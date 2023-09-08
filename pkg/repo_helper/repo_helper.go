package repo_helper

import (
	"github.com/cilloparch/cillop/i18np"
	"gorm.io/gorm"
)

func ParseGormError(err error, failedMsg string, notFoundMsg string) *i18np.Error {
	if err == gorm.ErrRecordNotFound {
		return i18np.NewError(notFoundMsg, i18np.P{
			"Error": err,
		})
	}
	return i18np.NewError(failedMsg, i18np.P{
		"Error": err,
	})
}
