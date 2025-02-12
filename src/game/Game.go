package game

import (
	"errors"
	"sync"
	"sync/atomic"
	"time"

	"github.com/HuXin0817/dots-and-boxes/src/board"
	"github.com/HuXin0817/dots-and-boxes/src/config"
	"github.com/HuXin0817/dots-and-boxes/src/match"
	"github.com/HuXin0817/dots-and-boxes/src/model"
)

var (
	ErrGameNotExist     = errors.New("game not exist")
	ErrGameAlreadyOver  = errors.New("game already over")
	ErrEdgeOutOfRange   = errors.New("edge out of range")
	ErrEdgeAlreadyExist = errors.New("edge already exist")
	ErrNotPlayerTurn    = errors.New("not player turn")
	ErrStepOutOfRange   = errors.New("step out of range")
)

type Game struct {
	board            board.BoardV2
	player1, player2 uint64
	lastActive       time.Time
}

var (
	lastId  atomic.Uint64
	gameMap = make(map[uint64]*Game)
	lock    sync.RWMutex
)

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
		if ok, player := match.Match(id); ok {
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

func AddEdge(id uint64, e model.Edge) (timeOut int, err error) {
	lock.RLock()
	game, ok := gameMap[id]
	lock.RUnlock()
	if !ok {
		return 0, ErrGameNotExist
	}
	if time.Since(game.lastActive) > config.PlayerTimeOut {
		return game.board.Turn, nil
	}
	if !game.board.NotOver() {
		return 0, ErrGameAlreadyOver
	}
	if game.board.Turn == model.Player1Turn && game.player1 != id {
		return 0, ErrNotPlayerTurn
	}
	if game.board.Turn == model.Player2Turn && game.player2 != id {
		return 0, ErrNotPlayerTurn
	}
	if e < 0 || e >= model.MaxEdge {
		return 0, ErrEdgeOutOfRange
	}
	if game.board.Contains(e) {
		return 0, ErrEdgeAlreadyExist
	}
	game.lastActive = time.Now()
	game.board.Add(e)
	return 0, nil
}

func Sync(id uint64, step model.Step) (edge model.Edge, timeOut int, err error) {
	lock.RLock()
	game, ok := gameMap[id]
	lock.RUnlock()
	if !ok {
		return 0, 0, ErrGameNotExist
	}
	if time.Since(game.lastActive) > config.PlayerTimeOut {
		return 0, game.board.Turn, nil
	}
	if step >= model.Step(model.MaxEdge) || step < 0 {
		return 0, 0, ErrStepOutOfRange
	}
	if step >= game.board.Step {
		return -1, 0, nil
	}
	edge = game.board.MoveRecord()[step]
	return edge, 0, nil
}
