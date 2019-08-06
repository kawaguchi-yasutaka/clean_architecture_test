package useCases

import (
	"request"
	"response"
)

type userLoginCase interface {
	Handle(request request.UserLoginRequest) response.UserLoginResponse
}
