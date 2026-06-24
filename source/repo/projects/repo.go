package projects

import (
	"context"

	"github.com/Shitcode-Swamp/unix-adm-project/source/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const Collection = "projects"

type Repo struct {
	col *mongo.Collection
}

func New(db *mongo.Database) *Repo {
	return &Repo{col: db.Collection(Collection)}
}

func (r *Repo) FindByName(ctx context.Context, name string) (*domain.Project, error) {
	var p domain.Project
	if err := r.col.FindOne(ctx, bson.M{"name": name}).Decode(&p); err != nil {
		return nil, err
	}
	return &p, nil
}

func (r *Repo) Create(ctx context.Context, p *domain.Project) error {
	_, err := r.col.InsertOne(ctx, p)
	return err
}

func (r *Repo) Delete(ctx context.Context, name string) error {
	_, err := r.col.DeleteOne(ctx, bson.M{"name": name})
	return err
}

func (r *Repo) List(ctx context.Context) ([]domain.Project, error) {
	cursor, err := r.col.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	var list []domain.Project
	if err := cursor.All(ctx, &list); err != nil {
		return nil, err
	}
	return list, nil
}
