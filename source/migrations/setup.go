package migrations

import (
	"context"

	"github.com/Shitcode-Swamp/unix-adm-project/source/migrator"
	projects "github.com/Shitcode-Swamp/unix-adm-project/source/repo/projects"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func projectsUniqueName(ctx context.Context, db *mongo.Database) error {
	_, err := db.Collection(projects.Collection).Indexes().CreateOne(ctx, mongo.IndexModel{
		Keys:    bson.D{{Key: "name", Value: 1}},
		Options: options.Index().SetUnique(true),
	})
	return err
}

func Setup(m *migrator.Migrator) {
	m.Register("0001_projects_unique_name", projectsUniqueName)
}
