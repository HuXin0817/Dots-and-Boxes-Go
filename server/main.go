package main

import (
	"errors"
	"flag"
	"log"
	"net/http"
	"runtime"
	"strconv"
	"sync"
	"sync/atomic"
	"time"

	conf "github.com/HuXin0817/dots-and-boxes/server/config"
	model2 "github.com/HuXin0817/dots-and-boxes/server/model"
	"github.com/HuXin0817/dots-and-boxes/src/board"
	"github.com/HuXin0817/dots-and-boxes/src/config"
	"github.com/HuXin0817/dots-and-boxes/src/model"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var configFile = flag.String("f", "etc/config.yaml", "the config file")

var (
	ErrGameNotExist     = errors.New("game not exist")
	ErrGameAlreadyOver  = errors.New("game already over")
	ErrEdgeOutOfRange   = errors.New("edge out of range")
	ErrEdgeAlreadyExist = errors.New("edge already exist")
	ErrNotPlayerTurn    = errors.New("not player turn")
	ErrStepOutOfRange   = errors.New("step out of range")
)

type Game struct {
	board                    board.BoardV2
	player1, player2         uint64
	player1Exit, player2Exit bool
	lastActive               time.Time
}

var (
	matcher   *uint64
	matchlock uintptr

	lastId  atomic.Uint64
	gameMap = make(map[uint64]*Game)
	lock    sync.RWMutex
)

func matchLock() {
	for !atomic.CompareAndSwapUintptr(&matchlock, 0, 1) {
		runtime.Gosched()
	}
}

func matchUnlock() { atomic.StoreUintptr(&matchlock, 0) }

func Match(id uint64) (success bool, player uint64) {
	matchLock()
	defer matchUnlock()

	success = matcher != nil
	if success {
		player = *matcher
		matcher = nil
	} else {
		matcher = &id
	}

	return success, player
}

func Drop(id uint64) bool {
	matchLock()
	defer matchUnlock()
	if matcher == nil {
		return false
	}
	if *matcher != id {
		return false
	}
	matcher = nil
	return true
}

func clean() {
	lock.Lock()
	defer lock.Unlock()
	for id, game := range gameMap {
		if time.Since(game.lastActive) > 5*time.Minute {
			delete(gameMap, id)
		}
	}
}

func init() {
	go func() {
		for range time.Tick(time.Minute) {
			clean()
		}
	}()
}

func NewPlayer() (id uint64) {
	id = lastId.Add(1)
	go func() {
		if ok, player := Match(id); ok {
			StartGame(id, player)
		}
	}()
	return id
}

func StartGame(player1, player2 uint64) {
	game := &Game{
		board:      *board.NewBoardV2(),
		player1:    player1,
		player2:    player2,
		lastActive: time.Now(),
	}
	lock.Lock()
	gameMap[player1] = game
	gameMap[player2] = game
	lock.Unlock()
}

func QueryMatch(id uint64) (matched, isFirst bool, err error) {
	now := time.Now()
	lock.RLock()
	game, ok := gameMap[id]
	lock.RUnlock()
	if ok {
		game.lastActive = now
		return true, game.player1 == id, nil
	}
	return false, false, nil
}

func checkGameStatus(game *Game, id uint64) (mess string, err error) {
	if game.player1Exit {
		if id == game.player1 {
			return "you exit", nil
		} else {
			return "enemy exit", nil
		}
	}
	if game.player2Exit {
		if id == game.player2 {
			return "you exit", nil
		} else {
			return "enemy exit", nil
		}
	}
	if time.Since(game.lastActive) > config.PlayerTimeOut {
		if id == game.player1 {
			if game.board.Turn == model.Player1Turn {
				return "your time out", nil
			} else {
				return "enemy time out", nil
			}
		} else if id == game.player2 {
			if game.board.Turn == model.Player2Turn {
				return "your time out", nil
			} else {
				return "enemy time out", nil
			}
		}
	}
	return "", nil
}

func DropPlayer(id uint64) error {
	lock.RLock()
	game, ok := gameMap[id]
	lock.RUnlock()
	if !ok {
		return ErrGameNotExist
	}
	if game.player1 == id {
		game.player1Exit = true
	} else {
		game.player2Exit = true
	}
	return nil
}

func AddEdge(id uint64, e model.Edge) (mess string, err error) {
	lock.RLock()
	game, ok := gameMap[id]
	lock.RUnlock()
	if !ok {
		return "", ErrGameNotExist
	}
	if mess, err = checkGameStatus(game, id); mess != "" || err != nil {
		return mess, err
	}
	if !game.board.NotOver() {
		return "", ErrGameAlreadyOver
	}
	if game.board.Turn == model.Player1Turn && game.player1 != id {
		return "", ErrNotPlayerTurn
	}
	if game.board.Turn == model.Player2Turn && game.player2 != id {
		return "", ErrNotPlayerTurn
	}
	if e < 0 || e >= model.MaxEdge {
		return "", ErrEdgeOutOfRange
	}
	if game.board.Contains(e) {
		return "", ErrEdgeAlreadyExist
	}
	game.lastActive = time.Now()
	game.board.Add(e)
	return "", nil
}

func Sync(id uint64, step model.Step) (edge model.Edge, mess string, err error) {
	lock.RLock()
	game, ok := gameMap[id]
	lock.RUnlock()
	if !ok {
		return 0, "", ErrGameNotExist
	}
	if mess, err = checkGameStatus(game, id); mess != "" || err != nil {
		return 0, mess, err
	}
	if step >= model.Step(model.MaxEdge) || step < 0 {
		return 0, "", ErrStepOutOfRange
	}
	if step >= game.board.Step {
		return -1, "", nil
	}
	edge = game.board.MoveRecord()[step]
	return edge, "", nil
}

func main() {
	flag.Parse()
	Config, err := conf.NewFromFile(*configFile)
	if err != nil {
		log.Fatal(err)
	}
	r := gin.Default()
	r.Use(cors.Default())
	g := r.Group("/api/game")
	{
		g.POST("/start", func(c *gin.Context) {
			c.JSON(http.StatusOK, &model2.GameStartResponse{
				Id: NewPlayer(),
			})
		})
		g.GET("/find", func(c *gin.Context) {
			Id, ok := c.GetQuery("id")
			if !ok {
				c.JSON(http.StatusBadRequest, &model2.FindEnemyResponse{
					Error: "id not found",
				})
				return
			}
			id, err := strconv.ParseUint(Id, 10, 64)
			if err != nil {
				c.JSON(http.StatusBadRequest, &model2.FindEnemyResponse{
					Error: err.Error(),
				})
				return
			}
			matched, isFirst, err := QueryMatch(id)
			if err != nil {
				c.JSON(http.StatusBadRequest, &model2.FindEnemyResponse{
					Error: err.Error(),
				})
				return
			}
			c.JSON(http.StatusOK, &model2.FindEnemyResponse{
				Waiting: !matched,
				IsFirst: isFirst,
			})
		})
		g.POST("dropid", func(c *gin.Context) {
			Id, ok := c.GetQuery("id")
			if !ok {
				c.JSON(http.StatusBadRequest, &model2.DropIDResponse{
					Error: "id not found",
				})
				return
			}
			id, err := strconv.ParseUint(Id, 10, 64)
			if err != nil {
				c.JSON(http.StatusBadRequest, &model2.DropIDResponse{
					Error: err.Error(),
				})
				return
			}
			if Drop(id) {
				c.JSON(http.StatusOK, &model2.DropIDResponse{})
				return
			}
			if err = DropPlayer(id); err != nil {
				c.JSON(http.StatusBadRequest, &model2.DropIDResponse{
					Error: err.Error(),
				})
				return
			}
			c.JSON(http.StatusOK, &model2.DropIDResponse{})
			return
		})
		g.POST("/add", func(c *gin.Context) {
			Id, ok := c.GetQuery("id")
			if !ok {
				c.JSON(http.StatusBadRequest, &model2.AddEdgeResponse{
					Error: "id not found",
				})
				return
			}
			id, err := strconv.ParseUint(Id, 10, 64)
			if err != nil {
				c.JSON(http.StatusBadRequest, &model2.AddEdgeResponse{
					Error: err.Error(),
				})
				return
			}
			Edge, ok := c.GetQuery("edge")
			if !ok {
				c.JSON(http.StatusBadRequest, &model2.AddEdgeResponse{
					Error: "edge not found",
				})
				return
			}
			edge, err := strconv.Atoi(Edge)
			if err != nil {
				c.JSON(http.StatusBadRequest, &model2.AddEdgeResponse{
					Error: err.Error(),
				})
				return
			}
			GameExitMess, err := AddEdge(id, model.Edge(edge))
			if err != nil {
				c.JSON(http.StatusBadRequest, &model2.AddEdgeResponse{
					Error: err.Error(),
				})
				return
			}
			c.JSON(http.StatusOK, &model2.AddEdgeResponse{
				GameExitMess: GameExitMess,
			})
		})
		g.GET("/sync", func(c *gin.Context) {
			Id, ok := c.GetQuery("id")
			if !ok {
				c.JSON(http.StatusBadRequest, &model2.GameSyncResponse{
					Error: "id not found",
				})
				return
			}
			id, err := strconv.ParseUint(Id, 10, 64)
			if err != nil {
				c.JSON(http.StatusBadRequest, &model2.GameSyncResponse{
					Error: err.Error(),
				})
				return
			}
			Step, ok := c.GetQuery("step")
			if !ok {
				c.JSON(http.StatusBadRequest, &model2.GameSyncResponse{
					Error: "step not found",
				})
				return
			}
			step, err := strconv.Atoi(Step)
			if err != nil {
				c.JSON(http.StatusBadRequest, &model2.GameSyncResponse{
					Error: err.Error(),
				})
				return
			}
			UnsyncEdge, GameExitMess, err := Sync(id, model.Step(step))
			if err != nil {
				c.JSON(http.StatusBadRequest, &model2.GameSyncResponse{
					Error: err.Error(),
				})
				return
			}
			if GameExitMess != "" {
				c.JSON(http.StatusOK, &model2.GameSyncResponse{
					GameExitMess: GameExitMess,
				})
				return
			}
			c.JSON(http.StatusOK, &model2.GameSyncResponse{
				UnsyncEdge: UnsyncEdge,
			})
		})
	}
	for range time.Tick(time.Second) {
		if err = r.Run(Config.ListenOn); err != nil {
			log.Println(err.Error() + ", retry.")
		}
	}
}
