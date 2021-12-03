package main

import (
	"log"
	"os"

	"github.com/extmatperez/w2GoPrueba/GoWeb/Clase4TT/proyecto/cmd/server/handler"
	"github.com/extmatperez/w2GoPrueba/GoWeb/Clase4TT/proyecto/docs"
	personas "github.com/extmatperez/w2GoPrueba/GoWeb/Clase4TT/proyecto/internal/personas"
	"github.com/extmatperez/w2GoPrueba/GoWeb/Clase4TT/proyecto/pkg/store"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
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

// @title MELI Bootcamp API
// @version 1.0
// @description This API Handle MELI Products.
// @termsOfService https://developers.mercadolibre.com.ar/es_ar/terminos-y-condiciones

// @contact.name API Support
// @contact.url https://developers.mercadolibre.com.ar/support

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
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

	docs.SwaggerInfo.Host = os.Getenv("HOST")
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	//	router.Use(TokenAuthMiddleware())

	router.GET("/personas/get", TokenAuthMiddleware(), controller.GetAll())
	router.POST("/personas/add", controller.Store())
	router.PUT("/personas/:id", controller.Update())
	router.PATCH("/personas/:id", controller.UpdateNombre())
	router.DELETE("/personas/:id", controller.Delete())

	router.Run()
}
