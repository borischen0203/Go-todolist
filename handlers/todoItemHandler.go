package handlers

import (
	"net/http"
	"strconv"

	"github.com/borischen0203/Go-todolist/database"
	"github.com/borischen0203/Go-todolist/dto"
	"github.com/borischen0203/Go-todolist/errors"
	"github.com/borischen0203/Go-todolist/logger"
	"github.com/borischen0203/Go-todolist/services"
	"github.com/gin-gonic/gin"
)

//TODO: finish handler function

func CreateItem(c *gin.Context) {
	request := dto.TodoRequest{}
	c.Bind(&request)
	logger.Info.Printf("[CreateItem Handler] request=%+v\n", request)
	// todo := dto.TodoItemModel{Description: request.Descripion, Completed: false}
	// db := database.DbConn()
	// db.Create(&todo)
	// result := db.Last(&todo)
	statusCode, result, err := services.CreateItemService(request)
	switch statusCode {
	case 200:
		c.JSON(http.StatusOK, result)
	case 500:
		c.JSON(http.StatusInternalServerError, err)
	default:
		c.JSON(http.StatusInternalServerError, errors.InternalServerError)
	}

	// result, _ := database.AddItem(todo)
	// if result.Error != nil {
	// 	// fmt.Println("Create fail")
	// 	logger.Info.Printf("[CreateItem Handler - Create fail] result=%+v\n", result)
	// }
	// if result.RowsAffected != 1 {
	// 	// fmt.Println("RowsAffected Number fail")
	// 	logger.Info.Printf("[CreateItem Handler - RowsAffected Number fail] result=%+v\n", result)
	// }
	// response := db.Createç(&todo)

	// c.JSON(http.StatusOK, result)
}

func UpdateItem(c *gin.Context) {

	// id, _ := strconv.Atoi(c.Param("id"))
	request := dto.TodoUpdateRequest{}
	request.Id, _ = strconv.Atoi(c.Param("id"))
	c.Bind(&request)
	logger.Info.Printf("[UpdateItem Handler] request:%+v\n", request)
	statusCode, _, _ := services.UpdateItemService(request)
	switch statusCode {
	case 200:
		c.JSON(http.StatusOK, `{"updated": true}`)
	case 404:
		c.JSON(http.StatusNotFound, `{"updated": false}`)
	// case 500:
	// 	c.JSON(http.StatusInternalServerError, err)
	default:
		c.JSON(http.StatusInternalServerError, errors.InternalServerError)
	}

	// err := database.GetItemByID(id)
	// if err != nil {
	// 	c.JSON(http.StatusNotFound, `{"updated": false}`)
	// } else {
	// 	request := dto.TodoUpdateRequest{}
	// 	c.Bind(&request)
	// 	logger.Info.Printf("[UpdateItem Handler] id=%+v\n, completed=%+v\n", id, request.Completed)
	// 	// todo := &dto.TodoItemModel{}
	// 	// db := database.DbConn()
	// 	// db.First(&todo, id)
	// 	// todo.Completed = request.Completed
	// 	// db.Save(&todo)
	// 	database.UpdateItemByID(id, request.Completed)
	// c.JSON(http.StatusOK, `{"updated": true}`)
	// }
}

func DeleteItem(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	logger.Info.Printf("[DeleteItem Handler] request id:%+v\n", id)
	statusCode, _, err := services.DeleteItemService(id)
	switch statusCode {
	case 200:
		c.JSON(http.StatusOK, `{"deleted": true}`)
	case 404:
		c.JSON(http.StatusNotFound, `{"deleted": false, "error": "Record Not Found"}`)
	case 500:
		c.JSON(http.StatusInternalServerError, err)
	default:
		c.JSON(http.StatusInternalServerError, errors.InternalServerError)
	}

	// _, err := database.GetItemByID(id)
	// if err != nil {
	// 	c.JSON(http.StatusNotFound, `{"deleted": false, "error": "Record Not Found"}`)
	// } else {
	// 	logger.Info.Printf("[DeleteItem Handler] id=%+v\n", id)
	// 	// todo := &dto.TodoItemModel{}
	// 	// db := database.DbConn()
	// 	// db.First(&todo, id)
	// 	// db.Delete(&todo)
	// 	database.DeleteItemByID(id)
	// 	c.JSON(http.StatusOK, `{"deleted": true}`)
	// }
}

func GetCompletedItems(c *gin.Context) {
	logger.Info.Print("[GetCompletedItems Handler]")
	completedTodoItems := database.GetTodoItems(true)
	c.JSON(http.StatusOK, completedTodoItems)
}

func GetIncompleteItems(c *gin.Context) {
	logger.Info.Print("[GetIncompleteItems Handler]")
	incompleteTodoItems := database.GetTodoItems(false)
	c.JSON(http.StatusOK, incompleteTodoItems)
}

// func GetTodoItems(completed bool) interface{} {
// 	var todos []dto.TodoItemModel
// 	db := database.DbConn()
// 	TodoItems := db.Where("completed = ?", completed).Find(&todos).Value
// 	return TodoItems
// }

// func GetItemByID(Id int) bool {
// 	todo := &dto.TodoItemModel{}
// 	db := database.DbConn()
// 	result := db.First(&todo, Id)
// 	if result.Error != nil {
// 		logger.Error.Printf("[GetItemByID] Id not found: %+v\n", Id)
// 		return false
// 	}
// 	return true
// }
