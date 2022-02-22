package main

import (
	"fmt"
	"os"
	"simple-chat-app/server/src/dal"
	"simple-chat-app/server/src/routes"
	"simple-chat-app/server/src/shared"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

const (
	serverStartMsg = "Gin server running on localhost"
	envFolderPath  = "../server/env"
)

/**
Load environment variables from ".env" files.
*/
func init() {
	env := os.Args[1]
	cwd, _ := os.Getwd()
	path := cwd + "/env/" + env + ".env"
	err := godotenv.Load(path)
	if err != nil {
		fmt.Println(err.Error())
	}
}

/**
Main()
*/
func main() {
	shared.Init()
	dal.Init()
	startServer() // <-- Must be last
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
	routes.Init(engine)
	// Start the server
	engine.Run()
}
