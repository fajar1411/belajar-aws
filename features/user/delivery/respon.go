package delivery

import "fajar/clean/features/user"

type UserRespon struct {
	// ID    uint   `json:"id`
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
	Role  string `json:"role" form:"role"`
	// Book  []BookRespon `json:"book"`
}

func UserCoreToUserRespon(dataCore user.UserCore) UserRespon { // data user core yang ada di controller yang memanggil user repository
	return UserRespon{
		// ID:    dataCore.ID,
		Name:  dataCore.Name,
		Email: dataCore.Email,
		Phone: dataCore.Phone,
		Role:  dataCore.Role,
	}
}
func ListUserCoreToUserRespon(dataCore []user.UserCore) []UserRespon { //data user.core data yang diambil dari entities ke respon struct
	var ResponData []UserRespon

	for _, value := range dataCore { //memanggil paramete data core yang berisi data user core
		ResponData = append(ResponData, UserCoreToUserRespon(value)) // mengambil data mapping dari user core to respon
	}
	return ResponData
}
