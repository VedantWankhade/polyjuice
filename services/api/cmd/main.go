package main

import (
	"github.com/vedantwankhade/polyjuice/services/api/internal/adapters/api/handlers"
	"github.com/vedantwankhade/polyjuice/services/api/internal/adapters/api/routers"
	"github.com/vedantwankhade/polyjuice/services/api/internal/core/services"
)

func main() {
	helloService := services.NewHelloService()
	helloHandler := handlers.NewHelloHandler(helloService)
	srv := routers.HelloRouter(*helloHandler)
	srv.Run(":8080")
}
