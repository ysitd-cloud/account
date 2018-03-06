package grpc

import (
	"strings"

	"code.ysitd.cloud/component/account/pkg/model/user"
	"code.ysitd.cloud/grpc/schema/account/actions"
	"code.ysitd.cloud/grpc/schema/account/models"
	"github.com/RangelReale/osin"
	"github.com/prometheus/client_golang/prometheus"
	"golang.org/x/net/context"
)

func (s *AccountService) ValidateUserPassword(_ context.Context, req *actions.ValidateUserRequest) (*actions.ValidateUserReply, error) {
	username := req.GetUsername()

	finish, err := s.Collector.InvokeRPC(validateUserPassword, prometheus.Labels{
		"user": username,
	})
	if err != nil {
		return nil, err
	}

	defer func() {
		finish <- err == nil
		close(finish)
	}()

	instance, err := user.LoadFromDBWithUsername(s.Pool, username)
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

func (s *AccountService) GetUserInfo(_ context.Context, req *actions.GetUserInfoRequest) (reply *actions.GetUserInfoReply, err error) {
	username := req.GetUsername()

	finish, err := s.Collector.InvokeRPC(getUser, prometheus.Labels{
		"user": username,
	})
	if err != nil {
		return
	}

	defer func() {
		finish <- err == nil
		close(finish)
	}()

	instance, err := user.LoadFromDBWithUsername(s.Pool, username)
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

func (s *AccountService) GetTokenInfo(_ context.Context, req *actions.GetTokenInfoRequest) (*actions.GetTokenInfoReply, error) {
	token := req.GetToken()

	finish, err := s.Collector.InvokeRPC(getToken, prometheus.Labels{
		"token": token,
	})
	if err != nil {
		return nil, err
	}

	defer func() {
		finish <- err == nil
		close(finish)
	}()

	oauth := s.getOAuthService()
	defer oauth.Storage.Close()

	reply := &actions.GetTokenInfoReply{
		Exists: false,
		Token:  nil,
	}

	if access, err := oauth.Storage.LoadAccess(token); err == nil {
		issuerID := access.UserData.(string)
		issuer, err := s.getUser(issuerID)
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

func (s *AccountService) getUser(username string) (*models.User, error) {
	instance, err := user.LoadFromDBWithUsername(s.Pool, username)
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
	return s.Container.Make("osin.server").(*osin.Server)
}
