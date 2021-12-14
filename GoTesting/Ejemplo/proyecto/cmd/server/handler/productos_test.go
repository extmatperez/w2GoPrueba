package handler

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/extmatperez/w2GoPrueba/GoTesting/Ejemplo/proyecto/pkg/store"

	productos "github.com/extmatperez/w2GoPrueba/GoTesting/Ejemplo/proyecto/internal/productos"
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
	db := store.New(store.FileType, "./productosTest.json")
	repo := productos.NewRepository(db)
	service := productos.NewService(repo)
	controller := NewProducto(service)

	router.Use(TokenAuthMiddleware())

	router.GET("/productos/get", controller.GetAll())
	router.POST("/productos/add", controller.Store())
	router.PUT("/productos/:id", controller.Update())
	router.DELETE("/productos/:id", controller.Delete())

	return router
}

func createRequestTest(method string, url string, body string) (*http.Request, *httptest.ResponseRecorder) {
	req := httptest.NewRequest(method, url, bytes.NewBuffer([]byte(body)))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("token", "123456")

	return req, httptest.NewRecorder()
}

func TestGetAllFuncional(t *testing.T) {
	router := createServer()

	req, res := createRequestTest(http.MethodGet, "/productos/get", "")

	router.ServeHTTP(res, req)

	assert.Equal(t, 200, res.Code)
}

func TestStoreFuncional(t *testing.T) {
	router := createServer()

	var nuevoProducto request = request{Nombre: "Cocina", Precio: 69852}

	sliceDeByte, _ := json.Marshal(nuevoProducto)

	req, res := createRequestTest(http.MethodPost, "/productos/add", string(sliceDeByte))

	router.ServeHTTP(res, req)

	assert.Equal(t, 204, res.Code)
}
