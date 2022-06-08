package server

import (
	"log"

	"github.com/FredsZDev/WebApi_GO_B2/server/routes"
	"github.com/gin-gonic/gin"
)

type Server struct {
	port   string
	server *gin.Engine
}

func NewServer() Server {
	return Server{
		port:   "5000",
		server: gin.Default(),
	}
}

func (s *Server) Run() {
	router := routes.Routes(s.server)

	log.Fatal(router.Run(":" + s.port))

}
