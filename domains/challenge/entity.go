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
	Slug                string          `json:"slug" gorm:"type:varchar(255)"`
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

func (d DifficultyLevel) String() string {
	return string(d)
}

type entityFields struct {
	ID                  string
	CreatedBy           string
	PreferredLanguageID string
	Title               string
	Slug                string
	Description         string
	DifficultyLevel     string
	IsPublic            string
}

var fields = entityFields{
	ID:                  "id",
	CreatedBy:           "created_by",
	PreferredLanguageID: "preferred_language_id",
	Title:               "title",
	Slug:                "slug",
	Description:         "description",
	DifficultyLevel:     "difficulty_level",
	IsPublic:            "is_public",
}

type Filter struct {
	CreatedBy        uuid.UUID         `json:"created_by"`
	LanguageID       uuid.UUID         `json:"language_id"`
	Keyword          string            `json:"keyword"`
	DifficultyLevels []DifficultyLevel `json:"difficulty_levels"`
}
