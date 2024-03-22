package factory

import (
	"database/sql"
	userDelivery "github.com/Achmadqizwini/SportKai/features/user/delivery"
	userRepo "github.com/Achmadqizwini/SportKai/features/user/repository"
	userService "github.com/Achmadqizwini/SportKai/features/user/service"

	clubDelivery "github.com/Achmadqizwini/SportKai/features/club/delivery"
	clubRepo "github.com/Achmadqizwini/SportKai/features/club/repository"
	clubService "github.com/Achmadqizwini/SportKai/features/club/service"

	memberDelivery "github.com/Achmadqizwini/SportKai/features/clubMember/delivery"
	memberRepo "github.com/Achmadqizwini/SportKai/features/clubMember/repository"
	memberService "github.com/Achmadqizwini/SportKai/features/clubMember/service"

	authDelivery "github.com/Achmadqizwini/SportKai/features/auth/delivery"
	authRepo "github.com/Achmadqizwini/SportKai/features/auth/repository"
	authService "github.com/Achmadqizwini/SportKai/features/auth/service"
	"net/http"
)

func InitFactory(r *http.ServeMux, db *sql.DB) {
	userRepoFactory := userRepo.New(db)
	userServiceFactory := userService.New(userRepoFactory)
	userDelivery.New(userServiceFactory, r)

	memberRepoFactory := memberRepo.New(db)
	memberServiceFactory := memberService.New(memberRepoFactory)
	memberDelivery.New(memberServiceFactory, r)

	clubRepoFactory := clubRepo.New(db)
	clubServiceFactory := clubService.New(clubRepoFactory, memberRepoFactory)
	clubDelivery.New(clubServiceFactory, r)

	authRepoFactory := authRepo.New(db)
	authServiceFactory := authService.New(authRepoFactory)
	authDelivery.New(authServiceFactory, r)

}
