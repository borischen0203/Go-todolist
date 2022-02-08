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

//Add todo item to DB
func AddItem(todo dto.TodoItemModel) interface{} {
	db := DbConn()
	err := db.Create(&todo)
	if err.Error != nil {
		logger.Error.Printf("[AddItem] add item error: %+v\n", &todo)
		return err.Error
	}
	result := db.Last(&todo)
	return result.Value
}

//Get todo item by ID
func GetItemByID(Id int) error {
	todo := &dto.TodoItemModel{}
	db := DbConn()
	result := db.First(&todo, Id)
	if result.Error != nil {
		logger.Error.Printf("[GetItemByID] query todo id error: %+v\n", Id)
		return result.Error
	}
	return nil
}

//Update todo item by ID
func UpdateItemByID(Id int, Completed bool) error {
	todo := &dto.TodoItemModel{}
	db := DbConn()
	result := db.First(&todo, Id)
	if result.Error != nil {
		logger.Error.Printf("[UpdateItemByID] query todo id error: %+v\n", Id)
		return result.Error
	}
	todo.Completed = Completed
	db.Save(&todo)
	return nil
}

//Delete item by ID
func DeleteItemByID(Id int) error {
	todo := &dto.TodoItemModel{}
	db := DbConn()
	result := db.First(&todo, Id)
	if result.Error != nil {
		logger.Error.Printf("[DeleteItemByID] query todo id error: %+v\n", Id)
		return result.Error
	}

	deleteResult := db.Delete(&todo)
	if deleteResult.Error != nil {
		logger.Error.Printf("[DeleteItemByID] delete id error: %+v\n", Id)
		return deleteResult.Error
	}

	return nil
}

//Get todo items by completed status
func GetTodoItems(completed bool) interface{} {
	todos := []dto.TodoItemModel{}
	db := DbConn()
	queryResult := db.Where("completed = ?", completed).Find(&todos)
	if queryResult.Error != nil {
		logger.Error.Printf("[GetTodoItems] query todo items status error: %+v\n", completed)
		return queryResult.Error
	}

	return queryResult.Value
}
