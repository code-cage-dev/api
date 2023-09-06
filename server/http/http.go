package http

import (
	"github.com/cilloparch/cillop/helpers/http"
	"github.com/cilloparch/cillop/i18np"
	"github.com/cilloparch/cillop/server"
	"github.com/cilloparch/cillop/validation"
	"github.com/code-cage-dev/api/app"
	"github.com/code-cage-dev/api/config"
	"github.com/code-cage-dev/api/pkg/auth"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

type srv struct {
	config config.App
	app    app.Application
	valid  validation.Validator
	i18n   *i18np.I18n
}

type Config struct {
	Config config.App
	App    app.Application
	Valid  validation.Validator
	I18n   *i18np.I18n
}

func New(cfg Config) server.Server {
	return &srv{
		config: cfg.Config,
		app:    cfg.App,
		valid:  cfg.Valid,
		i18n:   cfg.I18n,
	}
}

func (s *srv) Listen() error {
	http.RunServer(http.Config{
		Host:        s.config.Http.Host,
		Port:        s.config.Http.Port,
		I18n:        s.i18n,
		AcceptLangs: []string{},
		CreateHandler: func(router fiber.Router) fiber.Router {
			router.Post("/login", s.Login)
			router.Get("/current-user", s.cors(), auth.Check(Messages.Unauthorized), s.CurrentUser)
			router.Get("/@:user_name", s.cors(), s.ProfileView)
			return router
		},
	})
	return nil
}

func (s *srv) parseBody(ctx *fiber.Ctx, dto interface{}) {
	http.ParseBody(ctx, s.valid, *s.i18n, dto)
}

func (s *srv) parseParams(ctx *fiber.Ctx, dto interface{}) {
	http.ParseParams(ctx, s.valid, *s.i18n, dto)
}

func (s *srv) parseQuery(ctx *fiber.Ctx, dto interface{}) {
	http.ParseQuery(ctx, s.valid, *s.i18n, dto)
}

func (s *srv) cors() fiber.Handler {
	return cors.New(cors.Config{
		AllowOrigins:     s.config.Cors.AllowedOrigins,
		AllowMethods:     s.config.Cors.AllowedMethods,
		AllowHeaders:     s.config.Cors.AllowedHeaders,
		AllowCredentials: s.config.Cors.AllowCredentials,
	})
}
