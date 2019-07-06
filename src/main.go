package main

import (
	"controller"
	"fmt"
	"interactor"
	"repository"
	"request"
	"model"
)

func Init() controller.UserControllerInterface {
	repository.UsersStore = map[string]*model.User{}
	repository := repository.InmemoryUserRepository{}
	interactor := interactor.UserCreateInteractor{UserRepository: repository}
	return controller.UserController{UserCreateCase: interactor}
}

func main() {
	controller := Init()
	request := request.UserCreateRequest{Name: "kawaguchi"}
	response,_:= controller.Create(request)
	fmt.Println(response)
	fmt.Println(repository.UsersStore)
}

