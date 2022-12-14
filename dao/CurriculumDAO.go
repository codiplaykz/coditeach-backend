package dao

import (
	"coditeach/database"
	"coditeach/helpers"
	"coditeach/models"
	"context"
	"github.com/jackc/pgx/v4"
	"github.com/mborders/logmatic"
	"time"
)

type CurriculumDAO struct {
	Logger *logmatic.Logger
}

func (c *CurriculumDAO) Create(curriculum *models.Curriculum) error {
	err := database.DB.QueryRow(context.Background(),
		"insert into curriculums (user_id, title, description, created_at) VALUES($1,$2,$3,$4) returning id",
		curriculum.User_id,
		curriculum.Title,
		curriculum.Description,
		time.Now()).Scan(&curriculum.Id)

	if err != nil {
		c.Logger.Error("Unable to create curriculum.")
		return err
	}

	c.Logger.Info("Curriculum created with title: %s, description: %s", curriculum.Title, curriculum.Description)

	return nil
}

func (c *CurriculumDAO) Update(curriculum *models.Curriculum) error {
	err := database.DB.QueryRow(context.Background(),
		"update curriculums set user_id=$1, title=$2, description=$3 where id=$4 returning id",
		curriculum.User_id,
		curriculum.Title,
		curriculum.Description,
		curriculum.Id).Scan(&curriculum.Id)

	if err == pgx.ErrNoRows {
		c.Logger.Error("Curriculum not found")
		return err
	}

	if err != nil {
		c.Logger.Error("Unable to update curriculum.")
		return err
	}

	c.Logger.Info("Curriculum updated with id: %v, title: %s", curriculum.Id, curriculum.Title)

	return nil
}

func (c *CurriculumDAO) Delete(curriculum *models.Curriculum) error {
	_, err := database.DB.Exec(context.Background(), "delete from curriculum_lessons where block_id=any(select id from blocks where module_id=any(select id from modules where curriculum_id=$1)) returning id", curriculum.Id)
	if err != nil {
		c.Logger.Error("Unable to delete curriculum with id: %s", curriculum.Id)
		return err
	}
	_, err = database.DB.Exec(context.Background(), "delete from blocks where module_id=any(select id from modules where curriculum_id=$1) returning id", curriculum.Id)
	if err != nil {
		c.Logger.Error("Unable to delete curriculum with id: %s", curriculum.Id)
		return err
	}
	_, err = database.DB.Exec(context.Background(), "delete from modules where curriculum_id=$1 returning id", curriculum.Id)
	if err != nil {
		c.Logger.Error("Unable to delete curriculum with id: %s", curriculum.Id)
		return err
	}
	_, err = database.DB.Exec(context.Background(), "delete from curriculums where id=$1 returning id", curriculum.Id)
	if err != nil {
		c.Logger.Error("Unable to delete curriculum with id: %s", curriculum.Id)
		return err
	}
	c.Logger.Info("Deleted curriculum with id: %v", curriculum.Id)

	return nil
}

func (c *CurriculumDAO) GetById(curriculumId int) ([]map[string]interface{}, error) {
	rows, err := database.DB.Query(context.Background(),
		"select c.id, concat(u.name, ' ', u.surname) as name, c.title, c.description, c.created_at from curriculums as c inner join users as u on c.user_id=u.id where c.id=$1",
		curriculumId)

	if err != nil {
		c.Logger.Error("Unable to get curriculum with id: %s.", curriculumId)
		return nil, err
	}

	json := helpers.PgSqlRowsToJson(rows)

	return json, nil
}

func (c *CurriculumDAO) GetAll() ([]map[string]interface{}, error) {
	rows, err := database.DB.Query(context.Background(),
		"select c.id, concat(u.name, ' ', u.surname) as name, c.title, c.description, c.created_at from curriculums as c inner join users as u on c.user_id=u.id;")

	if err != nil {
		c.Logger.Error("Could not get curriculums")
		return nil, err
	}

	json := helpers.PgSqlRowsToJson(rows)

	if err == pgx.ErrNoRows {
		c.Logger.Error("Curriculums not found")
		return nil, err
	}

	if err != nil {
		c.Logger.Error("Unable to get curriculums")
		return nil, err
	}

	return json, nil
}
