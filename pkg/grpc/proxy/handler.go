package proxy

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/tonyhhyip/vodka"

	"code.ysitd.cloud/auth/account/pkg/grpc"
	"code.ysitd.cloud/grpc/schema/account/actions"
)

type Handler struct {
	http.Handler
	MetricsHandler http.Handler `inject:"metrics handler"`
	Service        *grpc.AccountService
}

func CreateProxy(service *grpc.AccountService) *Handler {
	return &Handler{
		Service: service,
	}
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if h.Handler == nil {
		h.initHandler()
	}
	h.Handler.ServeHTTP(w, r)
}

func (h *Handler) ConfigMetrics(handler http.Handler) {
	h.MetricsHandler = handler
}

func (h *Handler) initHandler() {
	router := vodka.NewRouter()
	router.GET("/token/:token", h.getTokenInfo)
	router.GET("/user/:username", h.getUserInfo)
	router.POST("/validate", h.validateUserPassword)
	regular := vodka.CastHandlerForHTTP(router.Handler(), nil)

	mux := http.NewServeMux()
	mux.Handle("/metrics", h.MetricsHandler)
	mux.Handle("/", regular)
	h.Handler = mux
}

func (h *Handler) getTokenInfo(c *vodka.Context) {
	reply, err := h.Service.GetTokenInfo(c, &actions.GetTokenInfoRequest{
		Token: c.UserValue("token").(string),
	})

	if err != nil {
		c.Error(err.Error(), http.StatusBadGateway)
		return
	}

	c.JSON(http.StatusOK, reply)
}

func (h *Handler) getUserInfo(c *vodka.Context) {
	reply, err := h.Service.GetUserInfo(c, &actions.GetUserInfoRequest{
		Username: c.UserValue("username").(string),
	})

	if err != nil {
		c.Error(err.Error(), http.StatusBadGateway)
		return
	}

	c.JSON(http.StatusOK, reply)
}

func (h *Handler) validateUserPassword(c *vodka.Context) {
	defer c.Request.Body.Close()
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.Error(err.Error(), http.StatusBadRequest)
		return
	}

	var req actions.ValidateUserRequest
	if err := json.Unmarshal(body, &req); err != nil {
		c.Error(err.Error(), http.StatusBadRequest)
		return
	}

	reply, err := h.Service.ValidateUserPassword(c, &req)

	if err != nil {
		c.Error(err.Error(), http.StatusBadGateway)
		return
	}

	c.JSON(http.StatusOK, reply)
}
