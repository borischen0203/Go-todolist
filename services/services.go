package services

import (
	"github.com/borischen0203/Go-todolist/database"
	"github.com/borischen0203/Go-todolist/dto"
)

func CreateItemService(request dto.TodoRequest) interface{} {
	todo := dto.TodoItemModel{Description: request.Descripion, Completed: false}
	result := database.AddItem(todo)

	return result
}

func UpdateItemService() {

}

func DeleteItemService() {

}

func GetCompletedItemsService() {

}

func GetInCompletedItemsService() {

}
