package hanlder

import (
	"net/http"
	"strconv"

	"github.com/HuXin0817/dots-and-boxes/server/internal/model"
	"github.com/HuXin0817/dots-and-boxes/src/game"
	"github.com/gin-gonic/gin"
)

func FindEnemyHandler(c *gin.Context) {
	Id, ok := c.GetQuery("id")
	if !ok {
		c.JSON(http.StatusBadRequest, &model.FindEnemyResponse{
			Error: "id not found",
		})
		return
	}
	id, err := strconv.ParseUint(Id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, &model.FindEnemyResponse{
			Error: err.Error(),
		})
		return
	}
	matched, isFirst, err := game.QueryMatch(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, &model.FindEnemyResponse{
			Error: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, &model.FindEnemyResponse{
		Waiting: !matched,
		IsFirst: isFirst,
	})
}
