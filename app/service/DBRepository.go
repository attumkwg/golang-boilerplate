package service

import (
	"github.com/jinzhu/gorm"
)

// DBRepository はserviceで利用するための定義です。
type DBRepository interface {
	Begin() *gorm.DB
	Connect() *gorm.DB
}
