package dao

import (
	"coditeach/database"
	"coditeach/models"
	"context"
	"github.com/mborders/logmatic"
	"time"
)

type CurriculumLessonDAO struct {
	Logger *logmatic.Logger
}

func (c *CurriculumLessonDAO) Create(curriculum_lesson *models.CurriculumLesson) error {
	err := database.DB.QueryRow(context.Background(),
		"insert into curriculum_lessons (block_id, title, description, type, content, created_at) VALUES($1,$2,$3,$4,$5,$6) returning id",
		curriculum_lesson.Block_id,
		curriculum_lesson.Title,
		curriculum_lesson.Description,
		curriculum_lesson.Lesson_type,
		curriculum_lesson.Content,
		time.Now()).Scan(&curriculum_lesson.Id)

	if err != nil {
		c.Logger.Error("Unable to create curriculum lesson.")
		return err
	}

	c.Logger.Info("Curriculum lesson created with title: %s, description: %s", curriculum_lesson.Title, curriculum_lesson.Description)

	return nil
}

func (c *CurriculumLessonDAO) Update(curriculum_lesson *models.CurriculumLesson) error {
	_, err := database.DB.Exec(context.Background(),
		"update curriculum_lessons set block_id=$1, title=$2, description=$3, type=$4, content=$5 where id=$6",
		curriculum_lesson.Block_id,
		curriculum_lesson.Title,
		curriculum_lesson.Description,
		curriculum_lesson.Lesson_type,
		curriculum_lesson.Content,
		curriculum_lesson.Id)

	if err != nil {
		c.Logger.Error("Unable to update curriculum lesson.")
		return err
	}

	c.Logger.Info("Curriculum lesson updated with id: %v, title: %s", curriculum_lesson.Id, curriculum_lesson.Title)

	return nil
}

func (c *CurriculumLessonDAO) Delete(curriculum_lesson *models.CurriculumLesson) error {
	_, err := database.DB.Exec(context.Background(),
		"delete from curriculum_lessons where id=$1",
		curriculum_lesson.Id)

	if err != nil {
		c.Logger.Error("Unable to delete curriculum lesson.")
		return err
	}

	c.Logger.Info("Curriculum lesson deleted with id: %v", curriculum_lesson.Id)

	return nil
}

func (c *CurriculumLessonDAO) GetById(curriculum_lesson *models.CurriculumLesson) error {
	row := database.DB.QueryRow(context.Background(),
		"select * from curriculum_lessons where id=$1",
		curriculum_lesson.Id)

	err := row.Scan(&curriculum_lesson.Id,
		&curriculum_lesson.Block_id,
		&curriculum_lesson.Title,
		&curriculum_lesson.Description,
		&curriculum_lesson.Lesson_type,
		&curriculum_lesson.Content,
		&curriculum_lesson.Created_at)

	if err != nil {
		c.Logger.Error("Unable to get curriculum lesson.")
		return err
	}

	return nil
}
