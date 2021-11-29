package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type persona struct {
	Nombre              string
	Apellido            string `json:"apellidito"`
	FechaNacimiento     string
	Direccion_domicilio string `json:"direccion_domicilio"`
}

func saludar(c *gin.Context) {

	c.JSON(200, "Hola")
}

func main() {
	router := gin.Default()

	per := persona{"Matias", "Perez", "17/11/1994", "Chilecito"}

	var personas []persona

	personas = append(personas, per, per, per)

	// salida, _ := json.Marshal(per)
	// salidaStr := string(salida)
	router.GET("/hello", saludar)
	router.GET("/hello2", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"mensaje": personas,
		})
	})

	router.Run()
	// router.Run(":8080")
}
