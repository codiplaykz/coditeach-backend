package dao

import (
	"coditeach/database"
	"coditeach/helpers"
	"coditeach/models"
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
	"github.com/mborders/logmatic"
	"time"
)

type ModuleDAO struct {
	Logger *logmatic.Logger
}

func (m *ModuleDAO) Create(module *models.Module) error {
	err := database.DB.QueryRow(context.Background(),
		"insert into modules (curriculum_id, title, description, created_at) VALUES($1,$2,$3,$4) RETURNING id",
		module.Curriculum_id,
		module.Title,
		module.Description,
		time.Now()).Scan(&module.Id)

	if err != nil {
		m.Logger.Error("Unable to create module.")
		return err
	}

	m.Logger.Info("Module created with curriculum id: %v, title: %s", module.Curriculum_id, module.Title)

	return nil
}

func (m *ModuleDAO) Update(module *models.Module) error {
	err := database.DB.QueryRow(context.Background(),
		"update modules set curriculum_id=$1, title=$2, description=$3 where id=$4 returning id",
		module.Curriculum_id,
		module.Title,
		module.Description,
		module.Id).Scan(&module.Id)

	if err == pgx.ErrNoRows {
		m.Logger.Error("Module not found")
		return err
	}

	if err != nil {
		m.Logger.Error("Unable to update module.")
		return err
	}

	m.Logger.Info("Module updated with curriculum id: %v, title: %s", module.Curriculum_id, module.Title)

	return nil
}

func (m *ModuleDAO) Delete(module *models.Module) error {
	_, err := database.DB.Exec(context.Background(), "delete from curriculum_lessons where block_id=any(select id from blocks where module_id=$1)", module.Id)
	if err != nil {
		fmt.Println("1")
		m.Logger.Error("Unable to delete module with id: %s", module.Id)
		return err
	}
	_, err = database.DB.Exec(context.Background(), "delete from blocks where module_id=$1", module.Id)
	if err != nil {
		fmt.Println("2")
		m.Logger.Error("Unable to delete module with id: %s", module.Id)
		return err
	}
	_, err = database.DB.Exec(context.Background(), "delete from modules where id=$1", module.Id)
	if err != nil {
		fmt.Println("3")
		m.Logger.Error("Unable to delete module with id: %s", module.Id)
		return err
	}

	m.Logger.Info("Module deleted with id: %v", module.Id)

	return nil
}

func (m *ModuleDAO) GetById(module *models.Module) error {
	row := database.DB.QueryRow(context.Background(),
		"select * from modules where id=$1",
		module.Id)

	err := row.Scan(&module.Id, &module.Curriculum_id, &module.Title, &module.Description, &module.Created_at)

	if err == pgx.ErrNoRows {
		m.Logger.Error("Module not found.")
		return err
	}

	if err != nil {
		m.Logger.Error("Unable to get module.")
		return err
	}

	return nil
}

func (m *ModuleDAO) GetAllByCurriculumId(curriculumId int) ([]map[string]interface{}, error) {
	rows, err := database.DB.Query(context.Background(),
		"select * from modules where curriculum_id=$1", curriculumId)

	if err != nil {
		m.Logger.Error("Could not get modules")
		return nil, err
	}

	json := helpers.PgSqlRowsToJson(rows)

	if err == pgx.ErrNoRows {
		m.Logger.Error("Modules not found")
		return nil, err
	}

	if err != nil {
		m.Logger.Error("Unable to get modules")
		return nil, err
	}

	return json, nil
}
