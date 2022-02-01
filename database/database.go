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

// const (
// 	UserName     string = "user"
// 	Password     string = "password"
// 	Addr         string = "127.0.0.1"
// 	Port         int    = 3306
// 	Database     string = "todolist"
// 	MaxLifetime  int    = 10
// 	MaxOpenConns int    = 10
// 	MaxIdleConns int    = 10
// )

func DbConn() (db *gorm.DB) {
	source := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", config.Env.DBUser, config.Env.DBPassword, config.Env.DBHost, config.Env.DBPort, config.Env.DBName)
	//Connect to mysql
	// conn, err := gorm.Open(mysql.Open(addr), &gorm.Config{})
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

}

// func Setup() {
// 	// var err error
// 	// addr := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True", UserName, Password, Addr, Port, Database)
// 	// addr := fmt.Sprintf("user:password@tcp(godockerDB)/todolist")
// 	source := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", config.Env.DBUser, config.Env.DBPassword, config.Env.DBHost, config.Env.DBPort, config.Env.DBName)
// 	//Connect to mysql
// 	// conn, err := gorm.Open(mysql.Open(addr), &gorm.Config{})
// 	_, err := gorm.Open(mysql.Open(source), &gorm.Config{})
// 	if err != nil {
// 		logger.Error.Fatalf("Setup MySQL connect error %+v\n", err)
// 	}
// 	// if err != nil {
// 	// 	panic("Could not connect with database!")
// 	// }
// 	//Set up ConnMaxLifetime/MaxIdleConns/MaxOpenConns
// 	// db, err1 := conn.DB()
// 	// if err1 != nil {
// 	// 	logger.Error.Fatalf("get db failed:", err) // return
// 	// }
// 	// db.SetConnMaxLifetime(time.Duration(MaxLifetime) * time.Second)
// 	// db.SetMaxIdleConns(MaxIdleConns)
// 	// db.SetMaxOpenConns(MaxOpenConns)

// 	// //create table
// 	// conn.Debug().AutoMigrate(&dto.TodoItemModel{})
// 	// //Check table exist or not
// 	// migrator := conn.Migrator()
// 	// has := migrator.HasTable(&dto.TodoItemModel{})
// 	// //has := migrator.HasTable("GG")
// 	// if !has {
// 	// 	fmt.Println("table not exist")
// 	// }
// }
