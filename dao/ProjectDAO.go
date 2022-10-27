package dao

import (
	"coditeach/database"
	"coditeach/helpers"
	"coditeach/models"
	"context"
	"github.com/jackc/pgx/v4"
	"github.com/mborders/logmatic"
)

type ProjectDAO struct {
	Logger *logmatic.Logger
}

func (p *ProjectDAO) Create(project *models.Project) error {
	err := database.DB.QueryRow(context.Background(), "insert into projects (name, type, level, tech_components, duration, purchase_box_link, description, creator_Id, source_Code, block_code, cover_img_url, scheme_img_url) values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12) returning id;",
		project.Name, project.Type, project.Level, project.Tech_components, project.Duration, project.Purchase_box_link, project.Description, 1, project.Source_Code, project.Block_code, project.Cover_img_url, project.Scheme_img_url).Scan(&project.Id)

	if err != nil {
		p.Logger.Error("Unable to create project. %s", err)
		return err
	}

	p.Logger.Info("Project created with id: %v", project.Id)

	return nil
}

func (p *ProjectDAO) GetAll() ([]map[string]interface{}, error) {
	rows, err := database.DB.Query(context.Background(),
		"select * from projects")

	if err != nil {
		p.Logger.Error("Could not get projects")
		return nil, err
	}

	json := helpers.PgSqlRowsToJson(rows)

	if err == pgx.ErrNoRows {
		p.Logger.Error("Projects not found")
		return nil, err
	}

	if err != nil {
		p.Logger.Error("Unable to get projects")
		return nil, err
	}

	return json, nil
}
