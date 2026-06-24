package secrets

import (
	"context"
	"fmt"
	"os"

	"github.com/Shitcode-Swamp/unix-adm-project/source/domain"
)

func (uc *UseCase) Upload(ctx context.Context, projectName string, env domain.Env, inputs []domain.SecretInput) error {
	project, err := uc.projects.FindByName(ctx, projectName)
	if err != nil {
		return fmt.Errorf("project not found: %w", err)
	}
	path := project.ResolvePath(env)

	pairs, err := uc.envFile.Read(path)
	if err != nil && !os.IsNotExist(err) {
		return err
	}
	if pairs == nil {
		pairs = make(map[string]string)
	}

	for _, s := range inputs {
		pairs[s.Key] = s.Value
	}
	return uc.envFile.Write(path, pairs)
}
