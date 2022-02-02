package database

import (
	"fmt"

	"github.com/borischen0203/Go-todolist/config"
	"github.com/borischen0203/Go-todolist/dto"
	"github.com/borischen0203/Go-todolist/logger"

	// "gorm.io/driver/mysql"
	// "gorm.io/gorm"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func DbConn() (db *gorm.DB) {
	// source := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", config.Env.DBUser, config.Env.DBPassword, config.Env.DBHost, config.Env.DBPort, config.Env.DBName)
	source := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", config.Env.DBUser, config.Env.DBPassword, config.Env.DBHost, config.Env.DBName)
	//Connect to mysql
	// conn, err := gorm.Open(mysql.Open(source), &gorm.Config{})
	db, err := gorm.Open("mysql", source)
	if err != nil {
		logger.Error.Fatalf("Setup MySQL connect error %+v\n", err)
	}
	return db
}

//Mysql db setup
func Setup() {
	// init mysql.
	db := DbConn()

	defer db.Close()
	db.Debug().DropTableIfExists(&dto.TodoItemModel{})
	db.Debug().AutoMigrate(&dto.TodoItemModel{})

	db.DB().SetConnMaxLifetime(10)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(10)
}
