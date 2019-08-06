package repository

import (
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
)
import "model"

var UsersStore map[string]*model.User

type UserRepository interface {
	Create(string, string)
	//Update(string)
	FindByEmail(string) (*model.User, error)
}

type InmemoryUserRepository struct {
}

func (repository InmemoryUserRepository) Create(name string, email string) {
	newUser := model.User{Name: name, Email: email}
	UsersStore[name] = &newUser
}

func (repository InmemoryUserRepository) FindByEmail(email string) (*model.User, error) {
	if x, ok := UsersStore[email]; ok {
		return x, nil
	} else {
		return nil, errors.New("user not found")
	}
}

type DbUserRepository struct {
	Db *gorm.DB
}

func (repository DbUserRepository) Create(name string, email string) {
	//現状ただ新規作成
	newUser := model.User{Name: name, Email: email}
	result := repository.Db.Create(&newUser)
	fmt.Println(result.Error)
}

func (repository DbUserRepository) FindByEmail(email string) (*model.User, error) {
	user := &model.User{}
	if err := repository.Db.Where("email = ?", email).First(user).Error; err != nil {
		fmt.Println(err)
		return nil, err
	}
	fmt.Println(user.Email)
	return user, nil
}
