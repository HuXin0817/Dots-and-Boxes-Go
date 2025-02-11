package hanlder

import (
	"net/http"
	"strconv"

	"github.com/HuXin0817/dots-and-boxes/server/internal/game"
	model2 "github.com/HuXin0817/dots-and-boxes/server/internal/model"
	"github.com/HuXin0817/dots-and-boxes/src/model"
	"github.com/gin-gonic/gin"
)

func AddEdgeHandler(c *gin.Context) {
	Id, ok := c.GetQuery("id")
	if !ok {
		c.JSON(http.StatusBadRequest, &model2.AddEdgeResponse{
			Error: "id not found",
		})
		return
	}
	id, err := strconv.ParseUint(Id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, &model2.AddEdgeResponse{
			Error: err.Error(),
		})
		return
	}
	Edge, ok := c.GetQuery("edge")
	if !ok {
		c.JSON(http.StatusBadRequest, &model2.AddEdgeResponse{
			Error: "edge not found",
		})
		return
	}
	edge, err := strconv.Atoi(Edge)
	if err != nil {
		c.JSON(http.StatusBadRequest, &model2.AddEdgeResponse{
			Error: err.Error(),
		})
		return
	}
	TimeOut, err := game.AddEdge(id, model.Edge(edge))
	if err != nil {
		c.JSON(http.StatusBadRequest, &model2.AddEdgeResponse{
			Error: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, &model2.AddEdgeResponse{
		TimeOut: TimeOut,
	})
}
