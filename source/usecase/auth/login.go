package auth

import (
	"context"
	"errors"
)

var ErrInvalidCredentials = errors.New("invalid credentials")

func (uc *UseCase) Login(ctx context.Context, username, password string) (string, error) {
	u, err := uc.users.FindByUsername(ctx, username)
	if err != nil {
		return "", ErrInvalidCredentials
	}
	if err := uc.hasher.Compare(u.PasswordHash, password); err != nil {
		return "", ErrInvalidCredentials
	}
	return uc.tokenizer.Sign(u.Username)
}
