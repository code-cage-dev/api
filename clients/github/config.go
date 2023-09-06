package github

type Config struct {
	ClientID     string `env:"GITHUB_CLIENT_ID"`
	ClientSecret string `env:"GITHUB_CLIENT_SECRET"`
	RedirectURL  string `env:"GITHUB_REDIRECT_URL"`
}
