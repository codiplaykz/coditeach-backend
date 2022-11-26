package dao

import (
	"coditeach/database"
	"coditeach/helpers"
	"coditeach/models"
	"context"
	"github.com/jackc/pgx/v4"
	"github.com/mborders/logmatic"
)

type SchoolAdminDAO struct {
	Logger *logmatic.Logger
}

func (s *SchoolAdminDAO) Create(schoolAdmin *models.SchoolAdmin) error {
	err := database.DB.QueryRow(context.Background(),
		"insert into school_admins (user_id, school_id) VALUES($1,$2) returning id",
		schoolAdmin.User_id,
		schoolAdmin.School_id).Scan(&schoolAdmin.Id)

	if err != nil {
		s.Logger.Error("Unable to create school admin.")
		return err
	}

	s.Logger.Info("School admin created.")

	return nil
}

func (s *SchoolAdminDAO) Update(schoolAdmin *models.SchoolAdmin) error {
	err := database.DB.QueryRow(context.Background(),
		"update school_admins set user_id=$1, school_id=$2 where id=$3 returning id",
		schoolAdmin.User_id,
		schoolAdmin.School_id,
		schoolAdmin.Id).Scan(&schoolAdmin.Id)

	if err == pgx.ErrNoRows {
		s.Logger.Error("school admin not found")
		return err
	}

	if err != nil {
		s.Logger.Error("Unable to update school admin.")
		return err
	}

	s.Logger.Info("Updated school admin with id: %v", schoolAdmin.Id)

	return nil
}

func (s *SchoolAdminDAO) Delete(schoolAdmin *models.SchoolAdmin) error {
	err := database.DB.QueryRow(context.Background(), "delete from school_admins where id=$1 returning id", schoolAdmin.Id).Scan(&schoolAdmin.Id)

	if err == pgx.ErrNoRows {
		s.Logger.Error("School admin not found")
		return err
	}

	if err != nil {
		s.Logger.Error("Unable to delete school admin with id: %v", schoolAdmin.Id)
		return err
	}

	s.Logger.Info("Deleted school admin with id: %v", schoolAdmin.Id)

	return nil
}

func (s *SchoolAdminDAO) GetById(schoolAdmin *models.SchoolAdmin) error {
	err := database.DB.QueryRow(context.Background(),
		"select * from school_admins where id=$1",
		schoolAdmin.Id).Scan(&schoolAdmin.Id, &schoolAdmin.User_id, &schoolAdmin.School_id)

	if err == pgx.ErrNoRows {
		s.Logger.Error("School admin not found")
		return err
	}

	if err != nil {
		s.Logger.Error("Unable to get school admin with id: %v.", schoolAdmin.Id)
		return err
	}

	return nil
}

func (s *SchoolAdminDAO) GetBySchoolId(schoolAdmin *models.SchoolAdmin) ([]map[string]interface{}, error) {
	rows, err := database.DB.Query(context.Background(),
		"select * from school_admins inner join users on school_admins.user_id=users.id where school_id=$1", schoolAdmin.School_id)

	if err != nil {
		s.Logger.Error("Could not get school admins")
		return nil, err
	}

	json := helpers.PgSqlRowsToJson(rows)

	if err == pgx.ErrNoRows {
		s.Logger.Error("School admins not found")
		return nil, err
	}

	if err != nil {
		s.Logger.Error("Unable to get school admins")
		return nil, err
	}

	return json, nil
}
