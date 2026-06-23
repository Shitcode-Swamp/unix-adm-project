package domain

import (
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Project struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Name      string             `bson:"name"`
	EnvPaths  map[Env]string     `bson:"env_paths"`
	CreatedAt time.Time          `bson:"created_at"`
}

func (p *Project) ResolvePath(env Env) (string, bool) {
	raw, ok := p.EnvPaths[env]
	if !ok {
		return "", false
	}
	return strings.Replace(raw, "~/", "/host/", 1), true
}
