package service

import (
	"github.com/jinzhu/gorm"
	"github.com/attumkwg/mplatmm/app/domain"
)

// UserRepository はserviceのための定義ファイルです。
type UserRepository interface {
	FindByID(db *gorm.DB, id int) (user domain.Users, err error)
}
