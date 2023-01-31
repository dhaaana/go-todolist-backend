package routes

import (
	"github.com/dhaaana/todolist-backend/controllers"
	"github.com/gin-gonic/gin"
)

type TodoRouteController struct {
	todoController controllers.TodoController
}

func NewTodoControllerRoute(todoController controllers.TodoController) TodoRouteController {
	return TodoRouteController{todoController}
}

func (r *TodoRouteController) TodoRoute(rg *gin.RouterGroup) {
	router := rg.Group("/todos")

	router.GET("/", r.todoController.FindTodos)
	router.GET("/:todoId", r.todoController.FindTodoById)
	router.POST("/", r.todoController.CreateTodo)
	router.PATCH("/:todoId", r.todoController.UpdateTodo)
	router.DELETE("/:todoId", r.todoController.DeleteTodo)
}
