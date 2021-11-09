package plan

import (
	"context"
)

type Repository interface {
	GetAllPlans(ctx context.Context) ([]Plan, error)
	GetPlan(ctx context.Context, planID uint) (*Plan, error)
	CreatePlan(ctx context.Context, plan Plan) (uint, error)
	UpdatePlan(ctx context.Context, plan Plan) error
	DeletePlan(ctx context.Context, planID uint) error
}

type service struct {
	repository Repository
}

// NewService creates a new plan service.
func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetAll(ctx context.Context) ([]Plan, error) {
	return s.repository.GetAllPlans(ctx)
}

func (s *service) Get(ctx context.Context, planID uint) (*Plan, error) {
	return s.repository.GetPlan(ctx, planID)
}

func (s *service) Create(ctx context.Context, plan Plan) (*Plan, error) {
	planID, err := s.repository.CreatePlan(ctx, plan)
	if err != nil {
		return nil, err
	}
	return s.repository.GetPlan(ctx, planID)

}

func (s *service) Update(ctx context.Context, plan Plan) (*Plan, error) {
	if err := s.repository.UpdatePlan(ctx, plan); err != nil {
		return nil, err
	}
	return s.repository.GetPlan(ctx, plan.ID)
}

func (s *service) Delete(ctx context.Context, planID uint) error {
	return s.repository.DeletePlan(ctx, planID)
}
