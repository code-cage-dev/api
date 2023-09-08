package challenge

import (
	"context"

	"github.com/cilloparch/cillop/i18np"
	"github.com/cilloparch/cillop/types/list"
	"github.com/code-cage-dev/api/pkg/repo_helper"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Repository interface {
	Filter(ctx context.Context, filter Filter, listConfig list.Config) (*list.Result[*Entity], *i18np.Error)
	Create(ctx context.Context, entity *Entity) *i18np.Error
	Update(ctx context.Context, id uuid.UUID, userID uuid.UUID, entity *Entity) *i18np.Error
	Delete(ctx context.Context, id uuid.UUID) *i18np.Error
	View(ctx context.Context, slug string, userID uuid.UUID) (*Entity, *i18np.Error)
	MarkPublic(ctx context.Context, id uuid.UUID, userID uuid.UUID) *i18np.Error
	MarkPrivate(ctx context.Context, id uuid.UUID, userID uuid.UUID) *i18np.Error
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
		db = db.Where("("+fields.Title+" LIKE ? OR "+fields.Slug+" LIKE ? OR "+fields.Description+" LIKE ?)", "%"+filter.Keyword+"%", "%"+filter.Keyword+"%", "%"+filter.Keyword+"%")
	}
	if len(filter.DifficultyLevels) > 0 {
		db = db.Where(fields.DifficultyLevel+" IN ?", filter.DifficultyLevels)
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

func (r *repo) Create(ctx context.Context, entity *Entity) *i18np.Error {
	if err := r.db.Create(entity).Error; err != nil {
		return r.parseErr(err)
	}
	return nil
}

func (r *repo) Update(ctx context.Context, id uuid.UUID, userID uuid.UUID, entity *Entity) *i18np.Error {
	if err := r.db.Where(fields.CreatedBy+" = ?", userID).Where(fields.ID+" = ?", id).Save(entity).Error; err != nil {
		return r.parseErr(err)
	}
	return nil
}

func (r *repo) Delete(ctx context.Context, id uuid.UUID) *i18np.Error {
	if err := r.db.Delete(&Entity{}, id).Error; err != nil {
		return r.parseErr(err)
	}
	return nil
}

func (r *repo) View(ctx context.Context, slug string, userID uuid.UUID) (*Entity, *i18np.Error) {
	var entity Entity
	if err := r.db.Where(fields.Slug+" = ?", slug).Where("("+fields.IsPublic+" = true OR "+fields.CreatedBy+" = ?)", userID).First(&entity).Error; err != nil {
		return nil, r.parseErr(err)
	}
	return &entity, nil
}

func (r *repo) MarkPublic(ctx context.Context, id uuid.UUID, userID uuid.UUID) *i18np.Error {
	if err := r.db.Model(&Entity{}).Where(fields.CreatedBy+" = ?", userID).Where(fields.ID+" = ?", id).Update(fields.IsPublic, true).Error; err != nil {
		return r.parseErr(err)
	}
	return nil
}

func (r *repo) MarkPrivate(ctx context.Context, id uuid.UUID, userID uuid.UUID) *i18np.Error {
	if err := r.db.Model(&Entity{}).Where(fields.CreatedBy+" = ?", userID).Where(fields.ID+" = ?", id).Update(fields.IsPublic, false).Error; err != nil {
		return r.parseErr(err)
	}
	return nil
}

func (r *repo) parseErr(err error) *i18np.Error {
	return repo_helper.ParseGormError(err, Msg.Failed, Msg.NotFound)
}
