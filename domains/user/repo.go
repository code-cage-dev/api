package user

import (
	"context"

	"github.com/cilloparch/cillop/i18np"
	"github.com/code-cage-dev/api/pkg/entity"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Repository interface {
	Login(ctx context.Context, githubID string, username string) (*Entity, *i18np.Error)
	Get(ctx context.Context, id uuid.UUID) (*Entity, *i18np.Error)
}

type repo struct {
	db *gorm.DB
}

func NewRepo(db *gorm.DB) Repository {
	return &repo{db: db}
}

func (r *repo) Login(ctx context.Context, githubID string, username string) (*Entity, *i18np.Error) {
	var user Entity
	if err := r.db.Where(fields.GithubID, githubID).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			user = Entity{
				Base:     entity.DefaultBase(),
				Username: username,
				GithubID: githubID,
			}
			if err := r.db.Create(&user).Error; err != nil {
				return nil, i18np.NewError(Msg.Failed)
			}
		} else {
			return nil, i18np.NewError(Msg.Failed)
		}
	}
	return &user, nil
}

func (r *repo) Get(ctx context.Context, id uuid.UUID) (*Entity, *i18np.Error) {
	var user Entity
	if err := r.db.Where(fields.ID, id).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, i18np.NewError(Msg.NotFound)
		} else {
			return nil, i18np.NewError(Msg.Failed)
		}
	}
	return &user, nil
}
