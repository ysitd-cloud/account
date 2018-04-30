package grpc

import (
	"context"
	"strings"

	"code.ysitd.cloud/auth/account/pkg/model/user"
	"code.ysitd.cloud/auth/account/third_party/forked/github.com/RangelReale/osin"
	"code.ysitd.cloud/grpc/schema/account/actions"
	"code.ysitd.cloud/grpc/schema/account/models"
	"github.com/prometheus/client_golang/prometheus"
)

func (s *AccountService) ValidateUserPassword(ctx context.Context, req *actions.ValidateUserRequest) (*actions.ValidateUserReply, error) {
	username := req.GetUsername()

	done, err := s.Collector.InvokeRPC(validateUserPassword, prometheus.Labels{
		"user": username,
	})
	if err != nil {
		return nil, err
	}

	defer done(err == nil)

	instance, err := user.LoadFromDBWithUsername(ctx, s.Pool, username)
	if err != nil {
		return nil, err
	}

	reply := &actions.ValidateUserReply{
		Valid: false,
		User:  nil,
	}

	if instance == nil {
		return reply, nil
	}

	password := req.GetPassword()
	if !instance.ValidatePassword(password) {
		return reply, nil
	}

	reply.User = &models.User{
		Username:    instance.Username,
		DisplayName: instance.DisplayName,
		AvatarUrl:   instance.AvatarUrl,
		Email:       instance.Email,
	}

	return reply, nil
}

func (s *AccountService) GetUserInfo(ctx context.Context, req *actions.GetUserInfoRequest) (reply *actions.GetUserInfoReply, err error) {
	username := req.GetUsername()

	done, err := s.Collector.InvokeRPC(getUser, prometheus.Labels{
		"user": username,
	})
	if err != nil {
		return
	}

	defer done(err == nil)

	instance, err := user.LoadFromDBWithUsername(ctx, s.Pool, username)
	if err != nil {
		return
	}

	err = nil

	reply = &actions.GetUserInfoReply{
		Exists: false,
		User:   nil,
	}

	if instance == nil {
		return reply, nil
	}

	reply.Exists = true
	reply.User = &models.User{
		Username:    instance.Username,
		DisplayName: instance.DisplayName,
		AvatarUrl:   instance.AvatarUrl,
		Email:       instance.Email,
	}

	return
}

func (s *AccountService) GetTokenInfo(ctx context.Context, req *actions.GetTokenInfoRequest) (*actions.GetTokenInfoReply, error) {
	token := req.GetToken()

	done, err := s.Collector.InvokeRPC(getToken, prometheus.Labels{
		"token": token,
	})
	if err != nil {
		return nil, err
	}

	defer done(err == nil)

	oauth := s.getOAuthService()
	defer oauth.Storage.Close()

	reply := &actions.GetTokenInfoReply{
		Exists: false,
		Token:  nil,
	}

	if access, err := oauth.Storage.LoadAccess(token); err == nil {
		issuerID := access.UserData.(string)
		issuer, err := s.getUser(ctx, issuerID)
		if err != nil {
			return nil, err
		}

		client := &models.Service{
			Id: access.Client.GetId(),
		}

		scopes := strings.Split(access.Scope, ",")

		token := &models.Token{
			Issuer:   issuer,
			Audience: client,
			Scopes:   scopes,
			Expire:   encodeToTimestamp(access.ExpireAt()),
		}

		reply.Token = token
		reply.Exists = true
	}

	return reply, nil
}

func (s *AccountService) getUser(ctx context.Context, username string) (*models.User, error) {
	instance, err := user.LoadFromDBWithUsername(ctx, s.Pool, username)
	if err != nil {
		return nil, err
	}

	return &models.User{
		Username:    instance.Username,
		DisplayName: instance.DisplayName,
		AvatarUrl:   instance.AvatarUrl,
		Email:       instance.Email,
	}, nil
}

func (s *AccountService) getOAuthService() *osin.Server {
	return s.Server
}
