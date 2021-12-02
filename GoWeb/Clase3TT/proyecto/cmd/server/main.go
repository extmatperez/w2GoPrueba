package main

import (
	"log"

	"github.com/extmatperez/w2GoPrueba/GoWeb/Clase3TT/proyecto/cmd/server/handler"
	personas "github.com/extmatperez/w2GoPrueba/GoWeb/Clase3TT/proyecto/internal/personas"
	"github.com/extmatperez/w2GoPrueba/GoWeb/Clase3TT/proyecto/pkg/store"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("No se pudo abrir el archivo .env")
	}

	router := gin.Default()

	db := store.New(store.FileType, "./personasSalida.json")
	repo := personas.NewRepository(db)
	service := personas.NewService(repo)
	controller := handler.NewPersona(service)

	router.GET("/personas/get", controller.GetAll())
	router.POST("/personas/add", controller.Store())
	router.PUT("/personas/:id", controller.Update())
	router.PATCH("/personas/:id", controller.UpdateNombre())
	router.DELETE("/personas/:id", controller.Delete())

	router.Run()
}
