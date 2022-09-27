package dao

import (
	"coditeach/database"
	"coditeach/models"
	"context"
	"github.com/jackc/pgx/v4"
	"github.com/mborders/logmatic"
)

type SubjectDAO struct {
	Logger *logmatic.Logger
}

func (s *SubjectDAO) Create(subject *models.Subject) error {
	err := database.DB.QueryRow(context.Background(),
		"insert into subjects (teacher_id, name, description) VALUES($1,$2,$3) returning id",
		subject.Teacher_id,
		subject.Name,
		subject.Description).Scan(&subject.Id)

	if err != nil {
		s.Logger.Error("Unable to create subject.")
		return err
	}

	s.Logger.Info("Subject created.")

	return nil
}

func (s *SubjectDAO) Update(subject *models.Subject) error {
	err := database.DB.QueryRow(context.Background(),
		"update subjects set teacher_id=$1, name=$2, description=$3 where id=$4 returning id",
		subject.Teacher_id,
		subject.Name,
		subject.Description,
		subject.Id).Scan(&subject.Id)

	if err == pgx.ErrNoRows {
		s.Logger.Error("Subject not found")
		return err
	}

	if err != nil {
		s.Logger.Error("Unable to update subject.")
		return err
	}

	s.Logger.Info("Updated subject with id: %v", subject.Id)

	return nil
}

func (s *SubjectDAO) Delete(subject *models.Subject) error {
	err := database.DB.QueryRow(context.Background(), "delete from subjects where id=$1 returning id", subject.Id).Scan(&subject.Id)

	if err == pgx.ErrNoRows {
		s.Logger.Error("Subject not found")
		return err
	}

	if err != nil {
		s.Logger.Error("Unable to delete subject with id: %v", subject.Id)
		return err
	}
	s.Logger.Info("Deleted subject with id: %v", subject.Id)

	return nil
}

func (s *SubjectDAO) GetById(subject *models.Subject) error {
	row := database.DB.QueryRow(context.Background(),
		"select * from subjects where id=$1",
		subject.Id)

	err := row.Scan(&subject.Id, &subject.Teacher_id, &subject.Name, &subject.Description)

	if err != nil {
		s.Logger.Error("Unable to get subject with id: %v.", subject.Id)
		return err
	}

	return nil
}
