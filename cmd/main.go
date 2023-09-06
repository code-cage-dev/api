package main

import (
	"fmt"

	"github.com/cilloparch/cillop/env"
	"github.com/cilloparch/cillop/i18np"
	"github.com/cilloparch/cillop/validation"
	"github.com/code-cage-dev/api/config"
	"github.com/code-cage-dev/api/pkg/infra"
	"github.com/code-cage-dev/api/server/http"
	"github.com/code-cage-dev/api/service"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	cnf := config.App{}
	env.Load(&cnf)
	i18n := i18np.New(cnf.I18n.Fallback)
	i18n.Load(cnf.I18n.Dir, cnf.I18n.Locales...)
	valid := validation.New(i18n)
	valid.ConnectCustom()
	valid.RegisterTagName()
	postgres := loadPostgres(cnf)
	app := service.NewApp(service.Config{
		App: cnf,
		DB:  postgres,
	})
	http := http.New(http.Config{
		Config: cnf,
		App:    app,
		Valid:  *valid,
		I18n:   i18n,
	})
	http.Listen()
}

func loadPostgres(cnf config.App) *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", cnf.Postgres.Host, cnf.Postgres.Username, cnf.Postgres.Password, cnf.Postgres.Database, cnf.Postgres.Port)
	fmt.Println("Connecting to postgres...")
	db, err := gorm.Open(
		postgres.New(
			postgres.Config{
				DSN:                  dsn,
				PreferSimpleProtocol: true,
			},
		),
	)
	if err != nil {
		panic(err)
	}
	fmt.Println("Connected to postgres")
	if cnf.Postgres.Migrate {
		fmt.Println("Migrating...")
		infra.RunMigrate(db)
		fmt.Println("Migrated")
	}
	if cnf.Postgres.Seed {
		fmt.Println("Seeding...")
		infra.RunSeed(db)
		fmt.Println("Seeded")
	}
	return db
}
