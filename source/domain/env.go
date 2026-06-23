package domain

type Env string

const (
	EnvProd    Env = "prod"
	EnvStaging Env = "staging"
)

func (e Env) Valid() bool {
	return e == EnvProd || e == EnvStaging
}
