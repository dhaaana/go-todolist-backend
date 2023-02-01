package controllers

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/dhaaana/todolist-backend/models"
	"github.com/dhaaana/todolist-backend/services"
	"github.com/gin-gonic/gin"
)

type TodoController struct {
	todoService services.TodoService
}

func NewTodoController(todoService services.TodoService) TodoController {
	return TodoController{todoService}
}

func (tc *TodoController) CreateTodo(ctx *gin.Context) {
	var todo *models.CreateTodoRequest

	if err := ctx.ShouldBindJSON(&todo); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	newTodo, err := tc.todoService.CreateTodo(todo)

	if err != nil {
		if strings.Contains(err.Error(), "task already exists") {
			ctx.JSON(http.StatusConflict, gin.H{"status": "fail", "message": err.Error()})
			return
		}

		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"status": "success", "data": newTodo})
}

func (tc *TodoController) UpdateTodo(ctx *gin.Context) {
	todoId := ctx.Param("todoId")

	var todo *models.UpdateTodo
	if err := ctx.ShouldBindJSON(&todo); err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	updatedTodo, err := tc.todoService.UpdateTodo(todoId, todo)
	if err != nil {
		if strings.Contains(err.Error(), "Id exists") {
			ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": err.Error()})
			return
		}
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": updatedTodo})
}

func (tc *TodoController) DeleteTodo(ctx *gin.Context) {
	todoId := ctx.Param("todoId")

	err := tc.todoService.DeleteTodo(todoId)

	if err != nil {
		if strings.Contains(err.Error(), "Id exists") {
			ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": err.Error()})
			return
		}
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}

func (pc *TodoController) FindTodoById(ctx *gin.Context) {
	todoId := ctx.Param("todoId")

	todo, err := pc.todoService.FindTodoById(todoId)

	if err != nil {
		if strings.Contains(err.Error(), "Id exists") {
			ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": err.Error()})
			return
		}
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": todo})
}

func (tc *TodoController) FindTodos(ctx *gin.Context) {
	var page = ctx.DefaultQuery("page", "1")
	var limit = ctx.DefaultQuery("limit", "10")

	intPage, err := strconv.Atoi(page)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	intLimit, err := strconv.Atoi(limit)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	todos, err := tc.todoService.FindTodos(intPage, intLimit)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "results": len(todos), "data": todos})
}
