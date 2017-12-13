package grpc

import (
	"github.com/ysitd-cloud/account/pkg/model"
	"github.com/ysitd-cloud/grpc-schema/account/actions"
	"golang.org/x/net/context"

	"github.com/ysitd-cloud/grpc-schema/account/models"
)

func (s *AccountService) ValidateUserPassword(_ context.Context, req *actions.ValidateUserRequest) (*actions.ValidateUserReply, error) {
	username := req.GetUsername()
	user, err := model.LoadUserFromDBWithUsername(s.DB, username)
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
