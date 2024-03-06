package factory

import (
	authDelivery "github.com/Achmadqizwini/SportKai/features/auth/delivery"
	authRepo "github.com/Achmadqizwini/SportKai/features/auth/repository"
	authService "github.com/Achmadqizwini/SportKai/features/auth/service"

	userDelivery "github.com/Achmadqizwini/SportKai/features/user/delivery"
	userRepo "github.com/Achmadqizwini/SportKai/features/user/repository"
	userService "github.com/Achmadqizwini/SportKai/features/user/service"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitFactory(e *echo.Echo, db *gorm.DB) {
	userRepoFactory := userRepo.New(db)
	userServiceFactory := userService.New(userRepoFactory)
	userDelivery.New(userServiceFactory, e)

	authRepoFactory := authRepo.New(db)
	authServiceFactory := authService.New(authRepoFactory)
	authDelivery.New(authServiceFactory, e)
}
