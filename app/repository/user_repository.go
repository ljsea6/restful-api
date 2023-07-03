package repository

import (
	"github.com/ljsea6/restful-api/app/domain/dao"

	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type IUserRepository interface {
	FindAllUser() ([]dao.User, error)
	FindUserByID(id int) (dao.User, error)
	Save(user *dao.User) (dao.User, error)
	DeleteUserByID(id int) error
}

type UserRepositoryImpl struct {
	db *gorm.DB
}

func (r *UserRepositoryImpl) FindAllUser() ([]dao.User, error) {
	var users []dao.User

	var err = r.db.Preload("Role").Find(&users).Error
	if err != nil {
		log.Error("Got an error finding all couples. Error: ", err)
		return nil, err
	}

	return users, nil
}

func (r *UserRepositoryImpl) FindUserByID(id int) (dao.User, error) {
	user := dao.User{
		ID: id,
	}
	err := r.db.Preload("Role").First(&user).Error
	if err != nil {
		log.Error("Got and error when find user by id. Error: ", err)
		return dao.User{}, err
	}
	return user, nil
}

func (r *UserRepositoryImpl) Save(user *dao.User) (dao.User, error) {
	var err = r.db.Save(user).Error
	if err != nil {
		log.Error("Got an error when save user. Error: ", err)
		return dao.User{}, err
	}
	return *user, nil
}

func (r *UserRepositoryImpl) DeleteUserByID(id int) error {
	err := r.db.Delete(&dao.User{}, id).Error
	if err != nil {
		log.Error("Got an error when delete user. Error: ", err)
		return err
	}
	return nil
}

func UserRepositoryImplInit(db *gorm.DB) *UserRepositoryImpl {
	db.AutoMigrate(&dao.User{})
	return &UserRepositoryImpl{
		db: db,
	}
}
