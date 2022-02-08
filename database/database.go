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

func DbConn() (db *gorm.DB) {
	//For docker-compose
	// source := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", config.Env.DBUser, config.Env.DBPassword, config.Env.DBHost, config.Env.DBName)

	//For
	source := "root:root@/todolist?charset=utf8&parseTime=True&loc=Local"

	//Connect to mysql
	db, err := gorm.Open("mysql", source)
	if err != nil {
		logger.Error.Fatalf("Setup MySQL connect error %+v\n", err)
	}
	return db
}

//Mysql db setup
func Setup() {
	// init mysql
	db := DbConn()

	defer db.Close()

	//drop exist table
	db.Debug().DropTableIfExists(&dto.TodoItemModel{})

	//Generate table
	db.Debug().AutoMigrate(&dto.TodoItemModel{})

	db.DB().SetConnMaxLifetime(10)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(10)
}

func AddItem(todo dto.TodoItemModel) interface{} {
	db := DbConn()
	db.Create(&todo)
	if db.Error != nil {
		logger.Error.Printf("[AddItem] todo: %+v\n", &todo)
		return false
	}
	result := db.Last(&todo)
	return result.Value
}

func GetItemByID(Id int) error {
	todo := &dto.TodoItemModel{}
	db := DbConn()
	result := db.First(&todo, Id)
	if result.Error != nil {
		logger.Error.Printf("[GetItemByID] Id not found: %+v\n", Id)
		return result.Error
	}
	return nil
}

func UpdateItemByID(Id int, Completed bool) error {
	todo := &dto.TodoItemModel{}
	db := DbConn()
	result := db.First(&todo, Id)
	if result.Error != nil {
		logger.Error.Printf("[UpdateItemByID] Id not found: %+v\n", Id)
		return result.Error
	}
	todo.Completed = Completed
	db.Save(&todo)
	return nil
}

func DeleteItemByID(Id int) error {
	todo := &dto.TodoItemModel{}
	db := DbConn()
	result := db.First(&todo, Id)
	if result.Error != nil {
		logger.Error.Printf("[DeleteItemByID] Id not found: %+v\n", Id)
		return result.Error
	} else {
		db.Delete(&todo)
	}
	return nil
}

func GetTodoItems(completed bool) interface{} {
	todos := []dto.TodoItemModel{}
	db := DbConn()
	result := db.Where("completed = ?", completed).Find(&todos).Value
	return result
}

// func AddItem(description string, completed bool) (string, error) {
// 	return "", nil

// }
