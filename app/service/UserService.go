package service

import (
	"github.com/attumkwg/mplatmm/app/domain"
)

// UserService はUserを取得するためのサービス層です。
type UserService struct {
	DB DBRepository
	User UserRepository
}

// Get はUserを取得します。
func (service *UserService) Get(id int) (user domain.UsersForGet, statusCode int, err error) {
	db := service.DB.Connect()
	foundUser, err := service.User.FindByID(db, id)
	if err != nil {
		return domain.UsersForGet{}, 404, err
	}
	user = foundUser.BuildForGet()
	return user, 200, nil
}
