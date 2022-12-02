package repository

import (
	_user "fajar/clean/features/user"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string
	Email    string
	Password string
	Phone    string `gorm:"type:varchar(15)"`
	Address  string
	Role     string
	// Books    []Book
}

func FromUserCore(dataCore _user.UserCore) User { //fungsi yang mengambil data dari entities usercore dan merubah data ke user gorm(model.go)
	userGorm := User{
		Name:     dataCore.Name,
		Email:    dataCore.Email, //mapping data core ke data gorm model
		Password: dataCore.Password,
		Phone:    dataCore.Phone,
		Address:  dataCore.Address,
		Role:     dataCore.Role,
	} ///formating data berdasarkan data gorm dan kita mapping data yang kita butuhkan untuk inputan  klien
	return userGorm //insert user
}

// methode karena mengambil data dari user struct
func (dataModel *User) ModelsToCore() _user.UserCore { //fungsi yang mengambil data dari  user gorm(model.go)  dan merubah data ke entities usercore
	return _user.UserCore{

		Name:    dataModel.Name,
		Email:   dataModel.Email, //mapping data model ke data entitis user core
		Phone:   dataModel.Phone,
		Address: dataModel.Address,
		Role:    dataModel.Role,
		// Book:    ListModelTOBookCore(dataModel.Books),
	}
}

// merubah slice[] dari user gorm(model.go)  ke data slice[]  entities usercore
func ListModelTOCore(dataModel []User) []_user.UserCore { //fungsi yang mengambil data dari  user gorm(model.go)  dan merubah data ke entities usercore
	var dataCore []_user.UserCore
	for _, value := range dataModel {
		dataCore = append(dataCore, value.ModelsToCore())
	}
	return dataCore //  untuk menampilkan data ke controller
}
