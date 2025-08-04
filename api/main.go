package main

import (
	"api/configs"
	"api/models"
	"api/notify"
	"api/user"
	"log"
	"net/http"
	"os"

	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"

	_ "api/docs" // This line is necessary for go-swagger to find your docs!
)

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

	e := echo.New()
	e.Validator = &user.CustomValidator{Validator: validator.New()}

	e.GET("/", healthCheck)
	e.POST("/notify/call", notify.Call)
	e.POST("/user/create", user.CreateUser)

	// Swagger endpoint
	e.GET("/swagger/*", echoSwagger.WrapHandler)

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
