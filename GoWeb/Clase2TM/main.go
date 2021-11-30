package main

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"
	"strings"

	"github.com/gin-gonic/gin"
)

type persona struct {
	ID       int    `json:"id"`
	Nombre   string `json:"nombre"`
	Apellido string `json:"apellido"`
	Edad     int    `json:"edad"`
}

var personas []persona

func AddPersona(ctx *gin.Context) {
	var per persona
	err := ctx.ShouldBindJSON(&per)
	if err != nil {
		ctx.JSON(400, gin.H{
			"error": err.Error(),
		})
		// ctx.String(400, "Se produjo un error: %v", err.Error())

	} else {

		// per.ID = len(personas) + 1
		if len(personas) == 0 {
			per.ID = 1
		} else {
			per.ID = personas[len(personas)-1].ID + 1
		}
		personas = append(personas, per)
		ctx.JSON(200, per)
	}

}

func GetPersonas(ctx *gin.Context) {

	token := ctx.GetHeader("token")

	if token != "" {
		if token == "123456" {
			if len(personas) > 0 {
				ctx.JSON(200, personas)
			} else {
				ctx.String(200, "No hay personas cargadas")
			}
		} else {
			ctx.String(401, "Token incorrecto")
		}
	} else {
		ctx.String(400, "No ingreso un token")
	}
}

func filtrar(slicePersonas []persona, campo string, valor string) []persona {
	var filtrado []persona

	var per persona
	tipos := reflect.TypeOf(per)
	i := 0
	for i = 0; i < tipos.NumField(); i++ {
		// fmt.Println(i, "->", tipos.Field(i).Name)
		if strings.ToLower(tipos.Field(i).Name) == campo {
			break
		}
	}

	for _, v := range slicePersonas {
		var cadena string
		cadena = fmt.Sprintf("%v", reflect.ValueOf(v).Field(i).Interface())
		if strings.Contains(cadena, valor) {
			// if reflect.ValueOf(v).Field(i).Interface() == valor {
			filtrado = append(filtrado, v)
		}
	}

	return filtrado
}

func FiltrarPersonas(ctx *gin.Context) {
	var etiquetas []string
	etiquetas = append(etiquetas, "nombre", "apellido")

	var personasFiltradas []persona

	personasFiltradas = personas

	for _, v := range etiquetas {
		if len(ctx.Query(v)) != 0 && len(personasFiltradas) != 0 {
			personasFiltradas = filtrar(personasFiltradas, v, ctx.Query(v))
		}
	}

	if len(personasFiltradas) == 0 {
		ctx.String(200, "No hay coincidencias")
	} else {
		ctx.JSON(200, personasFiltradas)
	}

}

func LoadData(ctx *gin.Context) {
	data, err := os.ReadFile("./personas.json")
	if err != nil {
		ctx.String(400, "No se pudo abrir el archivo")
	} else {
		json.Unmarshal(data, &personas)
		ctx.String(200, "Personas cargadas")
	}
}

func main() {

	router := gin.Default()
	personas := router.Group("/personas")

	personas.POST("/add", AddPersona)
	personas.GET("/", GetPersonas)
	personas.GET("/loadData", LoadData)
	personas.GET("/filtros", FiltrarPersonas)

	router.Run()

}
