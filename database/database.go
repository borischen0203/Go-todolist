package database

import (
	"github.com/borischen0203/Go-todolist/dto"
	"github.com/borischen0203/Go-todolist/logger"

	// "gorm.io/driver/mysql"
	// "gorm.io/gorm"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var Db *gorm.DB

func DbConn() (da *gorm.DB) {
	//For docker-compose
	// source := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", config.Env.DBUser, config.Env.DBPassword, config.Env.DBHost, config.Env.DBName)

	//For
	source := "root:root@/todolist?charset=utf8&parseTime=True&loc=Local"

	//Connect to mysql
	da, err := gorm.Open("mysql", source)
	if err != nil {
		logger.Error.Fatalf("Setup MySQL connect error %+v\n", err)
	}
	return da
}

//Mysql db setup
func Setup() {
	// init mysql.
	Db := DbConn()

	defer Db.Close()

	//drop exist table
	Db.Debug().DropTableIfExists(&dto.TodoItemModel{})

	//Generate table
	Db.Debug().AutoMigrate(&dto.TodoItemModel{})

	Db.DB().SetConnMaxLifetime(10)
	Db.DB().SetMaxIdleConns(10)
	Db.DB().SetMaxOpenConns(10)
}

// func AddItem(description string, completed bool) (string, error) {
// 	return "", nil

// }
