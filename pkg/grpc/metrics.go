package grpc

import "sync"

const (
	validateUserPassword = "validate_user_password"
	getUser              = "get_user"
	getToken             = "get_token"
)

var once sync.Once

func (s *AccountService) Init() {
	once.Do(func() {
		s.Collector.RegisterRPC(validateUserPassword, []string{"user"})
		s.Collector.RegisterRPC(getUser, []string{"user"})
		s.Collector.RegisterRPC(getToken, []string{"token"})
	})
}
