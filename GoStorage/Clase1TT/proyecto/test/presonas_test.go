package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/extmatperez/w2GoPrueba/GoStorage/Clase1TT/proyecto/cmd/server/handler"
	personas "github.com/extmatperez/w2GoPrueba/GoStorage/Clase1TT/proyecto/internal/personas"
	"github.com/extmatperez/w2GoPrueba/GoStorage/Clase1TT/proyecto/pkg/store"
	"github.com/extmatperez/w2GoPrueba/GoStorage/Clase1TT/proyecto/pkg/web"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func respondWithError(c *gin.Context, code int, message interface{}) {
	c.AbortWithStatusJSON(code, gin.H{"error": message})
}
func TokenAuthMiddleware() gin.HandlerFunc {
	requiredToken := os.Getenv("TOKEN")

	// We want to make sure the token is set, bail if not
	if requiredToken == "" {
		log.Fatal("Please set API_TOKEN environment variable")
	}

	return func(c *gin.Context) {
		token := c.GetHeader("token")

		if token == "" {
			respondWithError(c, 401, "API token required")
			return
		}

		if token != requiredToken {
			respondWithError(c, 401, "Invalid API token")
			return
		}

		c.Next()
	}
}

func createServer() *gin.Engine {
	router := gin.Default()
	_ = os.Setenv("TOKEN", "123456")
	db := store.New(store.FileType, "./personasSalidaTest.json")
	repo := personas.NewRepository(db)
	service := personas.NewService(repo)
	controller := handler.NewPersona(service)

	router.Use(TokenAuthMiddleware())

	router.GET("/personas/get", controller.GetAll())
	router.POST("/personas/add", controller.Store())
	router.PUT("/personas/:id", controller.Update())
	router.PATCH("/personas/:id", controller.UpdateNombre())
	router.DELETE("/personas/:id", controller.Delete())

	return router
}

func createRequestTest(method string, url string, body string) (*http.Request, *httptest.ResponseRecorder) {
	req := httptest.NewRequest(method, url, bytes.NewBuffer([]byte(body)))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("token", "123456")

	return req, httptest.NewRecorder()
}

func Test_GetPersonas(t *testing.T) {
	router := createServer()

	req, rr := createRequestTest(http.MethodGet, "/personas/get", "")

	router.ServeHTTP(rr, req)

	assert.Equal(t, 200, rr.Code)

	var respuesta web.Response

	err := json.Unmarshal(rr.Body.Bytes(), &respuesta)
	assert.Equal(t, 200, respuesta.Code)
	assert.Nil(t, err)
}

func Test_StorePersonas(t *testing.T) {
	router := createServer()

	nuevaPersona := personas.Persona{Nombre: "Juan",
		Apellido: "Pescie",
		Edad:     22}

	dataNueva, _ := json.Marshal(nuevaPersona)
	fmt.Println(string(dataNueva))
	req, rr := createRequestTest(http.MethodPost, "/personas/add", string(dataNueva))

	router.ServeHTTP(rr, req)

	assert.Equal(t, 200, rr.Code)

	var respuesta personas.Persona

	err := json.Unmarshal(rr.Body.Bytes(), &respuesta)
	assert.Equal(t, "Juan", respuesta.Nombre)
	assert.Nil(t, err)

	delete := fmt.Sprintf("/personas/%d", respuesta.ID)
	req, rr = createRequestTest(http.MethodDelete, delete, "")

	router.ServeHTTP(rr, req)

	assert.Equal(t, 200, rr.Code)

}
