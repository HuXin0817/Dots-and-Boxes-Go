package main

import (
	"fmt"
	"log"
	"sync/atomic"

	"github.com/HuXin0817/dots-and-boxes/server/api"
	"github.com/HuXin0817/dots-and-boxes/src/ai"
	"github.com/HuXin0817/dots-and-boxes/src/board"
	"github.com/HuXin0817/dots-and-boxes/src/model"
)

const N = 100

var Cli = api.New("127.0.0.1:8080")

func Run(gameID int64, m string) (err error) {
	b := board.NewBoardV2()
	id, err := Cli.StartGame()
	if err != nil {
		return err
	}
	isFirst, err := Cli.WaitJoin(id)
	if err != nil {
		return err
	}
	g, err := ai.New(m)
	if err != nil {
		return err
	}
	for b.NotOver() {
		var edge model.Edge
		var timeOut int
		if (b.Turn == model.Player1Turn && isFirst) || (b.Turn == model.Player2Turn && !isFirst) {
			edge = g(b)
			if timeOut, err = Cli.AddEdge(id, edge); err != nil {
				return err
			}
		} else {
			if edge, timeOut, err = Cli.GetOnlinePlayerEdge(id, b.Step); err != nil {
				return err
			}
		}
		if timeOut != 0 {
			return fmt.Errorf("game %d: time out", gameID)
		}
		b.Add(edge)
	}
	return nil
}

func main() {
	var id atomic.Int64
	n := 0
	for {
		if n < N {
			n++
			go func() {
				i := id.Add(1)
				if err := Run(i, "L0"); err != nil {
					log.Println(err)
				} else {
					log.Printf("game %d: done", i)
				}
				n--
			}()
		}
	}
}
