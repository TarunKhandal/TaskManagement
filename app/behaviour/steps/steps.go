package behaviourSteps

import (
	repoSteps "TaskManagement/app/repository/steps"
	"context"
)

type StepsRepository struct {
	ctx context.Context
}

func NewStepsRepository() *StepsRepository {
	return &StepsRepository{
		ctx: context.Background(),
	}
}

func (r *StepsRepository) GetSteps(id int) ([]repoSteps.Step, int, error) {
	resSet, cnt, err := repoSteps.GetStep(r.ctx, id)
	if err != nil {
		return nil, 0, err
	}
	return resSet, cnt, nil
}

func (r *StepsRepository) GetStep(taskId, stepId int) (interface{}, int, error) {
	resSet, cnt, err := repoSteps.GetSingleStep(r.ctx, taskId, stepId)
	if err != nil {
		return nil, 0, err
	}
	return resSet, cnt, nil
}

func (r *StepsRepository) AddSteps(obj repoSteps.Step) (interface{}, error) {
	resSet, err := repoSteps.AddStep(r.ctx, obj)
	if err != nil {
		return nil, err
	}
	return resSet, nil
}

func (r *StepsRepository) UpdateSteps(obj map[string]interface{}) error {
	err := repoSteps.UpdateStep(r.ctx, obj)
	if err != nil {
		return err
	}
	return nil
}

func (r *StepsRepository) DeleteSteps(id, taskId int) error {
	err := repoSteps.DeleteStep(r.ctx, id, taskId)
	if err != nil {
		return err
	}
	return nil

}
