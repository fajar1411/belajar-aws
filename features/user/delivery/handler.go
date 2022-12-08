package delivery

///handler = controller
import (
	"fajar/clean/features/user"
	"fajar/clean/utils/helper"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type UserDeliv struct {
	UserService user.ServiceEntities
}

func New(Service user.ServiceEntities, e *echo.Echo) {
	handler := &UserDeliv{
		UserService: Service,
	}
	e.GET("/user", handler.GetAll) // memanggil func getall
	e.POST("/user", handler.Create)
	e.PUT("/user/:id", handler.Update)
	e.GET("/user/:id", handler.GetById)
	e.DELETE("/user/:id", handler.DeleteById)
}

func (delivery *UserDeliv) GetAll(c echo.Context) error {
	result, err := delivery.UserService.GetAll() //memanggil fungsi service yang ada di folder service

	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.PesanGagalHelper("erorr read data"))
	}
	var ResponData = ListUserCoreToUserRespon(result)
	return c.JSON(http.StatusOK, helper.PesanDataBerhasilHelper("berhasil membaca  user", ResponData))
}

func (delivery *UserDeliv) Create(c echo.Context) error {
	Inputuser := UserRequest{} //penangkapan data user reques dari entities user
	errbind := c.Bind(&Inputuser)

	if errbind != nil {
		return c.JSON(http.StatusBadRequest, helper.PesanGagalHelper("erorr read data"+errbind.Error()))
	}
	dataCore := UserRequestToUserCore(Inputuser) //data mapping yang diminta create
	errResultCore := delivery.UserService.Create(dataCore)
	if errResultCore != nil {
		return c.JSON(http.StatusBadRequest, helper.PesanGagalHelper("erorr read data"+errResultCore.Error()))
	}
	return c.JSON(http.StatusCreated, helper.PesanSuksesHelper("berhasil create user"))
}
func (delivery *UserDeliv) Update(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	userInput := UserRequest{}

	errBind := c.Bind(&userInput) // menangkap data yg dikirim dari req body dan disimpan ke variabel
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.PesanGagalHelper("Error binding data "+errBind.Error()))
	}

	dataCore := UserRequestToUserCore(userInput)
	err := delivery.UserService.Update(id, dataCore)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.PesanGagalHelper("Failed update data"+err.Error()))
	}
	return c.JSON(http.StatusCreated, helper.PesanSuksesHelper("success Update data"))
}

func (delivery *UserDeliv) GetById(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	result, err := delivery.UserService.GetById(id) //memanggil fungsi service yang ada di folder service//jika return nya 2 maka variable harus ada 2

	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.PesanGagalHelper("erorr read data"))
	}
	var ResponData = UserCoreToUserRespon(result)
	return c.JSON(http.StatusOK, helper.PesanDataBerhasilHelper("berhasil membaca  user", ResponData))
}

func (delivery *UserDeliv) DeleteById(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	result := delivery.UserService.DeleteById(id) //memanggil fungsi service yang ada di folder service

	if result != nil {
		return c.JSON(http.StatusBadRequest, helper.PesanGagalHelper("erorr read data"))
	}

	return c.JSON(http.StatusCreated, helper.PesanSuksesHelper("success Hapus data"))
}
