package model

import (
	"github.com/HuXin0817/dots-and-boxes/src/model"
)

type GameStartResponse struct {
	Id uint64 `json:"id"`
}

type FindEnemyResponse struct {
	Waiting bool   `json:"waiting"`
	IsFirst bool   `json:"is_first"`
	Error   string `json:"error"`
}

type GameSyncResponse struct {
	TimeOut    int        `json:"timeout"`
	UnsyncEdge model.Edge `json:"unsync_edge"`
	Error      string     `json:"error"`
}

type AddEdgeResponse struct {
	TimeOut int    `json:"timeout"`
	Error   string `json:"error"`
}
