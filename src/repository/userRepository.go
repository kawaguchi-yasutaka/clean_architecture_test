package repository

import "errors"
import "model"

var UsersStore map[string]*model.User

type UserRepository interface {
	Save(string)
	FindByName(string) (*model.User, error)
}

type InmemoryUserRepository struct {
}

func (repository InmemoryUserRepository) Save(name string) {
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
