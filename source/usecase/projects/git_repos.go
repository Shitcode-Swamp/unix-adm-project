package projects

import (
	"context"
	"io/fs"
	"path/filepath"
	"strings"
)

const hostRoot = "/host"

func (uc *UseCase) ListGitRepoPaths(_ context.Context) ([]string, error) {
	var paths []string

	err := filepath.WalkDir(hostRoot, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return filepath.SkipDir
		}
		if !d.IsDir() {
			return nil
		}
		if d.Name() == ".git" {
			repoDir := filepath.Dir(path)
			homePath := strings.Replace(repoDir, hostRoot, "~", 1)
			paths = append(paths, homePath)
			return filepath.SkipDir
		}
		return nil
	})

	return paths, err
}
