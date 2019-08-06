package useCases

import "request"
import "response"

type UserSignUpCase interface {
	Handle(request.UserCreateRequest) (response.UserCreateResponse, error)
}
