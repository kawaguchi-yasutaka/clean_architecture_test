package repository

import (
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
)
import "model"

var UsersStore map[string]*model.User

type UserRepository interface {
	Create(string)
	//Update(string)
	FindByName(string) (*model.User, error)
}

type InmemoryUserRepository struct {
}

func (repository InmemoryUserRepository) Create(name string) {
	newUser := model.User{Name: name}
	UsersStore[name] = &newUser
}

func (repository InmemoryUserRepository) FindByName(name string) (*model.User, error) {
	if x, ok := UsersStore[name]; ok {
		return x, nil
	} else {
		return nil, errors.New("user not found")
	}
}

type DbUserRepository struct {
	Db *gorm.DB
}

func (repository DbUserRepository) Create(name string) {
	//現状ただ新規作成
	newUser := model.User{Name: name}
	result := repository.Db.Create(&newUser)
	fmt.Println(result.Error)
}

func (repository DbUserRepository) FindByName(name string) (*model.User, error) {
	user := &model.User{}
	if err := repository.Db.Where("name = ?", name).First(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}
