package projects

import (
	"context"
	"os"
	"strings"
)

func (uc *UseCase) ListGitRepoPaths(ctx context.Context) ([]string, error) {
	projects, err := uc.projects.List(ctx)
	if err != nil {
		return nil, err
	}

	var paths []string
	for _, p := range projects {
		dir := strings.Replace(p.Dir, "~/", "/host/", 1)
		if info, err := os.Stat(dir + "/.git"); err == nil && info.IsDir() {
			paths = append(paths, p.Dir)
		}
	}
	return paths, nil
}
