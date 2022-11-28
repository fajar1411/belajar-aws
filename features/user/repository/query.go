package repository

import (
	"errors"
	"fajar/clean/features/user"

	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

// DeleteById implements user.RepositoryEntities

// GetById implements user.RepositoryEntities

// Update implements user.RepositoryEntities

func New(db *gorm.DB) user.RepositoryEntities { // user.repository mengimplementasikan interface repository yang ada di entities
	return &userRepository{
		db: db,
	}

}

// Create implements user.Repository//create yang melekat di data userRepository yng ada di atas
func (repo *userRepository) Create(input user.UserCore) (row int, err error) { //parameter yang mengambil data usercore yang ada di entities
	userGorm := FromUserCore(input) //dari gorm model ke user core yang ada di entities

	tx := repo.db.Create(&userGorm) // proses insert data

	if tx.Error != nil {
		return -1, tx.Error
	}
	if tx.RowsAffected == 0 {
		return 0, errors.New("Insert failed")
	}
	return int(tx.RowsAffected), nil
}

// GetAll implements user.Repository
func (repo *userRepository) GetAll() (data []user.UserCore, err error) { // yang dibutuhkan adalah user core
	var users []User //mengambil data gorm model(model.go)
	tx := repo.db.Find(&users)
	if tx.Error != nil {
		return nil, tx.Error
	}

	var DataCore = ListModelTOCore(users) //mengambil data dari gorm model(file repository(model.go))
	// var DataCore []entities.UserCore
	// for _, val := range users {
	// 	var dataBuku []entities.BookCore
	// 	for _, valbook := range val.Books {
	// 		dataBuku = append(dataBuku, entities.BookCore{
	// 			ID:    valbook.ID,
	// 			Title: valbook.Title,
	// 		})
	// 	}
	// 	DataCore = append(DataCore, entities.UserCore{
	// 		ID:        val.ID,
	// 		Name:      val.Name,
	// 		Email:     val.Email,
	// 		Password:  val.Password,
	// 		Phone:     val.Phone,
	// 		Address:   val.Address,
	// 		CreatedAt: val.CreatedAt,
	// 		UpdatedAt: val.UpdatedAt,
	// 		Book:      dataBuku,
	// 	})
	// }
	return DataCore, nil

}
func (repo *userRepository) Update(id int, input user.UserCore) error {
	userGorm := FromUserCore(input)
	tx := repo.db.Model(&userGorm).Where("id = ?", id).Updates(&userGorm)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (repo *userRepository) GetById(id int) (data user.UserCore, err error) {
	var users User

	tx := repo.db.First(&users, id)

	if tx.Error != nil {

		return user.UserCore{}, tx.Error
	}
	gorms := users.ModelsToCore()
	return gorms, nil
}
func (repo *userRepository) DeleteById(id int) error {
	var users User
	tx := repo.db.Delete(&users, id)
	if tx.Error != nil {

		return tx.Error
	}
	return nil
}
