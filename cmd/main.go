package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/karpov-dmitry-py/jwt-with-gin-and-gorm/internal/api"
	"github.com/karpov-dmitry-py/jwt-with-gin-and-gorm/internal/pkg/initializers"
	"github.com/karpov-dmitry-py/jwt-with-gin-and-gorm/internal/pkg/repository"
	"github.com/karpov-dmitry-py/jwt-with-gin-and-gorm/internal/pkg/service"
)

func main() {
	const (
		port = ":5000"
	)

	if err := initializers.LoadEnvVariables(); err != nil {
		log.Fatal(err.Error())
	}

	dbConn, err := repository.ConnectToDB()
	if err != nil {
		log.Fatal(err.Error())
	}

	if err = repository.MigrateDB(); err != nil {
		log.Fatal(err.Error())
	}

	userService := service.NewUserService(dbConn, service.NewJwtService())
	server := api.NewServer(gin.Default(), userService)

	log.Fatal(server.Router.Run(port))
}
