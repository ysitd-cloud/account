package grpc

import (
	"strings"

	"github.com/RangelReale/osin"
	"github.com/ysitd-cloud/account/pkg/model"
	"github.com/ysitd-cloud/grpc-schema/account/actions"
	"golang.org/x/net/context"

	"github.com/ysitd-cloud/grpc-schema/account/models"
)

func (s *AccountService) ValidateUserPassword(_ context.Context, req *actions.ValidateUserRequest) (*actions.ValidateUserReply, error) {
	username := req.GetUsername()

	db, err := s.Pool.Acquire()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	user, err := model.LoadUserFromDBWithUsername(db, username)
	if err != nil {
		return nil, err
	}

	reply := &actions.ValidateUserReply{
		Valid: false,
		User:  nil,
	}

	if user == nil {
		return reply, nil
	}

	password := req.GetPassword()
	if !user.ValidatePassword(password) {
		return reply, nil
	}

	reply.User = &models.User{
		Username:    user.Username,
		DisplayName: user.DisplayName,
		AvatarUrl:   user.AvatarUrl,
		Email:       user.Email,
	}

	return reply, nil
}

func (s *AccountService) GetUserInfo(_ context.Context, req *actions.GetUserInfoRequest) (*actions.GetUserInfoReply, error) {
	username := req.GetUsername()

	db, err := s.Pool.Acquire()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	user, err := model.LoadUserFromDBWithUsername(db, username)
	if err != nil {
		return nil, err
	}

	reply := &actions.GetUserInfoReply{
		Exists: false,
		User:   nil,
	}

	if user == nil {
		return reply, nil
	}

	reply.Exists = true
	reply.User = &models.User{
		Username:    user.Username,
		DisplayName: user.DisplayName,
		AvatarUrl:   user.AvatarUrl,
		Email:       user.Email,
	}

	return reply, nil
}

func (s *AccountService) GetTokenInfo(_ context.Context, req *actions.GetTokenInfoRequest) (*actions.GetTokenInfoReply, error) {
	token := req.GetToken()

	oauth := s.getOAuthService()
	defer oauth.Storage.Close()

	reply := &actions.GetTokenInfoReply{
		Exists: false,
		Token:  nil,
	}

	if access, err := oauth.Storage.LoadAccess(token); err == nil {
		issuerId := access.UserData.(string)
		user, err := s.getUser(issuerId)
		if err != nil {
			return nil, err
		}

		client := &models.Service{
			Id: access.Client.GetId(),
		}

		scopes := strings.Split(access.Scope, ",")

		token := &models.Token{
			Issuer:   user,
			Audience: client,
			Scopes:   scopes,
			Expire:   encodeToTimestamp(access.ExpireAt()),
		}

		reply.Token = token
	}

	return reply, nil
}

func (s *AccountService) getUser(username string) (*models.User, error) {
	db, err := s.Pool.Acquire()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	user, err := model.LoadUserFromDBWithUsername(db, username)
	if err != nil {
		return nil, err
	}

	return &models.User{
		Username:    user.Username,
		DisplayName: user.DisplayName,
		AvatarUrl:   user.AvatarUrl,
		Email:       user.Email,
	}, nil
}

func (s *AccountService) getOAuthService() *osin.Server {
	return s.Container.Make("osin.server").(*osin.Server)
}
