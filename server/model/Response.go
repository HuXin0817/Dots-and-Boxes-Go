package model

import "github.com/HuXin0817/dots-and-boxes/src/model"

type GameStartResponse struct {
	Id uint64 `json:"id"`
}

type FindEnemyResponse struct {
	Waiting bool   `json:"waiting"`
	IsFirst bool   `json:"is_first"`
	Error   string `json:"error"`
}

type GameSyncResponse struct {
	GameExitMess string     `json:"game-exit-mess"`
	UnsyncEdge   model.Edge `json:"unsync-edge"`
	Error        string     `json:"error"`
}

type DropIDResponse struct {
	Error string `json:"error"`
}

type AddEdgeResponse struct {
	GameExitMess string `json:"game-exit-mess"`
	Error        string `json:"error"`
}
