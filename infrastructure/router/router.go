package router

import (
	"github.com/gin-gonic/gin"
	"golang-bootcamp-2020/controllers"
)

var (
	router = gin.Default()
)

func StartRouter(c controllers.AppController) error {

	router.GET("health", c.GetHealth)

	router.GET("data/fetch", c.FetchData)

	router.GET("api/character", c.GetCharacter)

	return router.Run(":8081")
}
