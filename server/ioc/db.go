package ioc

import (
	"database/sql"
	"fmt"
	"github.com/spf13/viper"
)

func InitDB() *sql.DB {
	type Config struct {
		DSN string `yaml:"dsn"`
	}
	var c Config
	err := viper.UnmarshalKey("MySQL", &c)
	if err != nil {
		panic(fmt.Errorf("初始化配置失败: %s \n", err))
	}
	fmt.Println("c", c.DSN)
	db, err := sql.Open("mysql", c.DSN)
	if err != nil {
		panic(err)
	}

	return db
}
