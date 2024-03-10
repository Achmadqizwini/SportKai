package factory

import (
	userDelivery "github.com/Achmadqizwini/SportKai/features/user/delivery"
	userRepo "github.com/Achmadqizwini/SportKai/features/user/repository"
	userService "github.com/Achmadqizwini/SportKai/features/user/service"
	"database/sql"
	"github.com/gorilla/mux"
)

func InitFactory(r *mux.Router, db *sql.DB) {
	userRepoFactory := userRepo.New(db)
	userServiceFactory := userService.New(userRepoFactory)
	userDelivery.New(userServiceFactory, r)

}
