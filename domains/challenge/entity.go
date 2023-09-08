package challenge

import (
	"github.com/code-cage-dev/api/pkg/entity"
	"github.com/google/uuid"
)

type Entity struct {
	entity.Base
	CreatedBy           uuid.UUID       `json:"created_by" gorm:"type:uuid"`
	PreferredLanguageID uuid.UUID       `json:"preferred_language_id" gorm:"type:uuid"`
	Title               string          `json:"title" gorm:"type:varchar(255)"`
	Description         string          `json:"description" gorm:"type:text"`
	DifficultyLevel     DifficultyLevel `json:"difficulty_level" gorm:"type:ENUM('easy', 'medium', 'hard', 'master')"`
	IsPublic            *bool           `json:"is_public" gorm:"type:boolean;default:false"`
}

type DifficultyLevel string

const (
	Easy   DifficultyLevel = "easy"
	Medium DifficultyLevel = "medium"
	Hard   DifficultyLevel = "hard"
	Master DifficultyLevel = "master"
)
