package repoSteps

import (
	"TaskManagement/app/globals/dbconnections"
	"context"
	"errors"

	"github.com/uptrace/bun"
)

type Step struct {
	bun.BaseModel `bun:"table:Steps"`
	ID            int    `json:"id" bun:"id,pk,autoincrement"`
	TaskID        int    `json:"task_id" validate:"required|min:1"`
	Step          string `json:"step" validate:"required|maxLen:125"`
	Completed     bool   `json:"completed"`
}

func GetStep(ctx context.Context, id int) ([]Step, int, error) {
	var steps []Step
	count, err := dbconnections.MainDB.NewSelect().Model(&steps).Where("task_id = ?", id).ScanAndCount(ctx)
	if err != nil {
		return nil, 0, err
	}
	return steps, count, nil
}

func GetSingleStep(ctx context.Context, taskId, stepId int) (interface{}, int, error) {
	var steps Step
	count, err := dbconnections.MainDB.NewSelect().Model(&steps).Where("id = ? AND task_id = ?", stepId, taskId).ScanAndCount(ctx)
	if err != nil {
		return nil, 0, err
	}
	return steps, count, nil
}

func AddStep(ctx context.Context, obj Step) (interface{}, error) {
	res, err := dbconnections.MainDB.NewInsert().Model(&obj).Exec(ctx)
	if err != nil {
		return nil, err
	}

	if aff, err := res.RowsAffected(); aff == 0 {
		return nil, err
	}

	lastID, err := dbconnections.MainDB.NewRaw(`SELECT last_insert_rowid()`).Exec(ctx)
	if err != nil {
		return nil, err
	}

	id, err := lastID.LastInsertId()
	if err != nil {
		return nil, err
	}

	return id, nil
}

func UpdateStep(ctx context.Context, obj map[string]interface{}) error {

	res, err := dbconnections.MainDB.NewUpdate().Model(&obj).TableExpr("Steps").Where("id = ? AND task_id = ?", obj["id"], obj["task_id"]).Exec(ctx)
	if err != nil {
		return err
	}

	rwsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rwsAffected == 0 {
		return errors.New("no rows affected")
	}

	return nil
}

func DeleteStep(ctx context.Context, id, taskId int) error {

	res, err := dbconnections.MainDB.NewDelete().Table("Steps").Where("id = ? AND task_id = ?", id, taskId).Exec(ctx)
	if err != nil {
		return err
	}

	rwsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rwsAffected == 0 {
		return errors.New("no rows affected")
	}

	return nil
}
