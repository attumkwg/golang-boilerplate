package repository

import (
	"github.com/jinzhu/gorm"
)

// DB はDB接続するための定義です。
type DB interface {
	Begin() *gorm.DB
	Connect() *gorm.DB
}
