package main

import (
	"fmt"
	"log"
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

	// server.LoadHTMLFiles("frontend/HTML/homepage.html", "frontend/HTML/get_products.html", "frontend/HTML/get_product.html")
	server.LoadHTMLFiles("frontend/HTML/*" )
	server.LoadHTMLGlob("frontend/HTML/*")
	server.Static("/static", "./frontend")

	routers.Router(server)

	err := server.SetTrustedProxies([]string{"127.0.0.1"})
	if err != nil {
		log.Fatal(err)
	}

	// crtFile := "server.crt"
	// keyFile := "server.key"
	// err := server.RunTLS(port, crtFile, keyFile)
	// err = server.RunTLS(port, crtFile, keyFile)

	port := os.Getenv("API_PORT")
	err = server.Run(port)
	if err != nil {
		fmt.Println("Failed to RUN the server:", err)
		return
	}
}
