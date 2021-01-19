package repository

import (
	"errors"
	"github.com/jinzhu/gorm"
	"github.com/attumkwg/mplatmm/app/domain"
)

// UserRepository .
type UserRepository struct {}

// FindByID .
func (repo *UserRepository) FindByID(db *gorm.DB, id int) (user domain.Users, err error) {
	user = domain.Users{}
	db.First(&user, id)
	if user.Id <= 0 {
		return domain.Users{}, errors.New("user is not found")
	}
	return user, nil
}
