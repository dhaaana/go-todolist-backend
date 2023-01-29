package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CreateTodoRequest struct {
	Title     string    `json:"title" bson:"title" binding:"required"`
	Content   string    `json:"content" bson:"content" binding:"required"`
	CreateAt  time.Time `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}

type DBTodo struct {
	Id        primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Title     string             `json:"title,omitempty" bson:"title,omitempty"`
	Content   string             `json:"content,omitempty" bson:"content,omitempty"`
	CreateAt  time.Time          `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt time.Time          `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}

type UpdateTodo struct {
	Id        primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Title     string             `json:"title,omitempty" bson:"title,omitempty"`
	Content   string             `json:"content,omitempty" bson:"content,omitempty"`
	CreateAt  time.Time          `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt time.Time          `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}

// A pointer in Go is a variable that stores the memory address of another variable. When you create a variable in Go, the value is stored in a specific location in memory. A pointer holds the memory address of that location, rather than the value itself.

// You can create a pointer by using the & operator before a variable. For example:

// go
// Copy code
// x := 5
// var p *int = &x
// In this example, x is an integer variable with the value 5. The & operator before x creates a pointer p that stores the memory address of x.

// You can access the value stored in a pointer by using the * operator.

// scss
// Copy code
// fmt.Println(*p) // Output: 5
// Here *p is known as dereferencing the pointer which means to access the value stored in the memory location that the pointer is pointing to.

// Using pointers can be useful when you want to pass a large data structure to a function, or when you want to modify the value of a variable inside a function. Instead of passing the entire data structure, you can pass a pointer to it, which is much more memory efficient.
