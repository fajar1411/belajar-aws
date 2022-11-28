package user

import "time"

type UserCore struct {
	ID        uint
	Name      string `validate:"required"`
	Password  string `validate:"required"`
	Email     string `validate:"required,email"`
	Phone     string `validate:"required"`
	Address   string `validate:"required"`
	Role      string `validate:"required"`
	CreatedAt time.Time
	UpdatedAt time.Time
	// Book      []BookCore
}

type ServiceEntities interface { //sebagai contract yang dibuat di layer service
	GetAll() (data []UserCore, err error) //yang returnnya(mengembalikan data core)
	Create(input UserCore) (err error)    // menambahkah data user berdasarkan data usercore
	Update(id int, input UserCore) error
	GetById(id int) (data UserCore, err error)
	DeleteById(id int) error
}

type RepositoryEntities interface { // berkaitan database
	GetAll() (data []UserCore, err error)
	Create(input UserCore) (row int, err error)
	Update(id int, input UserCore) error
	GetById(id int) (data UserCore, err error)
	DeleteById(id int) error
}
