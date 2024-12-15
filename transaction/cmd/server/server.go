package server

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/phuckhoa33/web-crawler/internal/config"
	db "github.com/phuckhoa33/web-crawler/internal/domain/database"
)

type Server struct {
	Gin    *gin.Engine
	DB     *gorm.DB
	Config *config.Config
}

func NewServer(config *config.Config) *Server {
	return &Server{
		Gin:    gin.Default(),
		DB:     db.Connect(config),
		Config: config,
	}
}

func (server *Server) Run(appPort string) error {
	// Configure swagger info
	InitSwaggerInfo(server.Config)
	// Run application
	return server.Gin.Run(":" + appPort)
}
