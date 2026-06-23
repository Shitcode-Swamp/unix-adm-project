package projects

import (
	"context"

	"github.com/Shitcode-Swamp/unix-adm-project/source/domain"
)

type ProjectRepo interface {
	FindByName(ctx context.Context, name string) (*domain.Project, error)
	Create(ctx context.Context, p *domain.Project) error
	Delete(ctx context.Context, name string) error
	List(ctx context.Context) ([]domain.Project, error)
}

type UseCase struct {
	projects ProjectRepo
}

func New(projects ProjectRepo) *UseCase {
	return &UseCase{projects: projects}
}
