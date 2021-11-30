package handler

import (
	personas "github.com/extmatperez/w2GoPrueba/GoWeb/Clase2TT/proyecto/internal/personas"
	"github.com/gin-gonic/gin"
)

type request struct {
	Nombre   string `json:"nombre"`
	Apellido string `json:"apellido"`
	Edad     int    `json:"edad"`
}

type Persona struct {
	service personas.Service
}

func NewPersona(ser personas.Service) *Persona {
	return &Persona{service: ser}
}

func (per *Persona) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		personas, err := per.service.GetAll()

		if err != nil {
			ctx.String(400, "Hubo un error %v", err)
		} else {
			ctx.JSON(200, personas)
		}
	}
}

func (controller *Persona) Store() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var perso request

		err := ctx.ShouldBindJSON(&perso)

		if err != nil {
			ctx.String(400, "Hubo un error al querer cargar una persona %v", err)
		} else {
			response, err := controller.service.Store(perso.Nombre, perso.Apellido, perso.Edad)
			if err != nil {
				ctx.String(400, "No se pudo cargar la persona %v", err)
			} else {
				ctx.JSON(200, response)
			}
		}

	}
}
