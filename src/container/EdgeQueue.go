package container

import (
	"github.com/HuXin0817/dots-and-boxes/src/config"
	"github.com/HuXin0817/dots-and-boxes/src/model"
	"github.com/stretchr/testify/assert"
)

type EdgeQueue struct {
	m          [model.MaxEdge]model.Edge
	front, end int
}

func (l *EdgeQueue) Clear() { l.front, l.end = 0, 0 }

func (l *EdgeQueue) Empty() bool { return l.front == l.end }

func (l *EdgeQueue) Pop() model.Edge {
	if config.DEBUG {
		assert.False(nil, l.Empty())
	}
	e := l.m[l.front]
	l.front++
	return e
}

func (l *EdgeQueue) Append(e model.Edge) {
	if config.DEBUG {
		assert.True(nil, l.end < int(model.MaxEdge))
	}
	l.m[l.end] = e
	l.end++
}
