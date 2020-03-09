package db

import (
	"context"

	"github.com/ThomasHamilton2/todo-list-6/schema"
)

const keyRepository = "Repository"

//Repository is an interface
type Repository interface {
	Close()
	Insert(todo *schema.Todo) (int, error)
	Update(todo *schema.Todo) error
	Delete(id int) error
	GetAll() ([]schema.Todo, error)
}

//SetRepository sets repository
func SetRepository(ctx context.Context, repository Repository) context.Context {
	return context.WithValue(ctx, keyRepository, repository)
}

//Close closes the repository
func Close(ctx context.Context) {
	getRepository(ctx).Close()
}

//Insert adds new Todo to database
func Insert(ctx context.Context, todo *schema.Todo) (int, error) {
	return getRepository(ctx).Insert(todo)
}

//Update updates existing Todo in database
func Update(ctx context.Context, todo *schema.Todo) error {
	return getRepository(ctx).Update(todo)
}

//Delete deletes existing Todo from database
func Delete(ctx context.Context, id int) error {
	return getRepository(ctx).Delete(id)
}

//GetAll gets all Todos from database
func GetAll(ctx context.Context) ([]schema.Todo, error) {
	return getRepository(ctx).GetAll()
}

func getRepository(ctx context.Context) Repository {
	return ctx.Value(keyRepository).(Repository)
}
