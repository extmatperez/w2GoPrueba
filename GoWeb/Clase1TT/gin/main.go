package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

var empleados []Empleado = []Empleado{{"Matias", "123", "Activo"}, {"Juan", "321", "Activo"}, {"Pedro", "0212", "Inactivo"}}

type Empleado struct {
	// Una etiqueta de struct se cierra con caracteres de acento grave `
	Nombre string `form:"name" json:"name"`
	Id     string `form:"id" json:"id"`
	Activo string `form:"active" json:"activa" `
}

func Buscar(ctx *gin.Context) {

	var emp Empleado

	if ctx.BindJSON(&emp) == nil {
		ctx.JSON(200, emp)
	}

}

func BuscarQuery(ctx *gin.Context) {

	var filtrados []*Empleado

	for i, v := range empleados {
		if ctx.Query("filtro") == v.Activo {
			filtrados = append(filtrados, &empleados[i])
		}
	}

	if len(filtrados) == 0 {
		ctx.String(400, "No se encontró nada")
	} else {
		ctx.JSON(200, filtrados)
	}

}

func BuscarEmpleado(ctx *gin.Context) {

	parametro := ctx.Param("id")
	var emp Empleado
	se := false
	for _, v := range empleados {
		if v.Id == parametro {
			emp = v
			se = true
			break
		}
	}

	if se {
		ctx.JSON(200, emp)
	} else {
		ctx.String(404, "No se encontro el empleado %s", parametro)
	}

}

func Ejemplo(ctx *gin.Context) {

	contenido := ctx.Request.Body
	header := ctx.Request.Header
	metodo := ctx.Request.Method

	fmt.Println("Recibi algo")
	fmt.Println("Metodo: ", metodo)
	fmt.Println("Cabecera:")

	for k, v := range header {
		fmt.Println(k, " : ", v)
	}

	fmt.Println("Contenido", contenido)
	salida := fmt.Sprintf("Termine %.2f", 2.6666)
	ctx.String(200, salida)
	salida = fmt.Sprintf("\nTermine %.2f", 56.222)
	ctx.String(200, salida)
	salida = fmt.Sprintf("\nTermineeeee %.2f", 15.5)
	ctx.String(200, salida)
	// ctx.JSON(400, "Salida")
}

func main() {
	router := gin.Default()

	router.GET("/ejemplo", Ejemplo)
	router.GET("/buscar/:id", BuscarEmpleado)
	router.GET("/crear", Buscar)
	router.GET("/buscarquery", BuscarQuery)

	router.Run()
}
