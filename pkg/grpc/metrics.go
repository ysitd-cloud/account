package grpc

const (
	validateUserPassword = "validate_user_password"
	getUser              = "get_user"
	getToken             = "get_token"
)

func (s *AccountService) Init() {
	s.Collector.RegisterRPC(validateUserPassword, []string{"user"})
	s.Collector.RegisterRPC(getUser, []string{"user"})
	s.Collector.RegisterRPC(getToken, []string{"token"})
}
