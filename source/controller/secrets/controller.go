package secrets

import (
	"context"

	"github.com/Shitcode-Swamp/unix-adm-project/source/domain"
	"github.com/gin-gonic/gin"
)

type UseCase interface {
	Upload(ctx context.Context, projectName string, env domain.Env, inputs []domain.SecretInput) error
	Delete(ctx context.Context, projectName string, env domain.Env, keys []string) error
	ListKeys(ctx context.Context, projectName string, env domain.Env) ([]string, error)
}

type Controller struct {
	uc UseCase
}

func New(uc UseCase) *Controller {
	return &Controller{uc: uc}
}

func (c *Controller) Register(r gin.IRouter) {
	r.POST("/secrets", c.post)
	r.PATCH("/secrets", c.patch)
	r.GET("/keys", c.getKeys)
}

func envParam(ctx *gin.Context) domain.Env {
	if e := ctx.Query("env"); e != "" {
		return domain.Env(e)
	}
	return domain.EnvStaging
}
