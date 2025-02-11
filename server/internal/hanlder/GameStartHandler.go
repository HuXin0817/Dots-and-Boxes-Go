package hanlder

import (
	"net/http"

	"github.com/HuXin0817/dots-and-boxes/server/internal/game"
	"github.com/HuXin0817/dots-and-boxes/server/internal/model"
	"github.com/gin-gonic/gin"
)

func GameStartHandler(c *gin.Context) {
	c.JSON(http.StatusOK, &model.GameStartResponse{
		Id: game.NewPlayer(),
	})
}
