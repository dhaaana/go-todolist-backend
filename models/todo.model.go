package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CreateTodoRequest struct {
	Task        string    `json:"task" bson:"task" binding:"required"`
	Description string    `json:"description" bson:"description" binding:"required"`
	Completed   bool      `json:"completed" bson:"completed"`
	CreateAt    time.Time `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt   time.Time `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}

type DBTodo struct {
	Id          primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Task        string             `json:"task,omitempty" bson:"task,omitempty"`
	Description string             `json:"description,omitempty" bson:"description,omitempty"`
	Completed   bool               `json:"completed,omitempty" bson:"completed,omitempty"`
	CreateAt    time.Time          `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt   time.Time          `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}

type UpdateTodo struct {
	Id          primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Task        string             `json:"task,omitempty" bson:"task,omitempty"`
	Description string             `json:"description,omitempty" bson:"description,omitempty"`
	Completed   bool               `json:"completed,omitempty" bson:"completed,omitempty"`
	CreateAt    time.Time          `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt   time.Time          `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}
