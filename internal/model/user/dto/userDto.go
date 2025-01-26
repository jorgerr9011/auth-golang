package dto

type UserDto struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

/*
type ListUserReq struct {
	Name      	string `json:"name,omitempty" form:"name"`
	Email 		string `json:`
}
*/

type ListUserReq struct {
	Name string `json:"name,omitempty" form:"name"`
}

type ListUserRes struct {
	Users []*UserDto `json:"users"`
}

type CreateUserReq struct {
	Name  string `json:"name" validate:"required"`
	Email string `json:"email" validate:"required"`
}

type UpdateUserReq struct {
	Name  string `json:"name,omitempty"`
	Email string `json:"email,omitempty"`
}
