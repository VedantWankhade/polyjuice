package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/vedantwankhade/polyjuice/services/api/internal/adapters/api/handlers"
)

func HelloRouter(helloHandler handlers.HelloHandler) *gin.Engine {
	helloRouter := gin.Default()
	helloRouter.GET("/hello", helloHandler.Hello)
	return helloRouter
}
