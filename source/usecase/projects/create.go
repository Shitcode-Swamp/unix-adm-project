package projects

import (
	"context"
	"strings"
	"time"

	"github.com/Shitcode-Swamp/unix-adm-project/source/domain"
)

func (uc *UseCase) Create(ctx context.Context, name string, envPaths map[domain.Env]string) error {
	for _, p := range envPaths {
		if !strings.HasPrefix(p, "~/") {
			return ErrInvalidPath
		}
	}
	return uc.projects.Create(ctx, &domain.Project{
		Name:      name,
		EnvPaths:  envPaths,
		CreatedAt: time.Now(),
	})
}
