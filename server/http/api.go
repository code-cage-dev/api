package http

import (
	"time"

	"github.com/cilloparch/cillop/result"
	"github.com/code-cage-dev/api/app/command"
	"github.com/code-cage-dev/api/app/query"
	"github.com/gofiber/fiber/v2"
)

func (s *srv) Login(ctx *fiber.Ctx) error {
	cmd := &command.LoginCmd{}
	s.parseBody(ctx, cmd)
	res, err := s.app.Commands.Login.Handle(ctx.Context(), *cmd)
	if err != nil {
		return err
	}
	ctx.Cookie(&fiber.Cookie{
		Name:     "access_token",
		Value:    res.AccessToken,
		Path:     "/",
		Domain:   s.config.Http.Domain,
		Secure:   true,
		HTTPOnly: true,
		Expires:  time.Now().Add(24 * time.Hour),
		SameSite: "Strict",
	})
	return result.SuccessDetail(Messages.Ok, res.User)
}

func (s *srv) CurrentUser(ctx *fiber.Ctx) error {
	token := ctx.Locals("token").(string)
	res, err := s.app.Queries.CurrentUser.Handle(ctx.Context(), query.CurrentUserQuery{
		Token: token,
	})
	if err != nil {
		return err
	}
	return result.SuccessDetail(Messages.Ok, res)
}

func (s *srv) ProfileView(ctx *fiber.Ctx) error {
	query := &query.ProfileViewQuery{}
	s.parseParams(ctx, query)
	res, err := s.app.Queries.ProfileView.Handle(ctx.Context(), *query)
	if err != nil {
		return err
	}
	return result.SuccessDetail(Messages.Ok, res)
}
