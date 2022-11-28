package factory

import (
	userDelivery "fajar/clean/features/user/delivery"
	userRepo "fajar/clean/features/user/repository" //alias
	userService "fajar/clean/features/user/service"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitFactory(e *echo.Echo, db *gorm.DB) {
	userRepofaktory := userRepo.New(db) //menginiasialisasi func new yang ada di repository
	userServiceFaktory := userService.New(userRepofaktory)
	userDelivery.New(userServiceFaktory, e)
}
