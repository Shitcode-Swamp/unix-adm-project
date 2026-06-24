package projects

import (
	"context"
	"strings"
	"time"

	"github.com/Shitcode-Swamp/unix-adm-project/source/domain"
)

func (uc *UseCase) Create(ctx context.Context, name, dir string) error {
	if !strings.HasPrefix(dir, "~/") {
		return ErrInvalidPath
	}
	return uc.projects.Create(ctx, &domain.Project{
		Name:      name,
		Dir:       dir,
		CreatedAt: time.Now(),
	})
}
