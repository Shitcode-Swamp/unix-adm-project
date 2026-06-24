package secrets

import (
	"context"
	"fmt"

	"github.com/Shitcode-Swamp/unix-adm-project/source/domain"
)

func (uc *UseCase) Delete(ctx context.Context, projectName string, env domain.Env, keys []string) error {
	project, err := uc.projects.FindByName(ctx, projectName)
	if err != nil {
		return fmt.Errorf("project not found: %w", err)
	}
	path := project.ResolvePath(env)

	pairs, err := uc.envFile.Read(path)
	if err != nil {
		return err
	}
	for _, k := range keys {
		delete(pairs, k)
	}
	return uc.envFile.Write(path, pairs)
}
