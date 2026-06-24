package domain

import (
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Project struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Name      string             `bson:"name"`
	Dir       string             `bson:"dir"`
	CreatedAt time.Time          `bson:"created_at"`
}

func (p *Project) ResolvePath(env Env) string {
	dir := strings.Replace(p.Dir, "~/", "/host/", 1)
	switch env {
	case EnvStaging:
		return dir + "/.staging.env"
	default:
		return dir + "/.env"
	}
}
