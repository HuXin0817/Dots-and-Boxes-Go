package container

import (
	"github.com/HuXin0817/dots-and-boxes/src/model"
)

type EdgeList struct {
	m   [model.MaxEdge]model.Edge
	len int64
}

func (l *EdgeList) Clear() { l.len = 0 }

func (l *EdgeList) Reset(e model.Edge) {
	l.m[0] = e
	l.len = 1
}

func (l *EdgeList) Empty() bool { return l.len == 0 }

func (l *EdgeList) Append(e model.Edge) {
	l.m[l.len] = e
	l.len++
}

func (l *EdgeList) Export() []model.Edge { return l.m[:l.len] }
