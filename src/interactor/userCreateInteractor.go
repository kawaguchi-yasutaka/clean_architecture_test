package interactor

import "repository"
import "request"
import "response"

type UserCreateInteractor struct {
	UserRepository repository.UserRepository
}

func (userCreateInteractor UserCreateInteractor) Handle(request request.UserCreateRequest) (response.UserCreateResponse, error) {
	if _, error := userCreateInteractor.UserRepository.FindByName(request.Name); error == nil {
		return response.UserCreateResponse{}, error
	}
	userCreateInteractor.UserRepository.Save(request.Name)
	return response.UserCreateResponse{Name: request.Name}, nil
}
