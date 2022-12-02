package service

import (
	"errors"
	"fajar/clean/features/user"
	"fajar/clean/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAll(t *testing.T) {
	repo := new(mocks.RepoUser)
	returndata := []user.UserCore{{ID: 1, Name: "fariz", Email: "fariz@gmail.com", Phone: "08778788", Address: "jakarta", Role: "User"}}
	t.Run("Succes Get All data", func(t *testing.T) {
		repo.On("GetAll").Return(returndata, nil).Once()
		srv := New(repo)
		result, err := srv.GetAll()
		assert.Nil(t, err)
		assert.Equal(t, returndata[0].Name, result[0].Name)
		repo.AssertExpectations(t)
	})
	t.Run("failed Get All data", func(t *testing.T) {
		repo.On("GetAll").Return(nil, errors.New("Failed to Get data")).Once()
		srv := New(repo)
		result, err := srv.GetAll()
		assert.NotNil(t, err)
		assert.Nil(t, result)
		repo.AssertExpectations(t)
	})
}

func TestGetByID(t *testing.T) {
	repo := new(mocks.RepoUser)
	returndata := user.UserCore{Name: "fariz", Email: "fariz@gmail.com", Phone: "08778788", Address: "jakarta", Role: "User"}
	t.Run("Succes Get id user", func(t *testing.T) {
		repo.On("GetById", 1).Return(returndata, nil).Once()
		srv := New(repo)
		result, err := srv.GetById(int(returndata.ID)) ///parameter get by id berdasarkan di entitie
		assert.Nil(t, err)
		assert.Equal(t, returndata.ID, result.ID)
		repo.AssertExpectations(t)
	})
	t.Run("failed Get id user", func(t *testing.T) {
		repo.On("GetById", 1).Return(nil, errors.New("Failed Get id user")).Once()
		srv := New(repo)
		result, err := srv.GetById(int(returndata.ID))
		assert.NotNil(t, err)
		assert.Nil(t, result.ID)
		repo.AssertExpectations(t)
	})
}

func TestCreate(t *testing.T) {
	repo := new(mocks.RepoUser)
	t.Run("Succes Creat data", func(t *testing.T) {
		InputRepo := user.UserCore{Name: "fariz", Email: "fariz@gmail.com", Phone: "08778788", Password: "12345", Address: "jakarta", Role: "User"}
		Inputdata := user.UserCore{Name: "fariz", Email: "fariz@gmail.com", Phone: "08778788", Password: "12345", Address: "jakarta"}
		repo.On("Create", InputRepo).Return(1, nil).Once()
		srv := New(repo)
		err := srv.Create(Inputdata)
		assert.NoError(t, err)
		repo.AssertExpectations(t)
	})
	t.Run("Fail Create data,duplicate entry", func(t *testing.T) {
		InputRepo := user.UserCore{Name: "fariz", Email: "fariz@gmail.com", Phone: "08778788", Password: "12345", Address: "jakarta", Role: "User"}
		Inputdata := user.UserCore{Name: "fariz", Email: "fariz@gmail.com", Phone: "08778788", Password: "12345", Address: "jakarta"}
		repo.On("Create", InputRepo).Return(0, errors.New("failed create data")).Once()
		srv := New(repo)
		err := srv.Create(Inputdata)
		assert.NotNil(t, err)
		assert.Equal(t, "GAGAL MENAMBAH DATA , QUERY ERROR", err.Error())
		repo.AssertExpectations(t)
	})
	/*
		dalam testcase ini kita akan melakukan test saat kondisi validationnya error,
		sehingga pengkondisian (if) validation akan terpenuhi dan akan menjalankan perintah return.
		jadi fungsi Create yang ada di repository tidak akan dijalankan dalam test case ini. sooo kita tidak perlu memanggil mock "Create" repo.On
	*/
	t.Run("Failed Create user, name empty", func(t *testing.T) {
		// inputRepo := user.Core{Email: "alta@mail.id", Password: "qwerty", Phone: "081234", Address: "Jakarta", Role: "user"}
		inputData := user.UserCore{Email: "alta@mail.id", Password: "qwerty", Phone: "081234", Address: "Jakarta"}
		// repo.On("Create", inputRepo).Return(0, errors.New("failed to insert data, error query")).Once()
		srv := New(repo)
		err := srv.Create(inputData)
		assert.NotNil(t, err)
		repo.AssertExpectations(t)
	})
}
func TestUpdate(t *testing.T) {
	repo := new(mocks.RepoUser)
	returndata := user.UserCore{Name: "fariz", Email: "fariz@gmail.com", Phone: "08778788", Address: "jakarta", Role: "User"}
	t.Run("Succes Update user", func(t *testing.T) {
		returnupdate := user.UserCore{Name: "japri", Email: "japri@gmail.com", Phone: "08778788", Address: "jakarta", Role: "User"}
		repo.On("Update", 1, returndata).Return(nil).Once()
		srv := New(repo)
		err := srv.Update(int(returnupdate.ID), returnupdate) ///parameter get by id berdasarkan di entitie
		assert.Nil(t, err)
		assert.NoError(t, err)
		repo.AssertExpectations(t)
	})
	t.Run("failed Update user", func(t *testing.T) {
		returnupdate := user.UserCore{Name: "japri", Email: "japri@gmail.com", Phone: "08778788", Address: "jakarta", Role: "User"}
		repo.On("Update", 1, returndata).Return(nil, errors.New("Failed  update user")).Once()
		srv := New(repo)
		err := srv.Update(int(returnupdate.ID), returnupdate)
		assert.NotNil(t, err)
		repo.AssertExpectations(t)
	})
}
