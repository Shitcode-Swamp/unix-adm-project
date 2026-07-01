package envfile

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type EnvFile struct{}

func New() *EnvFile { return &EnvFile{} }

func (e *EnvFile) Write(path string, pairs map[string]string) error {
	if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
		return err
	}
	var sb strings.Builder
	for k, v := range pairs {
		sb.WriteString(fmt.Sprintf("%s=%s\n", k, v))
	}
	return os.WriteFile(path, []byte(sb.String()), 0644)
}
