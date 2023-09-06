package github

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/cilloparch/cillop/i18np"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
)

type User struct {
	Login     string `json:"login"`
	ID        int    `json:"id"`
	AvatarURL string `json:"avatar_url"`
	Email     string `json:"email"`
	Name      string `json:"name"`
}

type AccessRequest struct {
	Code  string `json:"code"`
	State string `json:"state"`
}

var (
	oauthStateString = "random"
	userApiURL       = "https://api.github.com/user"
)

type Client interface {
	Access(ctx context.Context, req *AccessRequest) (*User, *i18np.Error)
	CurrentUser(ctx context.Context, token string) (*User, *i18np.Error)
}

type client struct {
	config oauth2.Config
}

func NewClient(config Config) Client {
	cnf := oauth2.Config{
		ClientID:     config.ClientID,
		ClientSecret: config.ClientSecret,
		Scopes:       []string{"user"},
		Endpoint:     github.Endpoint,
	}
	return &client{config: cnf}
}

func (c *client) Access(ctx context.Context, req *AccessRequest) (*User, *i18np.Error) {
	if req.State != oauthStateString {
		return nil, i18np.NewError(Msg.StateInvalid, i18np.P{
			"State":        req.State,
			"DefaultState": oauthStateString,
		})
	}
	token, err := c.config.Exchange(ctx, req.Code)
	if err != nil {
		return nil, i18np.NewError(Msg.CodeExchangeFailed, i18np.P{
			"Error": err.Error(),
		})
	}
	client := c.config.Client(ctx, token)
	return c.currentUser(ctx, client)
}

func (c *client) CurrentUser(ctx context.Context, token string) (*User, *i18np.Error) {
	client := c.config.Client(ctx, &oauth2.Token{AccessToken: token})
	return c.currentUser(ctx, client)
}

func (c *client) currentUser(ctx context.Context, httpClient *http.Client) (*User, *i18np.Error) {
	resp, err := httpClient.Get(userApiURL)
	if err != nil {
		return nil, i18np.NewError(Msg.GetUserInfoFailed, i18np.P{
			"Error": err.Error(),
		})
	}
	defer resp.Body.Close()
	var user User
	if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
		return nil, i18np.NewError(Msg.GetUserInfoFailed, i18np.P{
			"Error": err.Error(),
		})
	}
	return &user, nil
}
