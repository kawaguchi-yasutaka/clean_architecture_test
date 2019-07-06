package controller

import "request"
import "response"
import "useCases"

type UserControllerInterface interface {
	Create(request.UserCreateRequest) (response.UserCreateResponse, error)
}

type UserController struct {
	UserCreateCase useCases.UserCreateCase
}


func (controller UserController) Create(request request.UserCreateRequest) (response.UserCreateResponse, error) {
	return controller.UserCreateCase.Handle(request)
}
