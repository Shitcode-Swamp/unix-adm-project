package users

import (
	"context"

	"github.com/Shitcode-Swamp/unix-adm-project/source/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repo struct {
	col *mongo.Collection
}

func New(db *mongo.Database) *Repo {
	return &Repo{col: db.Collection("users")}
}

func (r *Repo) FindByUsername(ctx context.Context, username string) (*domain.User, error) {
	var u domain.User
	if err := r.col.FindOne(ctx, bson.M{"username": username}).Decode(&u); err != nil {
		return nil, err
	}
	return &u, nil
}

func (r *Repo) Create(ctx context.Context, u *domain.User) error {
	_, err := r.col.InsertOne(ctx, u)
	return err
}
