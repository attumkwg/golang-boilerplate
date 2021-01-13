package infrastructure

import (
	"os"
	"github.com/BurntSushi/toml"
)

// Config はDBとRoutingの設定を定義する
type Config struct {
	Db struct {
		Host string `toml:"host"`
		Username string `toml:"username"`
		Password string `toml:"password"`
		DbName string `toml:"db_name"`
	}
	Routing struct {
		Port string `toml:"port"`
	}
}

// NewConfig は外部に設定内容を提供する
func NewConfig() *Config {
	c := new(Config)
	_, err := toml.DecodeFile("config.tml", &c)
	if err != nil {
		panic(err)
	}
	if os.Getenv("DB_ENV") == "production" {
		c.Db.Host = os.Getenv("DB_HOST")
		c.Db.Username = os.Getenv("DB_USER")
		c.Db.Password = os.Getenv("DB_PASS")
		c.Db.DbName = os.Getenv("DB_NAME")
		c.Routing.Port = os.Getenv("AP_PORT")
	}
	return c
}
