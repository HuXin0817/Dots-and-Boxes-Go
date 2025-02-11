package api

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"

	model2 "github.com/HuXin0817/dots-and-boxes/server/internal/model"
	"github.com/HuXin0817/dots-and-boxes/src/model"
	"github.com/bytedance/sonic"
)

var lock sync.Mutex

type Api struct {
	addr string
}

func New(Addr string) *Api {
	return &Api{
		addr: Addr,
	}
}

func (api *Api) StartGame() (id uint64, err error) {
	lock.Lock()
	defer lock.Unlock()
	resp, err := http.Post("http://"+api.addr+"/api/game/start", "application/json", nil)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}
	var GameStartResponse model2.GameStartResponse
	if err = sonic.Unmarshal(body, &GameStartResponse); err != nil {
		return 0, err
	}
	id = GameStartResponse.Id
	return id, nil
}

func (api *Api) waitJoin(id uint64) (isFirst, wait bool, err error) {
	lock.Lock()
	defer lock.Unlock()
	resp, err := http.Get(fmt.Sprintf("http://%s/api/game/find?id=%d", api.addr, id))
	if err != nil {
		return false, false, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return false, false, err
	}
	var FindEnemyResponse model2.FindEnemyResponse
	if err = sonic.Unmarshal(body, &FindEnemyResponse); err != nil {
		return false, false, err
	}
	if FindEnemyResponse.Error != "" {
		return false, false, errors.New(FindEnemyResponse.Error)
	}
	return FindEnemyResponse.IsFirst, FindEnemyResponse.Waiting, nil
}

func (api *Api) WaitJoin(id uint64) (isFirst bool, err error) {
	for range time.Tick(time.Second) {
		var wait bool
		if isFirst, wait, err = api.waitJoin(id); err != nil {
			return false, err
		}
		if !wait {
			return isFirst, nil
		}
	}
	panic("unreachable")
}

func (api *Api) getOnlinePlayerEdge(id uint64, step model.Step) (e model.Edge, timeOut int, err error) {
	lock.Lock()
	defer lock.Unlock()
	resp, err := http.Get(fmt.Sprintf("http://%s/api/game/sync?id=%d&step=%d", api.addr, id, step))
	if err != nil {
		return 0, 0, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, 0, err
	}
	var r model2.GameSyncResponse
	if err = sonic.Unmarshal(body, &r); err != nil {
		return 0, 0, err
	}
	if r.Error != "" {
		return 0, 0, errors.New(r.Error)
	}
	if r.TimeOut != 0 {
		return 0, r.TimeOut, nil
	}
	return r.UnsyncEdge, 0, nil
}

func (api *Api) GetOnlinePlayerEdge(id uint64, step model.Step) (e model.Edge, timeOut int, err error) {
	for range time.Tick(time.Second) {
		if e, timeOut, err = api.getOnlinePlayerEdge(id, step); err != nil {
			return 0, 0, err
		}
		if timeOut != 0 {
			return 0, timeOut, nil
		}
		if e != -1 {
			return e, 0, nil
		}
	}
	return 0, 0, nil
}

func (api *Api) AddEdge(id uint64, e model.Edge) (timeOut int, err error) {
	lock.Lock()
	defer lock.Unlock()
	resp, err := http.Post(fmt.Sprintf("http://%s/api/game/add?id=%d&edge=%d", api.addr, id, e), "application/json", nil)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}
	var r model2.AddEdgeResponse
	if err = sonic.Unmarshal(body, &r); err != nil {
		return 0, err
	}
	if r.Error != "" {
		return 0, errors.New(r.Error)
	}
	return r.TimeOut, nil
}
