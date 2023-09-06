package infra

import (
	"github.com/code-cage-dev/api/domains/challenge"
	"github.com/code-cage-dev/api/domains/solution"
	"github.com/code-cage-dev/api/domains/user"
	"gorm.io/gorm"
)

func RunMigrate(db *gorm.DB) {
	db.AutoMigrate(&user.Entity{}, &solution.Entity{}, &challenge.Entity{})
}
