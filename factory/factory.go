package factory

import (
	userDelivery "github.com/Achmadqizwini/SportKai/features/user/delivery"
	userRepo "github.com/Achmadqizwini/SportKai/features/user/repository"
	userService "github.com/Achmadqizwini/SportKai/features/user/service"
	"database/sql"
	"github.com/labstack/echo/v4"
)

func InitFactory(e *echo.Echo, db *sql.DB) {
	userRepoFactory := userRepo.New(db)
	userServiceFactory := userService.New(userRepoFactory)
	userDelivery.New(userServiceFactory, e)

}
