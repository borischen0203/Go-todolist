package handlers

import (
	"net/http"
	"strconv"

	"github.com/borischen0203/Go-todolist/dto"
	"github.com/borischen0203/Go-todolist/errors"
	"github.com/borischen0203/Go-todolist/logger"
	"github.com/borischen0203/Go-todolist/services"
	"github.com/gin-gonic/gin"
)

func CreateItem(c *gin.Context) {
	request := dto.TodoRequest{}
	c.Bind(&request)
	logger.Info.Printf("[CreateItem Handler] request:%+v\n", request)
	statusCode, result, err := services.CreateItemService(request)
	switch statusCode {
	case 200:
		c.JSON(http.StatusOK, result)
	case 500:
		c.JSON(http.StatusInternalServerError, err)
	default:
		c.JSON(http.StatusInternalServerError, errors.InternalServerError)
	}
}

func UpdateItem(c *gin.Context) {
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
}

func GetCompletedItems(c *gin.Context) {
	logger.Info.Print("[GetCompletedItems Handler]")
	statusCode, result, err := services.GetItemsService(true)
	switch statusCode {
	case 200:
		c.JSON(http.StatusOK, result)
	case 500:
		c.JSON(http.StatusInternalServerError, err)
	default:
		c.JSON(http.StatusInternalServerError, errors.InternalServerError)
	}
}

func GetIncompleteItems(c *gin.Context) {
	logger.Info.Print("[GetIncompleteItems Handler]")
	statusCode, result, err := services.GetItemsService(false)
	switch statusCode {
	case 200:
		c.JSON(http.StatusOK, result)
	case 500:
		c.JSON(http.StatusInternalServerError, err)
	default:
		c.JSON(http.StatusInternalServerError, errors.InternalServerError)
	}
}
