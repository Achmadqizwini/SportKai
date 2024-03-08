package delivery

import (
	"net/http"

	"github.com/Achmadqizwini/SportKai/features/user"
	"github.com/Achmadqizwini/SportKai/utils/helper"

	"github.com/labstack/echo/v4"
)

type UserDelivery struct {
	userService user.ServiceInterface
}

func New(service user.ServiceInterface, e *echo.Echo) {
	handler := &UserDelivery{
		userService: service,
	}

	e.POST("/users", handler.Create)

}

func (delivery *UserDelivery) Create(c echo.Context) error {
	userInput := user.Core{}
	errBind := c.Bind(&userInput)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Error binding data "+errBind.Error()))
	}

	err := delivery.userService.Create(userInput)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponse("failed insert data"+err.Error()))
	}
	return c.JSON(http.StatusOK, helper.SuccessResponse("success create new users"))
}
