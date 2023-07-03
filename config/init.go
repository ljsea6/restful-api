package config

import (
	"github.com/ljsea6/restful-api/app/controller"
	"github.com/ljsea6/restful-api/app/repository"
	"github.com/ljsea6/restful-api/app/service"
)

type Initialization struct {
	UserRepo repository.IUserRepository
	UserSvc  service.IUserService
	UserCtrl controller.IUserController
}

func NewInitialization(
	userRepo repository.IUserRepository,
	userSvc service.IUserService,
	userCtrl controller.IUserController,
) *Initialization {
	return &Initialization{
		UserRepo: userRepo,
		UserSvc:  userSvc,
		UserCtrl: userCtrl,
	}
}
