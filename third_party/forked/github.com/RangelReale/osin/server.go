package osin

import (
	"time"
)

// Server is an OAuth2 implementation
type Server struct {
	Config            *ServerConfig     `inject:""`
	Storage           Storage           `inject:"osin storage"`
	AuthorizeTokenGen AuthorizeTokenGen `inject:"osin authorize token gen"`
	AccessTokenGen    AccessTokenGen    `inject:"osin access token gen"`
	Now               func() time.Time  `inject:"osin now func"`
	Logger            Logger            `inject:"osin logger"`
}

// NewServer creates a new server instance
func NewServer(config *ServerConfig, storage Storage) *Server {
	return &Server{
		Config:            config,
		Storage:           storage,
		AuthorizeTokenGen: &AuthorizeTokenGenDefault{},
		AccessTokenGen:    &AccessTokenGenDefault{},
		Now:               time.Now,
		Logger:            &LoggerDefault{},
	}
}

// NewResponse creates a new response for the server
func (s *Server) NewResponse() *Response {
	r := NewResponse(s.Storage)
	r.ErrorStatusCode = s.Config.ErrorStatusCode
	return r
}
