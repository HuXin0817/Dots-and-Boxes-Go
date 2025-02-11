package main

import (
	"flag"
	"log"
	"time"

	conf "github.com/HuXin0817/dots-and-boxes/server/internal/config"
	"github.com/HuXin0817/dots-and-boxes/server/internal/hanlder"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var configFile = flag.String("f", "etc/config.yaml", "the config file")

func main() {
	flag.Parse()
	Config, err := conf.NewFromFile(*configFile)
	if err != nil {
		log.Fatal(err)
	}
	r := gin.Default()
	r.Use(cors.Default())
	g := r.Group("/api/game")
	{
		g.POST("/start", hanlder.GameStartHandler)
		g.GET("/find", hanlder.FindEnemyHandler)
		g.POST("/add", hanlder.AddEdgeHandler)
		g.GET("/sync", hanlder.GameSyncHandler)
	}
	for range time.Tick(time.Second) {
		if err = r.Run(Config.ListenOn); err != nil {
			log.Println(err.Error() + ", retry.")
		}
	}
}
