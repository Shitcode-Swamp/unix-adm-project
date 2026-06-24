package migrator

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type migration struct {
	name string
	fn   func(ctx context.Context, db *mongo.Database) error
}

type Migrator struct {
	db         *mongo.Database
	col        *mongo.Collection
	migrations []migration
}

func New(db *mongo.Database) *Migrator {
	return &Migrator{
		db:  db,
		col: db.Collection("migrations"),
	}
}

func (m *Migrator) Register(name string, fn func(ctx context.Context, db *mongo.Database) error) {
	m.migrations = append(m.migrations, migration{name: name, fn: fn})
}

func (m *Migrator) Run(ctx context.Context) error {
	for _, mg := range m.migrations {
		count, err := m.col.CountDocuments(ctx, bson.M{"name": mg.name})
		if err != nil {
			return fmt.Errorf("check migration %s: %w", mg.name, err)
		}
		if count > 0 {
			continue
		}

		log.Printf("applying migration: %s", mg.name)
		if err := mg.fn(ctx, m.db); err != nil {
			return fmt.Errorf("migration %s: %w", mg.name, err)
		}

		_, err = m.col.InsertOne(ctx, bson.M{"name": mg.name, "applied_at": time.Now()})
		if err != nil {
			return fmt.Errorf("record migration %s: %w", mg.name, err)
		}
		log.Printf("migration applied: %s", mg.name)
	}
	return nil
}
