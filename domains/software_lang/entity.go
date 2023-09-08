package software_lang

import "github.com/code-cage-dev/api/pkg/entity"

type Entity struct {
	entity.Base
	Name        string `json:"name" gorm:"type:varchar(255)"`
	Description string `json:"description" gorm:"type:text"`
	Logo        string `json:"logo" gorm:"type:varchar(255)"`
	Website     string `json:"website" gorm:"type:varchar(255)"`
	Github      string `json:"github" gorm:"type:varchar(255)"`
	Twitter     string `json:"twitter" gorm:"type:varchar(255)"`
}
