package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vedantwankhade/polyjuice/services/api/internal/core/ports/outbound"
)

type HelloHandler struct {
	helloService outbound.HelloServices
}

func (handler *HelloHandler) Hello(c *gin.Context) {
	c.JSON(http.StatusOK, handler.helloService.Greet())
}

func NewHelloHandler(service outbound.HelloServices) *HelloHandler {
	return &HelloHandler{
		helloService: service,
	}
}
