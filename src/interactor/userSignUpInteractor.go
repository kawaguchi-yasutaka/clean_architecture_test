package interactor

import (
	"repository"
)
import "request"
import "response"

type UserSignUpInteractor struct {
	UserRepository repository.UserRepository
}

func (UserSignUpInteractor UserSignUpInteractor) Handle(request request.UserCreateRequest) (response.UserCreateResponse, error) {
	if _, error := UserSignUpInteractor.UserRepository.FindByName(request.Email); error == nil {
		return response.UserCreateResponse{}, error
	}
	UserSignUpInteractor.UserRepository.Create(request.Name)
	return response.UserCreateResponse{Name: request.Name}, nil
}
