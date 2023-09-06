package main

import (
	"github.com/cilloparch/cillop/env"
	"github.com/cilloparch/cillop/i18np"
	"github.com/cilloparch/cillop/validation"
	"github.com/code-cage-dev/api/config"
	"github.com/code-cage-dev/api/server/http"
	"github.com/code-cage-dev/api/service"
)

func main() {
	cnf := config.App{}
	env.Load(&cnf)
	i18n := i18np.New(cnf.I18n.Fallback)
	i18n.Load(cnf.I18n.Dir, cnf.I18n.Locales...)
	valid := validation.New(i18n)
	valid.ConnectCustom()
	valid.RegisterTagName()
	app := service.NewApp(service.Config{
		App: cnf,
	})
	http := http.New(http.Config{
		Config: cnf,
		App:    app,
		Valid:  *valid,
		I18n:   i18n,
	})
	http.Listen()
}
