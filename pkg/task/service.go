package task

import (
	"context"
)

type Repository interface {
	GetAllTasks(ctx context.Context, planID uint) ([]Task, error)
	GetTask(ctx context.Context, planID, taskID uint) (*Task, error)
	CreateTask(ctx context.Context, planID uint, task Task) (uint, error)
	UpdateTask(ctx context.Context, task Task) error
	DeleteTask(ctx context.Context, taskID uint) error
}

type service struct {
	repository Repository
}

// NewService creates a new task service.
func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetAll(ctx context.Context, planID uint) ([]Task, error) {
	return s.repository.GetAllTasks(ctx, planID)
}

func (s *service) Get(ctx context.Context, planID, taskID uint) (*Task, error) {
	return s.repository.GetTask(ctx, planID, taskID)
}

func (s *service) Create(ctx context.Context, planID uint, task Task) (*Task, error) {
	taskID, err := s.repository.CreateTask(ctx, planID, task)
	if err != nil {
		return nil, err
	}
	return s.repository.GetTask(ctx, planID, taskID)

}

func (s *service) Update(ctx context.Context, planID uint, task Task) (*Task, error) {
	if err := s.repository.UpdateTask(ctx, task); err != nil {
		return nil, err
	}
	return s.repository.GetTask(ctx, planID, task.ID)
}

func (s *service) Delete(ctx context.Context, taskID uint) error {
	return s.repository.DeleteTask(ctx, taskID)
}

func (s *service) Done(ctx context.Context, planID, taskID uint) (*Task, error) {
	return s.setDone(ctx, planID, taskID, true)
}

func (s *service) Undo(ctx context.Context, planID, taskID uint) (*Task, error) {
	return s.setDone(ctx, planID, taskID, false)
}

func (s *service) setDone(ctx context.Context, planID, taskID uint, done bool) (*Task, error) {
	task, err := s.repository.GetTask(ctx, planID, taskID)
	if err != nil {
		return nil, err
	}
	if task.Done == done {
		return task, nil
	}
	task.Done = done
	return s.Update(ctx, planID, *task)
}
