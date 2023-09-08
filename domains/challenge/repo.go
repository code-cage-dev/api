package challenge

import (
	"context"

	"github.com/cilloparch/cillop/i18np"
	"github.com/cilloparch/cillop/types/list"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Repository interface {
	Filter(ctx context.Context, filter Filter, listConfig list.Config) (*list.Result[*Entity], *i18np.Error)
}

type repo struct {
	db *gorm.DB
}

func NewRepo(db *gorm.DB) Repository {
	return &repo{db: db}
}

func (r *repo) Filter(ctx context.Context, filter Filter, listConfig list.Config) (*list.Result[*Entity], *i18np.Error) {
	var entities []*Entity
	var count int64
	db := r.db.Where(fields.IsPublic, true)
	if filter.CreatedBy != uuid.Nil {
		db = db.Where(fields.CreatedBy, filter.CreatedBy)
	}
	if filter.LanguageID != uuid.Nil {
		db = db.Where(fields.PreferredLanguageID, filter.LanguageID)
	}
	if filter.Keyword != "" {
		db = db.Where("(" + fields.Title + " LIKE ? OR " + fields.Slug + " LIKE ? OR " + fields.Description + " LIKE ?)", "%"+filter.Keyword+"%", "%"+filter.Keyword+"%", "%"+filter.Keyword+"%")
	}
	if len(filter.DifficultyLevels) > 0 {
		db = db.Where(fields.DifficultyLevel + " IN ?", filter.DifficultyLevels)
	}
	db = db.Model(&Entity{})
	if err := db.Count(&count).Error; err != nil {
		return nil, i18np.NewError(Msg.Failed, i18np.P{
			"Error": err,
		})
	}
	if err := db.Limit(int(listConfig.Limit)).Offset(int(listConfig.Offset)).Find(&entities).Error; err != nil {
		return nil, i18np.NewError(Msg.Failed, i18np.P{
			"Error": err,
		})
	}
	return &list.Result[*Entity]{
		Total:         count,
		List:          entities,
		FilteredTotal: int64(len(entities)),
		Page:          listConfig.Offset/listConfig.Limit + 1,
		IsNext:        count > listConfig.Offset+listConfig.Limit,
		IsPrev:        listConfig.Offset > 0 && count > 0,
	}, nil
}
