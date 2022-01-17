package app

import (
	"fmt"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/yescorihuela/walmart-products/config"
	"github.com/yescorihuela/walmart-products/domain"
	"github.com/yescorihuela/walmart-products/logger"
	"github.com/yescorihuela/walmart-products/services"
	"go.mongodb.org/mongo-driver/mongo"
)

type Server struct {
	engine   *gin.Engine
	logger   *logger.Logger
	database *mongo.Client
	httpAddr string
}

func NewServer(host string, port uint) Server {
	server := Server{
		engine:   gin.Default(), // New if your need incorporate middleware or your own logger
		logger:   logger.New(os.Stdout, logger.LevelDebug),
		database: config.ConnectToMongoDB(),
		httpAddr: fmt.Sprintf("%s:%d", host, port),
	}

	server.engine.Use(cors.Default())
	server.engine.Use(gin.Recovery())
	server.registerRoutes()
	return server
}

func (s *Server) Run() error {
	s.logger.PrintInfo(fmt.Sprintf("Server running on %v", s.httpAddr), nil)
	return s.engine.Run(s.httpAddr)
}

func (s *Server) registerRoutes() {
	ph := ProductHandlers{
		// productService: services.NewProductService(domain.NewRepositoryStub()),
		productService: services.NewProductService(domain.NewProductRepositoryMongo(s.database)),
	}
	s.engine.GET("/products", ph.GetAllProducts)
	s.engine.GET("/products/search", ph.SearchByCriteria)
}
