package secrets

import (
	"context"

	"github.com/Shitcode-Swamp/unix-adm-project/source/domain"
)

type ProjectRepo interface {
	FindByName(ctx context.Context, name string) (*domain.Project, error)
}

type EnvFile interface {
	Read(path string) (map[string]string, error)
	Write(path string, pairs map[string]string) error
}

type UseCase struct {
	projects ProjectRepo
	envFile  EnvFile
}

func New(projects ProjectRepo, envFile EnvFile) *UseCase {
	return &UseCase{projects: projects, envFile: envFile}
}
