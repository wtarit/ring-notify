package main

import (
	"api/configs"
	"api/models"
	"api/router"
	"api/util"
	"log"
	"net/http"
	"os"

	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"

	_ "api/docs" // This line is necessary for go-swagger to find your docs!
)

type CustomValidator struct {
	Validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.Validator.Struct(i)
}

//	@title			Ring Notify API
//	@version		0.0.1
//	@description	API Specification for Ring Notify app.

//	@securityDefinitions.apikey	BearerAuth
//	@in							header
//	@name						Authorization
//	@description				Type "Bearer" followed by a space and JWT token.

func main() {
	if err := godotenv.Load(); err != nil && !os.IsNotExist(err) {
		log.Fatalln("Error loading .env")
	}

	configs.InitFirebase()
	configs.InitDatabase()
	if err := configs.InitSupabase(); err != nil {
		log.Fatalf("Failed to initialize Supabase: %v", err)
	}

	e := echo.New()
	e.Validator = &CustomValidator{Validator: validator.New()}
	e.Binder = &util.CustomBinder{}

	e.GET("/", healthCheck)

	// Swagger endpoint
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	router.InitRoute(e)

	e.Logger.Fatal(e.Start(":1323"))
}

// healthCheck godoc
//
//	@Summary		Health check
//	@Description	Get the status of the API
//	@Tags			healthcheck
//	@Produce		json
//	@Success		200	{object}	models.HealthCheckResponse
//	@Router			/ [get]
func healthCheck(c echo.Context) error {
	response := models.HealthCheckResponse{Status: "active"}
	return c.JSON(http.StatusOK, response)
}
