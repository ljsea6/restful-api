// go:build wireinject
//go:build wireinject
// +build wireinject

package config

import (
	"github.com/google/wire"
	"github.com/ljsea6/restful-api/app/controller"
	"github.com/ljsea6/restful-api/app/repository"
	"github.com/ljsea6/restful-api/app/service"
)

var db = wire.NewSet(ConnectToDB)

var userRepoSet = wire.NewSet(repository.UserRepositoryImplInit,
	wire.Bind(new(repository.IUserRepository), new(*repository.UserRepositoryImpl)),
)

var userServiceSet = wire.NewSet(service.UserServiceImplInit,
	wire.Bind(new(service.IUserService), new(*service.UserServiceImpl)),
)

var userCtrlSet = wire.NewSet(controller.UserControllerImplInit,
	wire.Bind(new(controller.IUserController), new(*controller.UserControllerImpl)),
)

func Init() *Initialization {
	wire.Build(NewInitialization, db, userRepoSet, userServiceSet, userCtrlSet)
	return nil
}
