package repoAuthor

import (
	"TaskManagement/app/globals/dbconnections"
	"context"
	"errors"

	"github.com/uptrace/bun"
)

type Author struct {
	bun.BaseModel `bun:"table:Authors"`
	ID            int    `json:"id" bun:"id,pk,autoincrement"`
	FirstName     string `json:"first_name" filter:"Namify" validate:"maxLen:25"`
	LastName      string `json:"last_name" filter:"Namify" validate:"maxLen:25"`
	Username      string `json:"username" filter:"TrimMergeWhiteSpaces"  validate:"required|maxLen:25"`
}

func GetAllAuthor(ctx context.Context) ([]Author, int, error) {
	var author []Author
	count, err := dbconnections.MainDB.NewSelect().Model(&author).ScanAndCount(ctx)
	if err != nil {
		return nil, 0, err
	}
	return author, count, nil
}

func GetAuthor(ctx context.Context, id int) (*Author, error) {
	author := new(Author)
	err := dbconnections.MainDB.NewSelect().Model(author).Where("id = ?", id).Scan(ctx)
	if err != nil {
		return nil, err
	}
	return author, nil
}

func AddAuthor(ctx context.Context, obj Author) (interface{}, error) {
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

func UpdateAuthor(ctx context.Context, obj map[string]interface{}) error {

	res, err := dbconnections.MainDB.NewUpdate().Model(&obj).TableExpr("Authors").Where("id = ?", obj["id"]).Exec(ctx)
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

func DeleteAuthor(ctx context.Context, id int) error {

	res, err := dbconnections.MainDB.NewDelete().Table("Authors").Where("id = ?", id).Exec(ctx)
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
