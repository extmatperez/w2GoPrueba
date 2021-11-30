package main

import (
	"github.com/extmatperez/w2GoPrueba/GoWeb/Clase2TT/proyecto/cmd/server/handler"
	personas "github.com/extmatperez/w2GoPrueba/GoWeb/Clase2TT/proyecto/internal/personas"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	repo := personas.NewRepository()
	service := personas.NewService(repo)
	controller := handler.NewPersona(service)

	router.GET("/personas/get", controller.GetAll())
	router.POST("/personas/add", controller.Store())

	router.Run()
}
