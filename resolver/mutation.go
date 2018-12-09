package resolver

import (
	"context"
	"github.com/XMatrixStudio/BlogReaper/graphql"
	"github.com/kataras/iris/core/errors"
)

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) CreateLoginURL(ctx context.Context, backUrl string) (string, error) {
	if r.Session.GetString("id") != "" {
		return "", errors.New("already_login")
	}
	url, state := r.Service.User.GetLoginURL(backUrl)
	r.Session.Set("state", state)
	return url, nil
}

func (r *mutationResolver) Login(ctx context.Context, code string, state string) (*graphql.User, error) {
	if r.Session.GetString("id") != "" {
		return nil, errors.New("already_login")
	} else if r.Session.GetString("state") != state {
		return nil, errors.New("error_state")
	}
	// TODO
	return nil, nil
}

func (r *mutationResolver) Logout(ctx context.Context) (*graphql.User, error) {
	panic("not implemented")
	// TODO
}