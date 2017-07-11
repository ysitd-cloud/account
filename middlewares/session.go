package middlewares

import (
	"log"
	"os"
	"net/http"

	"gopkg.in/gin-gonic/gin.v1"
	"github.com/gin-contrib/sessions"

	"github.com/gorilla/context"
	gSession "github.com/gorilla/sessions"

	"github.com/ysitd-cloud/account/setup"
)

const (
	defaultKey  = "github.com/gin-contrib/sessions"
	errorFormat = "[sessions] ERROR! %s\n"
)

func Sessions() gin.HandlerFunc {
	store, err := setup.SetupSessionStore()
	if err != nil {
		panic(err)
	}

	name := os.Getenv("SESSION_NAME")

	return func (c *gin.Context) {
		s := &session{name, c.Request, store, nil, false, c.Writer}
		c.Set(defaultKey, s)
		defer context.Clear(c.Request)
		c.Next()
	}
}

type session struct {
	name    string
	request *http.Request
	store   sessions.Store
	session *gSession.Session
	written bool
	writer  http.ResponseWriter
}

func (s *session) Get(key interface{}) interface{} {
	return s.Session().Values[key]
}

func (s *session) Set(key interface{}, val interface{}) {
	s.Session().Values[key] = val
	s.written = true
}

func (s *session) Delete(key interface{}) {
	delete(s.Session().Values, key)
	s.written = true
}

func (s *session) Clear() {
	for key := range s.Session().Values {
		s.Delete(key)
	}
}

func (s *session) AddFlash(value interface{}, vars ...string) {
	s.Session().AddFlash(value, vars...)
	s.written = true
}

func (s *session) Flashes(vars ...string) []interface{} {
	s.written = true
	return s.Session().Flashes(vars...)
}

func (s *session) Options(options sessions.Options) {
	s.Session().Options = &gSession.Options{
		Path:     options.Path,
		Domain:   options.Domain,
		MaxAge:   options.MaxAge,
		Secure:   options.Secure,
		HttpOnly: options.HttpOnly,
	}
}

func (s *session) Save() error {
	if s.Written() {
		e := s.Session().Save(s.request, s.writer)
		if e == nil {
			s.written = false
		}
		return e
	}
	return nil
}

func (s *session) Session() *gSession.Session {
	if s.session == nil {
		var err error
		s.session, err = s.store.Get(s.request, s.name)
		if err != nil {
			log.Printf(errorFormat, err)
		}
	}
	return s.session
}

func (s *session) Written() bool {
	return s.written
}
