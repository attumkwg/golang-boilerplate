package repository

import (
	"github.com/jinzhu/gorm"
)

// DBRepository はDB接続のための定義です。
type DBRepository struct {
	DB DB
}

// Begin はDB接続を開始するメソッドです。
func (db *DBRepository) Begin() *gorm.DB {
	return db.DB.Begin()
}

// Connect はDB接続を行うメソッドです。
func (db *DBRepository) Connect() *gorm.DB {
	return db.DB.Connect()
}
