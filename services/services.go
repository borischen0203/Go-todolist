package services

import (
	"github.com/borischen0203/Go-todolist/database"
	"github.com/borischen0203/Go-todolist/dto"
	e "github.com/borischen0203/Go-todolist/errors"
	"github.com/borischen0203/Go-todolist/logger"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func CreateItemService(request dto.TodoRequest) (int64, interface{}, e.ErrorInfo) {
	todo := dto.TodoItemModel{Description: request.Descripion, Completed: false}
	result, err := database.AddItem(todo)
	if err != nil {
		logger.Error.Printf("[CreateItemService] add item error ,request=%+v\n", request)
		return 500, dto.TodoItemModel{}, e.InternalServerError
	}
	return 200, result, e.Ok
}

func UpdateItemService(request dto.TodoUpdateRequest) (int64, interface{}, e.ErrorInfo) {
	result, err := database.GetItemByID(request.Id)
	if err != nil {
		logger.Error.Printf("[UpdateItemService] id not found ,request id:%+v\n", request)
		return 404, dto.TodoItemModel{}, e.NotFoundError
	}
	database.UpdateItemByID(request.Id, request.Completed)
	return 200, result, e.Ok
}

func DeleteItemService(id int) (int64, bool, e.ErrorInfo) {
	err := database.DeleteItemByID(id)
	if err == nil {
		return 200, true, e.Ok
	}
	if err == gorm.ErrRecordNotFound {
		logger.Error.Printf("[DeleteItemService] id not found ,request id:%+v\n", id)
		return 404, false, e.NotFoundError
	}
	logger.Error.Printf("[DeleteItemService] delete id error ,request id:%+v\n", id)
	return 500, false, e.InternalServerError
}

func GetCompletedItemsService() {

}

func GetInCompletedItemsService() {

}
