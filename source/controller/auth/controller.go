package auth

import (
	"context"

	"github.com/gin-gonic/gin"
)

type UseCase interface {
	Login(ctx context.Context, username, password string) (string, error)
}

type Controller struct {
	uc UseCase
}

func New(uc UseCase) *Controller {
	return &Controller{uc: uc}
}

func (c *Controller) Register(r gin.IRouter) {
	r.POST("/login", c.login)
}
