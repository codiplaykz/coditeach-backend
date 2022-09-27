package dao

import (
	"coditeach/database"
	"coditeach/models"
	"context"
	"github.com/jackc/pgx/v4"
	"github.com/mborders/logmatic"
)

type ParentDAO struct {
	Logger *logmatic.Logger
}

func (p *ParentDAO) Create(parent *models.Parent) error {
	err := database.DB.QueryRow(context.Background(),
		"insert into parents (user_id, student_id) VALUES($1,$2) returning id",
		parent.User_id,
		parent.Student_id).Scan(&parent.Id)

	if err != nil {
		p.Logger.Error("Unable to create parent.")
		return err
	}

	p.Logger.Info("Parent created.")

	return nil
}

func (p *ParentDAO) Update(parent *models.Parent) error {
	err := database.DB.QueryRow(context.Background(),
		"update parents set user_id=$1, student_id=$2 where id=$3 returning id",
		parent.User_id,
		parent.Student_id,
		parent.Id).Scan(&parent.Id)

	if err == pgx.ErrNoRows {
		p.Logger.Error("Parent not found")
		return err
	}

	if err != nil {
		p.Logger.Error("Unable to update parent.")
		return err
	}

	p.Logger.Info("Updated parent with id: %v", parent.Id)

	return nil
}

func (p *ParentDAO) Delete(parent *models.Parent) error {
	err := database.DB.QueryRow(context.Background(), "delete from parents where id=$1 returning id", parent.Id).Scan(&parent.Id)

	if err == pgx.ErrNoRows {
		p.Logger.Error("Parent not found")
		return err
	}

	if err != nil {
		p.Logger.Error("Unable to delete parent with id: %v", parent.Id)
		return err
	}

	p.Logger.Info("Deleted parent with id: %v", parent.Id)

	return nil
}

func (p *ParentDAO) GetById(parent *models.Parent) error {
	err := database.DB.QueryRow(context.Background(),
		"select * from parents where id=$1",
		parent.Id).Scan(&parent.Id, &parent.User_id, &parent.Student_id)

	if err == pgx.ErrNoRows {
		p.Logger.Error("Student not found")
		return err
	}

	if err != nil {
		p.Logger.Error("Unable to get parent with id: %v.", parent.Id)
		return err
	}

	return nil
}
