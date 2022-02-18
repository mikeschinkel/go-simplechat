package main

import (
	"fmt"
	"os"
	"path/filepath"
	"simple-chat-app/server/src/daos"
	"simple-chat-app/server/src/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

const (
	serverStartMsg = "Gin server running on localhost"
	envFolderPath  = "../server/env"
)

/**
Main
*/
func main() {
	env := os.Args[2]
	loadEnv(env)
	daos.InitConn()
	startServer()
}

/**
Load environment variables from ".env" files.
*/
func loadEnv(env string) {
	path := filepath.Join(envFolderPath, env+".env")
	err := godotenv.Load(path)
	if err != nil {
		fmt.Println(err.Error())
	}
}

/**
Start the Gin server.
*/
func startServer() {
	engine := gin.Default()
	engine.GET("/", func(c *gin.Context) {
		c.String(200, serverStartMsg)
	})
	// Add routers (Groups) to the gin-engine
	routes.SetupApiRouter(engine)
	// Start the server
	engine.Run()
}
