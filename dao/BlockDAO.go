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

type BlockDAO struct {
	Logger *logmatic.Logger
}

func (b *BlockDAO) Create(block *models.Block) error {
	err := database.DB.QueryRow(context.Background(),
		"insert into blocks (module_id, title, description, created_at) VALUES($1,$2,$3,$4) returning id",
		block.Module_id,
		block.Title,
		block.Description,
		time.Now()).Scan(&block.Id)

	if err != nil {
		b.Logger.Error("Unable to create block. %s", err)
		return err
	}

	b.Logger.Info("Block created with module id: %v, title: %s", block.Module_id, block.Title)

	return nil
}

func (b *BlockDAO) Update(block *models.Block) error {
	_, err := database.DB.Exec(context.Background(),
		"update blocks set module_id=$1, title=$2, description=$3 where id=$4",
		block.Module_id,
		block.Title,
		block.Description,
		block.Id)

	if err != nil {
		b.Logger.Error("Unable to update block. %s", err)
		return err
	}

	b.Logger.Info("Block updated with module id: %v, title: %s", block.Module_id, block.Title)

	return nil
}

func (b *BlockDAO) Delete(block *models.Block) error {
	_, err := database.DB.Exec(context.Background(),
		"delete from curriculum_lessons where block_id=$1",
		block.Id)

	if err != nil {
		b.Logger.Error("Unable to delete block. %s", err)
		return err
	}

	_, err = database.DB.Exec(context.Background(),
		"delete from blocks where id=$1",
		block.Id)

	if err != nil {
		b.Logger.Error("Unable to delete block. %s", err)
		return err
	}

	b.Logger.Info("Block deleted with id: %v", block.Id)

	return nil
}

func (b *BlockDAO) GetById(block *models.Block) error {
	row := database.DB.QueryRow(context.Background(),
		"select * from blocks where id=$1",
		block.Id)

	err := row.Scan(&block.Id, &block.Module_id, &block.Title, &block.Description, &block.Created_at)

	if err == pgx.ErrNoRows {
		b.Logger.Error("Block with id %v not found.", block.Id)
		return err
	}

	if err != nil {
		b.Logger.Error("Unable to get block. %s", err)
		return err
	}

	return nil
}

func (b *BlockDAO) GetAllByModuleId(moduleId int) ([]map[string]interface{}, error) {
	rows, err := database.DB.Query(context.Background(),
		"select * from blocks where module_id=$1", moduleId)

	if err != nil {
		b.Logger.Error("Could not get blocks")
		return nil, err
	}

	json := helpers.PgSqlRowsToJson(rows)

	if err == pgx.ErrNoRows {
		b.Logger.Error("Blocks not found")
		return nil, err
	}

	if err != nil {
		b.Logger.Error("Unable to get blocks")
		return nil, err
	}

	return json, nil
}
