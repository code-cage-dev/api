package config

import "github.com/code-cage-dev/api/clients/github"

type I18n struct {
	Fallback string   `env:"I18N_FALLBACK_LANGUAGE" envDefault:"en"`
	Dir      string   `env:"I18N_DIR" envDefault:"./src/locales"`
	Locales  []string `env:"I18N_LOCALES" envDefault:"en,tr"`
}

type Cors struct {
	AllowedOrigins   string `env:"CORS_ALLOWED_ORIGINS" envDefault:"*"`
	AllowedMethods   string `env:"CORS_ALLOWED_METHODS" envDefault:"GET,POST,PUT,DELETE,OPTIONS"`
	AllowedHeaders   string `env:"CORS_ALLOWED_HEADERS" envDefault:"*"`
	AllowCredentials bool   `env:"CORS_ALLOW_CREDENTIALS" envDefault:"true"`
}

type HttpServer struct {
	Host string `env:"HTTP_SERVER_HOST" envDefault:"localhost"`
	Port int    `env:"HTTP_SERVER_PORT" envDefault:"3000"`

	Domain string `env:"HTTP_SERVER_DOMAIN" envDefault:"localhost"`
}

type Postgres struct {
	Host     string `env:"POSTGRES_HOST" envDefault:"localhost"`
	Port     string `env:"POSTGRES_PORT" envDefault:"5432"`
	Database string `env:"POSTGRES_DATABASE" envDefault:"postgres"`
	Username string `env:"POSTGRES_USERNAME" envDefault:"postgres"`
	Password string `env:"POSTGRES_PASSWORD" envDefault:"postgres"`
	Migrate  bool   `env:"POSTGRES_MIGRATE" envDefault:"true"`
	Seed     bool   `env:"POSTGRES_SEED" envDefault:"true"`
}

type App struct {
	I18n     I18n
	Http     HttpServer
	Postgres Postgres
	Github   github.Config
	Cors     Cors
}
