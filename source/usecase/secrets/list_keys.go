package secrets

import (
	"context"
	"fmt"
	"os"

	"github.com/Shitcode-Swamp/unix-adm-project/source/domain"
)

func (uc *UseCase) ListKeys(ctx context.Context, projectName string, env domain.Env) ([]string, error) {
	project, err := uc.projects.FindByName(ctx, projectName)
	if err != nil {
		return nil, fmt.Errorf("project not found: %w", err)
	}
	path := project.ResolvePath(env)

	pairs, err := uc.envFile.Read(path)
	if err != nil {
		if os.IsNotExist(err) {
			return []string{}, nil
		}
		return nil, err
	}

	keys := make([]string, 0, len(pairs))
	for k := range pairs {
		keys = append(keys, k)
	}
	return keys, nil
}
