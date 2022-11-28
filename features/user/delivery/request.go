package delivery

import "fajar/clean/features/user"

type UserRequest struct {
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
	Phone    string `json:"phone" form:"phone"`
	Address  string `json:"address" form:"address"`
	Role     string `json:"role" form:"role"`
}

func UserRequestToUserCore(data UserRequest) user.UserCore {
	return user.UserCore{
		Name:     data.Name,
		Password: data.Password,
		Email:    data.Email,
		Phone:    data.Phone,
		Address:  data.Address,
		Role:     data.Role,
	}
}
