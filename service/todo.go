package service

import (
	"context"

	"github.com/ThomasHamilton2/todo-list-6/db"
	"github.com/ThomasHamilton2/todo-list-6/schema"
)

//Close closes DB
func Close(ctx context.Context) {
	db.Close(ctx)
}

//Insert inserts Todo into DB
func Insert(ctx context.Context, todo *schema.Todo) (int, error) {
	return db.Insert(ctx, todo)
}

//Update updates Todo into DB
func Update(ctx context.Context, todo *schema.Todo) error {
	return db.Update(ctx, todo)
}

//Delete delets Todo from DB
func Delete(ctx context.Context, id int) error {
	return db.Delete(ctx, id)
}

//GetAll gets all Todos from DB
func GetAll(ctx context.Context) ([]schema.Todo, error) {
	return db.GetAll(ctx)
}
