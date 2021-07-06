package service

import (
	"github.com/fungerouscode/todo-app"
	"github.com/fungerouscode/todo-app/pkg/repository"
)

type Authorization interface {
	CreateUser(user todo.User) (int, error)
	GenerateToken(username string, password string) (string, error)
	ParseToken(token string) (int, error)
}

type TodoList interface {
	Create(userId int, list todo.TodoList) (int, error)
	GetAll(userId int) ([]todo.TodoList, error)
	GetById(userId int, listId int) (todo.TodoList, error)
	Delete(userId int, listId int) error
	Update(userId int, listId int, input todo.UpdateListInput) error
}

type TodoItem interface {
	Create(userId int, listId int, todo todo.TodoItem) (int, error)
	GetAll(userId int, listId int) ([]todo.TodoItem, error)
	GetById(userId int, itemId int) (todo.TodoItem, error)
	Delete(userId int, itemId int) error
	Update(userId int, itemId int, input todo.UpdateItemInput) error
}

type Service struct {
	Authorization
	TodoList
	TodoItem
}

func NewService(repos *repository.Repository) *Service {
	return &Service{Authorization: NewAuthService(repos), TodoList: NewTodoListService(repos), TodoItem: NewTodoItemService(repos.TodoItem, repos.TodoList)}
}
