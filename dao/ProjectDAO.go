package dao

import (
	"coditeach/database"
	"coditeach/models"
	"context"
	"github.com/mborders/logmatic"
)

type ProjectDAO struct {
	Logger *logmatic.Logger
}

func (p *ProjectDAO) Create(project *models.Project) error {
	err := database.DB.QueryRow(context.Background(), "insert into projects (name, type, level, tech_components, description, creator_Id, source_Code, block_code, cover_img_url, scheme_img_url) values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) returning id;",
		project.Name, project.Type, project.Level, project.Tech_components, project.Description, 1, project.Source_Code, project.Block_code, project.Cover_img_url, project.Scheme_img_url).Scan(&project.Id)

	if err != nil {
		p.Logger.Error("Unable to create project. %s", err)
		return err
	}

	p.Logger.Info("Project created with id: %v", project.Id)

	return nil
}
