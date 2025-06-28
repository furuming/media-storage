package router

import (
	"storage/cmd/config"

	"github.com/gin-gonic/gin"
)

func New(env *config.Config) *gin.Engine {

	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	return router
}
