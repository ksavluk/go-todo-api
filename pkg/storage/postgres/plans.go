package postgres

import (
	"context"
	"github.com/pkg/errors"
	"gorm.io/gorm"

	"github.com/ksavluk/go-todo-api/pkg/plan"
)

type Plan struct {
	gorm.Model
	Name string `gorm:"name"`
}

func (s *storage) GetAllPlans(ctx context.Context) ([]plan.Plan, error) {
	var plans []Plan

	if err := s.db.WithContext(ctx).Find(&plans).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, errors.Wrap(err, "get_plans")
	}

	result := make([]plan.Plan, len(plans))
	for i, p := range plans {
		result[i] = convertPlanToEntity(p)
	}
	return result, nil
}

func (s storage) GetPlan(ctx context.Context, planID uint) (*plan.Plan, error) {
	var plan Plan

	if err := s.db.WithContext(ctx).First(&plan, planID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, errors.Wrap(err, "get_plan")
	}
	result := convertPlanToEntity(plan)
	return &result, nil
}

func (s storage) CreatePlan(ctx context.Context, plan plan.Plan) (uint, error) {
	newPlan := convertEntityToPlan(plan)

	if err := s.db.WithContext(ctx).Create(&newPlan).Error; err != nil {
		return 0, errors.Wrap(err, "create_plan")
	}
	return newPlan.ID, nil
}

func (s storage) UpdatePlan(ctx context.Context, plan plan.Plan) error {
	newPlan := convertEntityToPlan(plan)

	err := s.db.WithContext(ctx).Model(&newPlan).Updates(newPlan).Error
	return errors.Wrap(err, "update_plan")
}

func (s storage) DeletePlan(ctx context.Context, planID uint) error {
	err := s.db.WithContext(ctx).Unscoped().Delete(&Plan{}, planID).Error
	return errors.Wrap(err, "delete_plan")
}

func convertPlanToEntity(p Plan) plan.Plan {
	return plan.Plan{
		ID:      p.ID,
		Name:    p.Name,
		Created: p.CreatedAt,
	}
}

func convertEntityToPlan(p plan.Plan) Plan {
	plan := Plan{
		Name: p.Name,
	}
	plan.ID = p.ID
	return plan
}
