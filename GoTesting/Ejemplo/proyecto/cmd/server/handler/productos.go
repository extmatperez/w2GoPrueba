package handler

import (
	"fmt"
	"strconv"

	productos "github.com/extmatperez/w2GoPrueba/GoTesting/Ejemplo/proyecto/internal/productos"
	"github.com/extmatperez/w2GoPrueba/GoTesting/Ejemplo/proyecto/pkg/web"
	"github.com/gin-gonic/gin"
)

type request struct {
	Nombre string  `json:"nombre"`
	Precio float64 `json:"precio"`
}

type Producto struct {
	service productos.Service
}

func NewProducto(ser productos.Service) *Producto {
	return &Producto{service: ser}
}

// func validarToken(ctx *gin.Context) bool {
// 	token := ctx.GetHeader("token")
// 	if token == "" {
// 		ctx.JSON(400, web.NewResponse(400, nil, "Falta token"))
// 		// ctx.String(400, "Falta token")
// 		return false
// 	}
// 	tokenENV := os.Getenv("TOKEN")
// 	if token != tokenENV {
// 		ctx.JSON(404, web.NewResponse(404, nil, "Token incorrecto"))
// 		// ctx.String(404, "Token incorrecto")
// 		return false
// 	}

// 	return true
// }

//2 test

func (per *Producto) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		Productos, err := per.service.GetAll()
		if err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, fmt.Sprintf("Hubo un error %v", err)))
		} else {
			ctx.JSON(200, web.NewResponse(200, Productos, ""))
		}
	}
}

// 3 test

func (controller *Producto) Store() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var prod request
		err := ctx.ShouldBindJSON(&prod)
		if err != nil {
			ctx.JSON(400, web.NewResponse(400, fmt.Sprintf("Hubo un error al querer cargar una Producto %v", err), ""))
		} else {
			response, err := controller.service.Store(prod.Nombre, prod.Precio)
			if err != nil {
				ctx.JSON(400, web.NewResponse(400, fmt.Sprintf("No se pudo cargar la Producto %v", err), ""))
			} else {
				ctx.JSON(204, web.NewResponse(200, response, ""))
			}
		}
	}
}

//4 test

func (controller *Producto) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var prod request
		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.String(400, "El id es invalido")
			return
		}
		err = ctx.ShouldBindJSON(&prod)
		if err != nil {
			ctx.String(400, "Error en el body")
		} else {
			ProductoActualizada, err := controller.service.Update(int(id), prod.Nombre, prod.Precio)
			if err != nil {
				ctx.JSON(400, err.Error())
			} else {
				ctx.JSON(200, ProductoActualizada)
			}
		}

	}
}

func (controller *Producto) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)

		if err != nil {
			ctx.String(400, "El id es invalido")
			return
		}

		err = controller.service.Delete(int(id))
		if err != nil {
			ctx.JSON(400, err.Error())
		} else {
			ctx.String(200, "El Producto %d ha sido eliminada", id)
		}

	}
}
