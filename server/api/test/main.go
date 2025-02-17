package main

import (
	"errors"
	"log"
	"sync/atomic"

	"github.com/HuXin0817/dots-and-boxes/server/api"
	"github.com/HuXin0817/dots-and-boxes/src/ai"
	"github.com/HuXin0817/dots-and-boxes/src/board"
	"github.com/HuXin0817/dots-and-boxes/src/mock"
	"github.com/HuXin0817/dots-and-boxes/src/model"
)

const N = 100

var Cli = api.New("127.0.0.1:8080")

func Run(m string) (id uint64, err error) {
	b := board.NewBoardV2()
	if id, err = Cli.StartGame(); err != nil {
		return id, err
	}
	isFirst, err := Cli.WaitJoin(id)
	if err != nil {
		return id, err
	}
	g, err := ai.New(m)
	if err != nil {
		return id, err
	}
	var gameExit string
	getEdge := func() (edge model.Edge) {
		if (b.Turn == model.Player1Turn && isFirst) || (b.Turn == model.Player2Turn && !isFirst) {
			edge = g(b)
			gameExit, err = Cli.AddEdge(id, edge)
		} else {
			edge, gameExit, err = Cli.GetOnlinePlayerEdge(id, b.Step)
		}
		if gameExit != "" || err != nil {
			return 0
		}
		return edge
	}
	mock.Run(
		b,
		func() bool { return err == nil && gameExit == "" && b.NotOver() },
		func(edge model.Edge) { b.Add(edge) },
		getEdge,
		getEdge,
	)
	if err == nil && gameExit != "" {
		err = errors.New(gameExit)
	}
	return id, err
}

func main() {
	var runNumber atomic.Uint64
	for {
		if runNumber.Load() < N {
			runNumber.Add(1)
			go func() {
				defer runNumber.Add(-1)
				if id, err := Run("L0"); err != nil {
					log.Printf("game %d: %s", id, err)
				} else {
					log.Printf("game %d: done", id)
				}
			}()
		}
	}
}
