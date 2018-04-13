package service

import (
	"context"
	"time"

	"code.ysitd.cloud/component/account/pkg/modals/token"
	"code.ysitd.cloud/component/account/pkg/modals/user"
)

const extendDuration = 30 * time.Minute

type Service struct {
	UserStore  *user.Store  `inject:""`
	TokenStore *token.Store `inject:""`
}

func (s *Service) ValidaUserSignIn(ctx context.Context, username, password string) (instance *user.User, err error) {
	instance, err = s.GetUserInfo(ctx, username)
	if err != nil {
		return
	}

	if valid := instance.ValidatePassword(password); !valid {
		return nil, IncorrectPasswordError
	}
	return
}

func (s *Service) GetUserInfo(ctx context.Context, username string) (instance *user.User, err error) {
	return s.UserStore.GetByUsername(ctx, username)
}

func (s *Service) GetTokenInfo(ctx context.Context, token string) (t *token.Token, err error) {
	return s.TokenStore.GetToken(ctx, token)
}

func (s *Service) RevokeToken(ctx context.Context, token string) (err error) {
	return s.TokenStore.Revoke(ctx, token)
}

func (s *Service) ExtendToken(ctx context.Context, token string) (err error) {
	return s.TokenStore.ExtendToken(ctx, token, extendDuration)
}
