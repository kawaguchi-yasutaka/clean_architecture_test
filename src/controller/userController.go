package controller

import "request"
import "response"
import "useCases"

type UserControllerInterface interface {
	SignUp(request.UserCreateRequest) (response.UserCreateResponse, error)
}

type UserController struct {
	UserSignUpCase useCases.UserSignUpCase
}

func (controller UserController) SignUp(request request.UserCreateRequest) (response.UserCreateResponse, error) {
	return controller.UserSignUpCase.Handle(request)
}
