package public

import (
	"net/http"
	"net/url"

	"code.ysitd.cloud/component/account/pkg/server/public/session"
	"code.ysitd.cloud/component/account/pkg/service"
	"github.com/tonyhhyip/vodka"
)

func (p *Provider) loginForm(c *vodka.Context) {
	q := c.Request.URL.Query()
	next := q.Get("next")
	if next == "" {
		next = "/"
	}
	_, err := p.Manager.LoadSession(c.Request)
	if err != nil && err != http.ErrNoCookie {
		c.Logger().Error(err)
		c.Error(err.Error(), http.StatusBadRequest)
		return
	}

	if err != nil {
		p.Renderer.Render(c, http.StatusOK, "login", nil)
	} else {
		c.Redirect(next, http.StatusFound)
	}
}

func (p *Provider) loginHandle(c *vodka.Context) {
	if err := c.Request.ParseForm(); err != nil {
		c.Logger().Error(err)
		c.Error(err.Error(), http.StatusBadRequest)
	}

	reason := ""
	username := c.PostFormValue("username")
	password := c.PostFormValue("password")

	user, err := p.Service.ValidaUserSignIn(c, username, password)
	if err != nil {
		switch err {
		case service.ErrUserNotExists:
			reason = "not_found"
			break
		case service.ErrIncorrectPassword:
			reason = "not_match"
			break
		default:
			c.Logger().Error(err)
			c.Error(err.Error(), http.StatusInternalServerError)
			return
		}
	}

	next := c.PostFormValue("next")
	if next == "" {
		next = "/"
	}

	if reason == "" {
		if err := p.Manager.WriteSession(c.Response, &session.Session{User: user}); err != nil {
			c.Logger().Error(err)
			c.Error(err.Error(), http.StatusInternalServerError)
		}
	} else {
		redirect := &url.URL{
			Path: "/login",
		}
		query := redirect.Query()
		query.Set("next", next)
		query.Set("error", reason)
		redirect.RawQuery = query.Encode()
		c.Redirect(redirect.String(), http.StatusFound)
	}

}
