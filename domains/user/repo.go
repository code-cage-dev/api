package user

import (
	"context"
	"strconv"

	"github.com/cilloparch/cillop/i18np"
	"github.com/code-cage-dev/api/clients/github"
	"github.com/code-cage-dev/api/pkg/entity"
	"gorm.io/gorm"
)

type Repository interface {
	Login(ctx context.Context, user *github.User) (*Entity, *i18np.Error)
	Get(ctx context.Context, username string) (*Entity, *i18np.Error)
}

type repo struct {
	db *gorm.DB
}

func NewRepo(db *gorm.DB) Repository {
	return &repo{db: db}
}

func (r *repo) Login(ctx context.Context, githubUser *github.User) (*Entity, *i18np.Error) {
	var user Entity
	if err := r.db.Where(fields.GithubID, githubUser.ID).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			user = Entity{
				Base:      entity.DefaultBase(),
				Username:  githubUser.Name,
				GithubID:  strconv.Itoa(githubUser.ID),
				AvatarURL: githubUser.AvatarURL,
				Email:     githubUser.Email,
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

func (r *repo) Get(ctx context.Context, username string) (*Entity, *i18np.Error) {
	var user Entity
	if err := r.db.Where(fields.Username, username).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, i18np.NewError(Msg.NotFound)
		} else {
			return nil, i18np.NewError(Msg.Failed)
		}
	}
	return &user, nil
}
