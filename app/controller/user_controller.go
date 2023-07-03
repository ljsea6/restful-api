package controller

import (
	"github.com/ljsea6/restful-api/app/service"

	"github.com/gin-gonic/gin"
)

type IUserController interface {
	GetAllUserData(c *gin.Context)
	AddUserData(c *gin.Context)
	GetUserByID(c *gin.Context)
	UpdateUserData(c *gin.Context)
	DeleteUser(c *gin.Context)
}

type UserControllerImpl struct {
	svc service.IUserService
}

func (ctrl *UserControllerImpl) GetAllUserData(c *gin.Context) {
	ctrl.svc.GetAllUser(c)
}

func (ctrl *UserControllerImpl) AddUserData(c *gin.Context) {
	ctrl.svc.AddUserData(c)
}

func (ctrl *UserControllerImpl) GetUserByID(c *gin.Context) {
	ctrl.svc.GetUserByID(c)
}

func (ctrl *UserControllerImpl) UpdateUserData(c *gin.Context) {
	ctrl.svc.UpdateUserData(c)
}

func (ctrl *UserControllerImpl) DeleteUser(c *gin.Context) {
	ctrl.svc.DeleteUser(c)
}

func UserControllerImplInit(userService service.IUserService) *UserControllerImpl {
	return &UserControllerImpl{
		svc: userService,
	}
}
