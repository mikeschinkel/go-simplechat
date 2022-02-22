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
Main()
*/
func main() {
	loadEnv() // <-- Must be first
	shared.Init()
	dal.Init()
	startServer() // <-- Must be last
}

/**
Load environment variables from ".env" files.
*/
func loadEnv() {
	env := os.Args[1]
	cwd, _ := os.Getwd()
	path := cwd + "/env/" + env + ".env"
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
	routes.Init(engine)
	// Start the server
	engine.Run()
}
