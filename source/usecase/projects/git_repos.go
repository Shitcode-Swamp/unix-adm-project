package projects

import (
	"context"
	"io/fs"
	"path/filepath"
	"strings"
)

const hostRoot = "/host"

func (uc *UseCase) ListGitRepoPaths(ctx context.Context) ([]string, error) {
	projects, err := uc.projects.List(ctx)
	if err != nil {
		return nil, err
	}

	registered := make(map[string]struct{}, len(projects))
	for _, p := range projects {
		registered[p.Dir] = struct{}{}
	}

	var paths []string

	err = filepath.WalkDir(hostRoot, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return filepath.SkipDir
		}
		if !d.IsDir() {
			return nil
		}
		if d.Name() == ".git" {
			repoDir := filepath.Dir(path)
			homePath := strings.Replace(repoDir, hostRoot, "~", 1)
			if _, ok := registered[homePath]; !ok {
				paths = append(paths, homePath)
			}
			return filepath.SkipDir
		}
		return nil
	})

	return paths, err
}
