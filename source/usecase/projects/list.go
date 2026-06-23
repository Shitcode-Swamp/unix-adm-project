package projects

import (
	"context"

	"github.com/Shitcode-Swamp/unix-adm-project/source/domain"
)

func (uc *UseCase) List(ctx context.Context) ([]domain.Project, error) {
	return uc.projects.List(ctx)
}
