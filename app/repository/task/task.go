package repoTask

import (
	"TaskManagement/app/globals/dbconnections"
	"context"
	"errors"

	"github.com/uptrace/bun"
)

type Task struct {
	bun.BaseModel `bun:"table:Tasks"`
	ID            int    `json:"id" bun:"id,pk,autoincrement"`
	Name          string `json:"name" filter:"Namify" validate:"required|maxLen:125"`
	Description   string `json:"description"`
	AuthorID      int    `json:"author_id" validate:"required|min:1"`
}

func GetAllTask(ctx context.Context) ([]Task, int, error) {
	var task []Task
	count, err := dbconnections.MainDB.NewSelect().Model(&task).ScanAndCount(ctx)
	if err != nil {
		return nil, 0, err
	}
	return task, count, nil
}

func GetTask(ctx context.Context, id int) (*Task, error) {
	task := new(Task)
	err := dbconnections.MainDB.NewSelect().Model(task).Where("id = ?", id).Scan(ctx)
	if err != nil {
		return nil, err
	}
	return task, nil
}

func AddTask(ctx context.Context, obj Task) (interface{}, error) {
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

func UpdateTask(ctx context.Context, obj map[string]interface{}) error {

	res, err := dbconnections.MainDB.NewUpdate().Model(&obj).TableExpr("Tasks").Where("id = ?", obj["id"]).Exec(ctx)
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

func DeleteTask(ctx context.Context, id int) error {

	res, err := dbconnections.MainDB.NewDelete().Table("Tasks").Where("id = ?", id).Exec(ctx)
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
