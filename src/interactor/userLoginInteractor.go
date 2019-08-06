package interactor

import (
	"model"
	"repository"
	"request"
	"response"
)

type UserLoginInteractor struct {
	UserRepository repository.UserRepository
}

func (userLoginInteractor UserLoginInteractor) handle(request request.UserLoginRequest) (response.UserLoginResponse, error) {
	//認証処理が必要
	var user *model.User
	var err error
	if user, err = userLoginInteractor.UserRepository.FindByEmail(request.Email); err != nil {
		return response.UserLoginResponse{}, err
	}
	return response.UserLoginResponse{Email: user.Email}, nil
}
