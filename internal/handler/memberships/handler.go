package memberships

import (
	"github.com/azybk/music-catalog/internal/models/memberships"
	"github.com/gin-gonic/gin"
)

type service interface {
	SignUp(request memberships.SignUpRequest) error
	Login(req memberships.LoginRequest) (string, error)
}

type Handler struct {
	*gin.Engine
	service service
}

func NewHandler(api *gin.Engine, service service) *Handler {
	return &Handler{
		api,
		service,
	}
}

func (h *Handler) RegisterRoute() {
	route := h.Group("/memberships")
	route.POST("/sign-up", h.SignUp)
	route.POST("/login", h.Login)
}
