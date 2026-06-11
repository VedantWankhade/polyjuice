package services

import "github.com/vedantwankhade/polyjuice/services/api/internal/core/ports/outbound"

type HelloService struct {
}

func (s *HelloService) Greet() string {
	return "Hello, world!"
}

func NewHelloService() outbound.HelloServices {
	return &HelloService{}
}
