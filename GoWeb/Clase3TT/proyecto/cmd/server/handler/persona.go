package handler

import (
	"os"
	"strconv"

	personas "github.com/extmatperez/w2GoPrueba/GoWeb/Clase3TT/proyecto/internal/personas"
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

func validarToken(ctx *gin.Context) bool {
	token := ctx.GetHeader("token")
	if token == "" {
		ctx.String(400, "Falta token")
		return false
	}
	tokenENV := os.Getenv("TOKEN")
	if token != tokenENV {
		ctx.String(404, "Token incorrecto")
		return false
	}

	return true
}

func (per *Persona) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		if !validarToken(ctx) {
			return
		}

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

		if !validarToken(ctx) {
			return
		}

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

func (controller *Persona) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		if !validarToken(ctx) {
			return
		}

		var per request

		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)

		if err != nil {
			ctx.String(400, "El id es invalido")
		}

		err = ctx.ShouldBindJSON(&per)

		if err != nil {
			ctx.String(400, "Error en el body")
		} else {
			personaActualizada, err := controller.service.Update(int(id), per.Nombre, per.Apellido, per.Edad)
			if err != nil {
				ctx.JSON(400, err.Error())
			} else {
				ctx.JSON(200, personaActualizada)
			}
		}

	}
}

func (controller *Persona) UpdateNombre() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		if !validarToken(ctx) {
			return
		}

		var per request

		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)

		if err != nil {
			ctx.String(400, "El id es invalido")
		}

		err = ctx.ShouldBindJSON(&per)

		if err != nil {
			ctx.String(400, "Error en el body")
		} else {
			if per.Nombre == "" {
				ctx.String(404, "El nombre no puede estar vacío")
				return
			}
			personaActualizada, err := controller.service.UpdateNombre(int(id), per.Nombre)
			if err != nil {
				ctx.JSON(400, err.Error())
			} else {
				ctx.JSON(200, personaActualizada)
			}
		}

	}
}

func (controller *Persona) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		if !validarToken(ctx) {
			return
		}

		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)

		if err != nil {
			ctx.String(400, "El id es invalido")
		}

		err = controller.service.Delete(int(id))
		if err != nil {
			ctx.JSON(400, err.Error())
		} else {
			ctx.String(200, "La persona %d ha sido eliminada", id)
		}

	}
}
