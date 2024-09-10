package behaviourTask

import (
	"TaskManagement/app/helpers"
	repoTask "TaskManagement/app/repository/task"
	"context"
)

type TaskRepository struct {
	ctx context.Context
}

func NewTaskRepository() *TaskRepository {
	return &TaskRepository{
		ctx: context.Background(),
	}
}

func (r *TaskRepository) GetAllTask() ([]repoTask.Task, int, error) {
	resSet, cnt, err := repoTask.GetAllTask(r.ctx)
	if err != nil {
		return nil, 0, err
	}
	return resSet, cnt, nil
}

func (r *TaskRepository) GetTask(id int) (*repoTask.Task, error) {
	resSet, err := repoTask.GetTask(r.ctx, id)
	if err != nil {
		return nil, err
	}
	return resSet, nil
}

func (r *TaskRepository) AddTask(obj repoTask.Task) (interface{}, error) {
	resSet, err := repoTask.AddTask(r.ctx, obj)
	if err != nil {
		return nil, err
	}
	return resSet, nil
}

func (r *TaskRepository) UpdateTask(obj map[string]interface{}) error {
	helpers.RemoveNilKeys(obj)
	if obj["name"] != nil {
		obj["name"] = helpers.Namify(obj["name"])
	}
	err := repoTask.UpdateTask(r.ctx, obj)
	if err != nil {
		return err
	}
	return nil
}

func (r *TaskRepository) DeleteTask(id int) error {
	err := repoTask.DeleteTask(r.ctx, id)
	if err != nil {
		return err
	}
	return nil

}
