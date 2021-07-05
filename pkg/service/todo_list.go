package service

import (
	"github.com/fungerouscode/todo-app"
	"github.com/fungerouscode/todo-app/pkg/repository"
)

type TodoListService struct {
	repo *repository.Repository
}

func NewTodoListService(repo *repository.Repository) *TodoListService {
	return &TodoListService{repo: repo}
}

func (s *TodoListService) Create(userId int, list todo.TodoList) (int, error) {
	return s.repo.TodoList.Create(userId, list)
}

func (s *TodoListService) GetAll(userId int) ([]todo.TodoList, error) {
	return s.repo.TodoList.GetAll(userId)
}

func (s *TodoListService) GetById(userId int, listId int) (todo.TodoList, error) {
	return s.repo.TodoList.GetById(userId, listId)
}
