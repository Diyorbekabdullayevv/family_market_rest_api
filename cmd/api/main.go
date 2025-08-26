package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"practice_gin.com/internal/api/routers"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Failed to LOAD environment variables:", err)
		return
	}
	useGin()
}

func useGin() {
	server := gin.Default()

	routers.Router(server)

	port := os.Getenv("API_PORT")
	crtFile := "server.crt"
	keyFile := "server.key"
	err := server.RunTLS(port, crtFile, keyFile)
	if err != nil {
		fmt.Println("Failed to RUN the server:", err)
		return
	}
}
