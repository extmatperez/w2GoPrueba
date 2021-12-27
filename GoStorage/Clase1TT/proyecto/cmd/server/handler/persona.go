package handler

import (
	"fmt"
	"os"
	"strconv"

	personas "github.com/extmatperez/w2GoPrueba/GoStorage/Clase1TT/proyecto/internal/personas"
	"github.com/extmatperez/w2GoPrueba/GoStorage/Clase1TT/proyecto/pkg/web"
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
		ctx.JSON(400, web.NewResponse(400, nil, "Falta token"))
		// ctx.String(400, "Falta token")
		return false
	}
	tokenENV := os.Getenv("TOKEN")
	if token != tokenENV {
		ctx.JSON(404, web.NewResponse(404, nil, "Token incorrecto"))
		// ctx.String(404, "Token incorrecto")
		return false
	}

	return true
}

// ListProducts godoc
// @Summary List personas
// @Tags Persona
// @Description get personas
// @Accept  json
// @Produce  json
// @Param token header string true "token"
// @Success 200 {object} web.Response
// @Router /personas/get [get]
func (per *Persona) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		// if !validarToken(ctx) {
		// 	return
		// }

		personas, err := per.service.GetAll()

		if err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, fmt.Sprintf("Hubo un error %v", err)))
			// ctx.String(400, "Hubo un error %v", err)
		} else {
			ctx.JSON(200, web.NewResponse(200, personas, ""))
			// ctx.JSON(200, personas)
		}
	}
}

// StoreProducts godoc
// @Summary Store persona
// @Tags Persona
// @Description store persona
// @Accept  json
// @Produce  json
// @Param token header string true "token"
// @Param persona body request true "Persona to store"
// @Success 200 {object} web.Response
// @Router /personas/add [post]
func (controller *Persona) Store() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		// if !validarToken(ctx) {
		// 	return
		// }

		var perso request

		err := ctx.ShouldBindJSON(&perso)

		fmt.Println("Desde el controller: ", perso)
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

		// if !validarToken(ctx) {
		// 	return
		// }

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

		// if !validarToken(ctx) {
		// 	return
		// }

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

		// if !validarToken(ctx) {
		// 	return
		// }

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
