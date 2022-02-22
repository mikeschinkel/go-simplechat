package simplechat

import (
	"github.com/gin-gonic/gin"
)

const serverStartMsg = "Gin server running on localhost"

type Server struct {
	*gin.Engine
}

func NewServer() *Server {
	return &Server{
		Engine: gin.Default(),
	}
}

// Start starts the Gin server.
func (s *Server) Start() error {

	// Add routers (Groups) to the gin-engine
	addRoutes(s.Engine)

	// Start the server
	return s.Engine.Run()
}
