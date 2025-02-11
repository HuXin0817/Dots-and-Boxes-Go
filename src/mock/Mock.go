package mock

import (
	"sync"

	"github.com/HuXin0817/dots-and-boxes/src/ai"
	"github.com/HuXin0817/dots-and-boxes/src/model"
	"github.com/HuXin0817/dots-and-boxes/src/model/board"
)

func Run(b *board.BoardV2, Add func(model.Edge), GetEdge1, GetEdge2 func() model.Edge) {
	for b.NotOver() {
		if b.Turn == model.Player1Turn {
			Add(GetEdge1())
		} else {
			Add(GetEdge2())
		}
	}
}

func RunAILocal(m1, m2 func(v2 *board.BoardV2) model.Edge) (Player1Score, Player2Score int) {
	b := board.NewBoardV2()
	Run(b,
		func(edge model.Edge) { b.Add(edge) },
		func() model.Edge { return m1(b) },
		func() model.Edge { return m2(b) },
	)
	return b.Player1Score, b.Player2Score
}

func RunAILocalN(m1, m2 func(v2 *board.BoardV2) model.Edge, N int) (Player1Score, Player2Score int) {
	for range N {
		s1, s2 := RunAILocal(m1, m2)
		Player1Score += s1
		Player2Score += s2
	}
	return
}

func RunAILocalM(m1, m2 string, N int) (Player1Score, Player2Score int) {
	l1, err := ai.New(m1)
	if err != nil {
		panic(err)
	}
	l2, err := ai.New(m2)
	if err != nil {
		panic(err)
	}
	return RunAILocalN(l1, l2, N)
}

func RunAILocalMParallel(m1, m2 string, N int, T int) (Player1Score, Player2Score int) {
	m := N / T
	var wg sync.WaitGroup
	wg.Add(T - 1)
	for range T - 1 {
		go func() {
			s1, s2 := RunAILocalM(m1, m2, m)
			Player1Score += s1
			Player2Score += s2
			wg.Done()
		}()
	}
	s1, s2 := RunAILocalM(m1, m2, N-m*(T-1))
	Player1Score += s1
	Player2Score += s2
	wg.Wait()
	return
}
