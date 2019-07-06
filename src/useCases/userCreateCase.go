package useCases

import "request"
import "response"


type UserCreateCase interface {
	Handle(request.UserCreateRequest) (response.UserCreateResponse, error)
}
