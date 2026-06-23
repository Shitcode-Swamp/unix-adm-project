package projects

import "context"

func (uc *UseCase) Delete(ctx context.Context, name string) error {
	return uc.projects.Delete(ctx, name)
}
