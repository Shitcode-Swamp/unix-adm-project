package auth

import (
	"context"

	"github.com/Shitcode-Swamp/unix-adm-project/source/domain"
)

type UserRepo interface {
	FindByUsername(ctx context.Context, username string) (*domain.User, error)
	Create(ctx context.Context, u *domain.User) error
}

type Hasher interface {
	Compare(hash, password string) error
	Hash(password string) (string, error)
}

type Tokenizer interface {
	Sign(username string) (string, error)
}

type UseCase struct {
	users     UserRepo
	hasher    Hasher
	tokenizer Tokenizer
}

func New(users UserRepo, hasher Hasher, tokenizer Tokenizer) *UseCase {
	return &UseCase{users: users, hasher: hasher, tokenizer: tokenizer}
}
