package user

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type CreateUserRequest struct {
	FcmToken string `query:"fcmToken"`
}

type User struct {
	UserId string `json:"userId"`
	Name   string `json:"name"`
}

func CreateUser(c echo.Context) error {
	var reqBody CreateUserRequest
	err := c.Bind(&reqBody)
	if err != nil {
		return c.String(http.StatusBadRequest, "Bad Request")
	}
	fmt.Printf("id: %s\n", reqBody.FcmToken)
	u := &User{
		UserId: "test",
		Name:   "Test",
	}
	return c.JSON(http.StatusOK, u)
}
