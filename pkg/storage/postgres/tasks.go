package postgres

import (
	"context"
	"github.com/pkg/errors"
	"gorm.io/gorm"

	"github.com/ksavluk/go-todo-api/pkg/task"
)

type Task struct {
	gorm.Model
	Description string `gorm:"description"`
	Done        bool   `gorm:"done"`
	PlanID      uint   `gorm:"plan_id"`
	Plan        Plan   `gorm:"constraint:OnUpdate:NO ACTION,OnDelete:CASCADE;"`
}

func (s *storage) GetAllTasks(ctx context.Context, planID uint) ([]task.Task, error) {
	var tasks []Task

	if err := s.db.WithContext(ctx).Find(&tasks, Task{PlanID: planID}).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, errors.Wrap(err, "get_tasks")
	}

	result := make([]task.Task, len(tasks))
	for i, t := range tasks {
		result[i] = convertTaskToEntity(t)
	}
	return result, nil
}

func (s storage) GetTask(ctx context.Context, planID, taskID uint) (*task.Task, error) {
	var task Task

	if err := s.db.WithContext(ctx).Where(&Task{PlanID: planID}).First(&task, taskID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, errors.Wrap(err, "get_task")
	}
	result := convertTaskToEntity(task)
	return &result, nil
}

func (s storage) CreateTask(ctx context.Context, planID uint, task task.Task) (uint, error) {
	newTask := Task{
		PlanID:      planID,
		Description: task.Description,
		Done:        task.Done,
	}

	if err := s.db.WithContext(ctx).Create(&newTask).Error; err != nil {
		return 0, errors.Wrap(err, "create_task")
	}
	return newTask.ID, nil
}

func (s storage) UpdateTask(ctx context.Context, task task.Task) error {
	newTask := Task{
		Description: task.Description,
		Done:        task.Done,
	}
	newTask.ID = task.ID

	err := s.db.WithContext(ctx).Model(&newTask).Select("description", "done").Updates(newTask).Error
	return errors.Wrap(err, "update_task")
}

func (s storage) DeleteTask(ctx context.Context, taskID uint) error {
	err := s.db.WithContext(ctx).Unscoped().Delete(&Task{}, taskID).Error
	return errors.Wrap(err, "delete_task")
}

func convertTaskToEntity(t Task) task.Task {
	return task.Task{
		ID:          t.ID,
		Description: t.Description,
		Done:        t.Done,
		Created:     t.CreatedAt,
	}
}
