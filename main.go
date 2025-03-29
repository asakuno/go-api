package main

import (
	"log"
	"os"

	"github.com/asakuno/go-api/command"
	"github.com/asakuno/go-api/constants"
	"github.com/asakuno/go-api/middlewares"
	"github.com/asakuno/go-api/providers"
	"github.com/asakuno/go-api/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/samber/do"
)

func init() {
	if os.Getenv("APP_ENV") != constants.ENUM_RUN_PRODUCTION {
		err := godotenv.Load()
		if err != nil {
			log.Println("Warning: .env file not found")
		}
	}
}

func args(injector *do.Injector) bool {
	if len(os.Args) > 1 {
		flag := command.Commands(injector)
		return flag
	}

	return true
}

func main() {
	var (
		injector = do.New()
	)
	providers.RegisterDependencies(injector)

	if !args(injector) {
		return
	}

	server := gin.Default()
	server.Use(middlewares.CORSMiddleware())

	routes.RegisterRoutes(server, injector)

	port := os.Getenv("GOLANG_PORT")
	if port == "" {
		port = "8888"
	}

	serve := "0.0.0.0:" + port

	log.Printf("Server is running on %s\n", serve)
	if err := server.Run(serve); err != nil {
		log.Fatalf("Error running server: %v", err)
	}
}
