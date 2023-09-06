package config

type I18n struct {
	Fallback string   `env:"I18N_FALLBACK_LANGUAGE" envDefault:"en"`
	Dir      string   `env:"I18N_DIR" envDefault:"./src/locales"`
	Locales  []string `env:"I18N_LOCALES" envDefault:"en,tr"`
}

type HttpServer struct {
	Host string `env:"HTTP_SERVER_HOST" envDefault:"localhost"`
	Port int    `env:"HTTP_SERVER_PORT" envDefault:"3000"`
}

type Postgres struct {
	Host     string `env:"POSTGRES_HOST" envDefault:"localhost"`
	Port     int    `env:"POSTGRES_PORT" envDefault:"5432"`
	Database string `env:"POSTGRES_DATABASE" envDefault:"postgres"`
	Username string `env:"POSTGRES_USERNAME" envDefault:"postgres"`
	Password string `env:"POSTGRES_PASSWORD" envDefault:"postgres"`
}

type App struct {
	I18n     I18n
	Http     HttpServer
	Postgres Postgres
}
