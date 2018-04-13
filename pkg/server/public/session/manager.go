package session

import "net/http"

const authCookieName = "account.auth"

type Manager struct {
	Coder *Coder `inject:""`
}

func (m *Manager) LoadSession(req *http.Request) (s *Session, err error) {
	cookie, err := req.Cookie(authCookieName)
	if err != nil {
		return
	}

	return m.Coder.ParseCookie(cookie.Value)
}

func (m *Manager) WriteSession(w http.ResponseWriter, s *Session) (err error) {
	token, err := m.Coder.CreateCookie(s)
	if err != nil {
		return
	}
	http.SetCookie(w, &http.Cookie{
		Name:     authCookieName,
		Value:    token,
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
	})
	return
}
