package session

import "github.com/dgrijalva/jwt-go"

type Coder struct {
	key        []byte
	signMethod jwt.SigningMethod
	Parser     *jwt.Parser
}

func NewTransCoder(key string) *Coder {
	return &Coder{
		key:        []byte(key),
		signMethod: jwt.SigningMethodES512,
		Parser:     new(jwt.Parser),
	}
}

func (s *Coder) getKey(token *jwt.Token) (interface{}, error) {
	return s.key, nil
}

func (s *Coder) CreateCookie(claim jwt.Claims) (string, error) {
	token := jwt.NewWithClaims(s.signMethod, claim)
	return token.SignedString(s.key)
}

func (s *Coder) ParseCookie(tokenString string) (session *Session, err error) {
	token, err := s.Parser.ParseWithClaims(tokenString, new(Session), s.getKey)
	if err != nil {
		return
	}

	if session, ok := token.Claims.(*Session); ok && token.Valid {
		return session, nil
	}
	return
}
