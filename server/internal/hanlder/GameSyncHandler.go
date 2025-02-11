package hanlder

import (
	"net/http"
	"strconv"

	"github.com/HuXin0817/dots-and-boxes/server/internal/game"
	model2 "github.com/HuXin0817/dots-and-boxes/server/internal/model"
	"github.com/HuXin0817/dots-and-boxes/src/model"
	"github.com/gin-gonic/gin"
)

func GameSyncHandler(c *gin.Context) {
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
	UnsyncEdge, TimeOut, err := game.Sync(id, model.Step(step))
	if err != nil {
		c.JSON(http.StatusBadRequest, &model2.GameSyncResponse{
			Error: err.Error(),
		})
		return
	}
	if TimeOut != 0 {
		c.JSON(http.StatusOK, &model2.GameSyncResponse{
			TimeOut: TimeOut,
		})
		return
	}
	c.JSON(http.StatusOK, &model2.GameSyncResponse{
		UnsyncEdge: UnsyncEdge,
	})
}
