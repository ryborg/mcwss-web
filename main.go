package main

import (
	"github.com/ryborg/mcwss"
//	"io/ioutil"
//	"log"
	"net/http"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"

)

func main() {
router := gin.Default()
router.Use(static.Serve("/", static.LocalFile("./views",true)))
api := router.Group("/api")
	{
		api.GET("/", func(c *gin.Context){
			c.JSON(http.StatusOK, gin.H {
			"message": "pong",
			})
		})
	}
//router.Run(":8100")


	server := mcwss.NewServer(nil)

	server.OnConnection(func(player *mcwss.Player){
player.Tell("Hello %s!", player.Name())
})
	server.OnDisconnection(func(player *mcwss.Player){
	})
	go server.Run()
	router.Run(":8100")
}
