package main

import (
	"github.com/extmatperez/w2GoPrueba/GoWeb/Clase3TM/proyecto/cmd/server/handler"
	personas "github.com/extmatperez/w2GoPrueba/GoWeb/Clase3TM/proyecto/internal/personas"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	repo := personas.NewRepository()
	service := personas.NewService(repo)
	controller := handler.NewPersona(service)

	router.GET("/personas/get", controller.GetAll())
	router.POST("/personas/add", controller.Store())
	router.PUT("/personas/:id", controller.Update())
	router.PATCH("/personas/:id", controller.UpdateNombre())
	router.DELETE("/personas/:id", controller.Delete())

	router.Run()
}
