package services

import (
	"context"
	"errors"
	"time"

	"github.com/dhaaana/todolist-backend/models"
	"github.com/dhaaana/todolist-backend/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type TodoService interface {
	CreateTodo(*models.CreateTodoRequest) (*models.DBTodo, error)
	UpdateTodo(string, *models.UpdateTodo) (*models.DBTodo, error)
	FindTodoById(string) (*models.DBTodo, error)
	FindTodos(page int, limit int) ([]*models.DBTodo, error)
	DeleteTodo(string) error
}

type TodoServiceImpl struct {
	todoCollection *mongo.Collection
	ctx            context.Context
}

func NewTodoService(todoCollection *mongo.Collection, ctx context.Context) TodoService {
	return &TodoServiceImpl{todoCollection, ctx}
}

func (t *TodoServiceImpl) CreateTodo(todo *models.CreateTodoRequest) (*models.DBTodo, error) {
	todo.CreateAt = time.Now()
	todo.UpdatedAt = todo.CreateAt
	res, err := t.todoCollection.InsertOne(t.ctx, todo)

	if err != nil {
		if er, ok := err.(mongo.WriteException); ok && er.WriteErrors[0].Code == 11000 {
			return nil, errors.New("todo with that title already exists")
		}
		return nil, err
	}

	opt := options.Index()
	opt.SetUnique(true)

	index := mongo.IndexModel{Keys: bson.M{"title": 1}, Options: opt}

	if _, err := t.todoCollection.Indexes().CreateOne(t.ctx, index); err != nil {
		return nil, errors.New("could not create index for title")
	}

	var newTodo *models.DBTodo
	query := bson.M{"_id": res.InsertedID}
	if err = t.todoCollection.FindOne(t.ctx, query).Decode(&newTodo); err != nil {
		return nil, err
	}

	return newTodo, nil
}

func (t *TodoServiceImpl) UpdateTodo(id string, data *models.UpdateTodo) (*models.DBTodo, error) {
	doc, err := utils.ToDoc(data)
	if err != nil {
		return nil, err
	}

	obId, _ := primitive.ObjectIDFromHex(id)
	query := bson.D{{Key: "_id", Value: obId}}
	update := bson.D{{Key: "$set", Value: doc}}
	res := t.todoCollection.FindOneAndUpdate(t.ctx, query, update, options.FindOneAndUpdate().SetReturnDocument(1))

	var updatedTodo *models.DBTodo

	if err := res.Decode(&updatedTodo); err != nil {
		return nil, errors.New("no Todo with that Id exists")
	}

	return updatedTodo, nil
}

func (t *TodoServiceImpl) DeleteTodo(id string) error {
	obId, _ := primitive.ObjectIDFromHex(id)
	query := bson.M{"_id": obId}

	res, err := t.todoCollection.DeleteOne(t.ctx, query)
	if err != nil {
		return err
	}

	if res.DeletedCount == 0 {
		return errors.New("no document with that Id exists")
	}

	return nil
}

func (t *TodoServiceImpl) FindTodoById(id string) (*models.DBTodo, error) {
	obId, _ := primitive.ObjectIDFromHex(id)

	query := bson.M{"_id": obId}

	var Todo *models.DBTodo

	if err := t.todoCollection.FindOne(t.ctx, query).Decode(&Todo); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("no document with that Id exists")
		}

		return nil, err
	}

	return Todo, nil
}

func (t *TodoServiceImpl) FindTodos(page int, limit int) ([]*models.DBTodo, error) {
	if page == 0 {
		page = 1
	}

	if limit == 0 {
		limit = 10
	}

	skip := (page - 1) * limit

	opt := options.FindOptions{}
	opt.SetLimit(int64(limit))
	opt.SetSkip(int64(skip))

	query := bson.M{}

	cursor, err := t.todoCollection.Find(t.ctx, query, &opt)
	if err != nil {
		return nil, err
	}

	defer cursor.Close(t.ctx)

	var Todos []*models.DBTodo

	for cursor.Next(t.ctx) {
		Todo := &models.DBTodo{}
		err := cursor.Decode(Todo)

		if err != nil {
			return nil, err
		}

		Todos = append(Todos, Todo)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	if len(Todos) == 0 {
		return []*models.DBTodo{}, nil
	}

	return Todos, nil
}
