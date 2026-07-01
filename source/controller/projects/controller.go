package projects

import (
	"context"

	"github.com/Shitcode-Swamp/unix-adm-project/source/domain"
	"github.com/gin-gonic/gin"
)

type UseCase interface {
	Create(ctx context.Context, name, dir string) error
	Delete(ctx context.Context, name string) error
	List(ctx context.Context) ([]domain.Project, error)
	ListGitRepoPaths(ctx context.Context) ([]string, error)
}

type Controller struct {
	uc UseCase
}

func New(uc UseCase) *Controller {
	return &Controller{uc: uc}
}

func (c *Controller) Register(r gin.IRouter) {
	r.GET("", c.list)
	r.GET("/git-repos", c.listGitRepoPaths)
	r.POST("", c.create)
	r.DELETE("/:name", c.delete)
}
