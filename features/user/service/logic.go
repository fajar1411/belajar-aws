package service

import (
	"errors"
	"fajar/clean/features/user"

	"github.com/go-playground/validator/v10"
)

// bisnis Logic
type UserService struct {
	userRepository user.RepositoryEntities //data repository dri entities
	validate       *validator.Validate
}

// DeleteById implements user.ServiceEntities

// GetById implements user.ServiceEntities

// Update implements user.ServiceEntities

func New(repo user.RepositoryEntities) user.ServiceEntities { //dengan kembalian user.service
	return &UserService{
		userRepository: repo,
		validate:       validator.New(),
	}
}

// Create implements user.ServiceEntities
func (service *UserService) Create(input user.UserCore) (err error) {

	// if input.Name == "" || input.Email == "" || input.Password == "" || input.Address == "" || input.Phone == "" {
	// 	return errors.New("Name, email, password, phone dan addres harus diisi")
	// }
	input.Role = "User"
	if validateERR := service.validate.Struct(input); validateERR != nil {
		return validateERR
	}

	_, errCreate := service.userRepository.Create(input)
	if errCreate != nil {
		return errors.New("GAGAL MENAMBAH DATA , QUERY ERROR")
	}

	return nil
}

// GetAll implements user.ServiceEntities
func (service *UserService) GetAll() (data []user.UserCore, err error) {
	data, err = service.userRepository.GetAll() // memanggil struct entities repository yang ada di entities yang berisi coding logic
	return
}
func (service *UserService) Update(id int, input user.UserCore) error {
	// data, err = service.userRepository.Update(id int, input)
	// return
	errUpdate := service.userRepository.Update(id, input)
	if errUpdate != nil {
		return errors.New("GAGAL mengupdate data , QUERY ERROR")
	}

	return nil
}

func (service *UserService) GetById(id int) (data user.UserCore, err error) {
	data, err = service.userRepository.GetById(id) // memanggil struct entities repository yang ada di entities yang berisi coding logic
	return
}
func (service *UserService) DeleteById(id int) error {
	data := service.userRepository.DeleteById(id) // memanggil struct entities repository yang ada di entities yang berisi coding logic
	return data
}
