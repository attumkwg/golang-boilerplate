package controller

import (
	"strconv"
	"github.com/attumkwg/mplatmm/app/repository"
	"github.com/attumkwg/mplatmm/app/service"
)

// UsersController はユーザー用のControllerです。
type UsersController struct {
	Service service.UserService
}

// NewUsersController はUsersControllerを新規作成するメソッドです。
func NewUsersController(db repository.DB) *UsersController {
	return &UsersController{
		Service: service.UserService{
			DB: &repository.DBRepository{ DB: db },
			User: &repository.UserRepository{},
		},
	}
}

// Get は一人のユーザー情報を返却します。
func (controller *UsersController) Get(c Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	user, statusCode, err := controller.Service.Get(id)
	if err != nil {
		c.JSON(statusCode, NewResponse(err.Error(), nil))
		return
	}
	c.JSON(statusCode, NewResponse("success", user))
}

// List は複数人分のユーザー情報を返却します。
func (controller *UsersController) List(c Context) {
	c.JSON(200, NewResponse("success", nil))
}


