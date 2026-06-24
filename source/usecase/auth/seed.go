package auth

import (
	"context"
	"time"

	"github.com/Shitcode-Swamp/unix-adm-project/source/domain"
)

func (uc *UseCase) SeedAdmin(ctx context.Context, username, password string) error {
	_, err := uc.users.FindByUsername(ctx, username)
	if err == nil {
		return nil
	}
	hash, err := uc.hasher.Hash(password)
	if err != nil {
		return err
	}
	return uc.users.Create(ctx, &domain.User{
		Username:     username,
		PasswordHash: hash,
		CreatedAt:    time.Now(),
	})
}
