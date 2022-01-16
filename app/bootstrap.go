package app

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/yescorihuela/walmart-products/domain"
	"github.com/yescorihuela/walmart-products/services"
)

type Server struct {
	engine   *gin.Engine
	httpAddr string
}

func NewServer(host string, port uint) Server {
	server := Server{
		engine:   gin.Default(), // New if your need incorporate middleware or your own logger | Default is better this case
		httpAddr: fmt.Sprintf("%s:%d", host, port),
	}
	server.registerRoutes()
	return server
}

func (s *Server) Run() error {
	log.Println("Server running on ", s.httpAddr)
	return s.engine.Run(s.httpAddr)
}

func (s *Server) registerRoutes() {
	ph := ProductHandlers{
		productService: services.NewProductService(domain.NewRepositoryStub()),
	}
	s.engine.GET("/products", ph.GetAllProducts)
}
