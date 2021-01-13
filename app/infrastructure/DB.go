package infrastructure

import (
	"github.com/jinzhu/gorm"
	// gormでmysqlを利用するためのパッケージ
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// DB はDB接続のための定義です。
type DB struct {
	Host string
	Username string
	Password string
	DBName string
	Connection *gorm.DB
}

// NewDB はコネクション取得のためのfunctionです。
func NewDB() *DB {
	c := NewConfig()
	return newDB(&DB{
		Host: c.Db.Host,
		Username: c.Db.Username,
		Password: c.Db.Password,
		DBName: c.Db.DbName,
	})
}

func newDB(d *DB) *DB {
	db, err := gorm.Open("mysql", d.Username + ":" + d.Password + "@tcp(" + d.Host + ")/" + d.DBName + "?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err.Error())
	}
	d.Connection = db
	return d
}

// Begin begins a transction
func (db *DB) Begin() *gorm.DB {
	return db.Connection.Begin()
}

// Connect returns the transaction
func (db *DB) Connect() *gorm.DB {
	return db.Connection
}
