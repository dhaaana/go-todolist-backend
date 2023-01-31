package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/dhaaana/todolist-backend/config"
	"github.com/dhaaana/todolist-backend/controllers"
	"github.com/dhaaana/todolist-backend/routes"
	"github.com/dhaaana/todolist-backend/services"
	"github.com/gin-gonic/gin"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	server      *gin.Engine
	ctx         context.Context
	mongoclient *mongo.Client

	todoService         services.TodoService
	TodoController      controllers.TodoController
	todoCollection      *mongo.Collection
	TodoRouteController routes.TodoRouteController
)

func init() {
	config, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("Could not load environment variables", err)
	}

	ctx = context.TODO()

	// Connect to MongoDB
	mongoconn := options.Client().ApplyURI(config.DBUri)
	mongoclient, err := mongo.Connect(ctx, mongoconn)

	if err != nil {
		panic(err)
	}

	if err := mongoclient.Ping(ctx, readpref.Primary()); err != nil {
		panic(err)
	}

	fmt.Println("MongoDB successfully connected...")

	todoCollection = mongoclient.Database("golang_mongodb").Collection("todos")
	todoService = services.NewTodoService(todoCollection, ctx)
	TodoController = controllers.NewTodoController(todoService)
	TodoRouteController = routes.NewTodoControllerRoute(TodoController)

	server = gin.Default()
}

func main() {
	config, err := config.LoadConfig(".")

	if err != nil {
		log.Fatal("Could not load config", err)
	}

	defer mongoclient.Disconnect(ctx)

	startGinServer(config)
}

func startGinServer(config config.Config) {
	router := server.Group("/api")
	router.GET("/healthchecker", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": "server is running"})
	})

	TodoRouteController.TodoRoute(router)
	log.Fatal(server.Run(":" + config.Port))
}
