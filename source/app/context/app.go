package appctx

import (
	"fmt"
	"os"
	"strconv"

	authctrl "github.com/Shitcode-Swamp/unix-adm-project/source/controller/auth"
	"github.com/Shitcode-Swamp/unix-adm-project/source/controller/middleware"
	projectsctrl "github.com/Shitcode-Swamp/unix-adm-project/source/controller/projects"
	secretsctrl "github.com/Shitcode-Swamp/unix-adm-project/source/controller/secrets"
	"github.com/Shitcode-Swamp/unix-adm-project/source/repo"
	repoprojects "github.com/Shitcode-Swamp/unix-adm-project/source/repo/projects"
	repousers "github.com/Shitcode-Swamp/unix-adm-project/source/repo/users"
	"github.com/Shitcode-Swamp/unix-adm-project/source/usecase/auth"
	ucprojects "github.com/Shitcode-Swamp/unix-adm-project/source/usecase/projects"
	ucsecrets "github.com/Shitcode-Swamp/unix-adm-project/source/usecase/secrets"
	"github.com/Shitcode-Swamp/unix-adm-project/source/utils/envfile"
	"github.com/Shitcode-Swamp/unix-adm-project/source/utils/hash"
	jwtutil "github.com/Shitcode-Swamp/unix-adm-project/source/utils/jwt"
)

type App struct {
	authUC     *auth.UseCase
	projectsUC *ucprojects.UseCase
	secretsUC  *ucsecrets.UseCase
	jwt        *jwtutil.JWT
}

func New() (*App, error) {
	mongoURI := mustEnv("MONGO_URI")
	mongoDB := envOrDefault("MONGO_DB", "secrets_registry")
	jwtSecret := mustEnv("JWT_SECRET")
	jwtTTL, _ := strconv.Atoi(envOrDefault("JWT_TTL_HOURS", "24"))

	db, err := repo.Connect(mongoURI, mongoDB)
	if err != nil {
		return nil, fmt.Errorf("mongo connect: %w", err)
	}

	usersRepo := repousers.New(db)
	projectsRepo := repoprojects.New(db)
	jwtUtil := jwtutil.New(jwtSecret, jwtTTL)

	return &App{
		authUC:     auth.New(usersRepo, hash.New(), jwtUtil),
		projectsUC: ucprojects.New(projectsRepo),
		secretsUC:  ucsecrets.New(projectsRepo, envfile.New()),
		jwt:        jwtUtil,
	}, nil
}

func (a *App) AuthUseCase() authctrl.UseCase         { return a.authUC }
func (a *App) ProjectsUseCase() projectsctrl.UseCase { return a.projectsUC }
func (a *App) SecretsUseCase() secretsctrl.UseCase   { return a.secretsUC }
func (a *App) JWTValidator() middleware.Validator     { return a.jwt }

func mustEnv(key string) string {
	v := os.Getenv(key)
	if v == "" {
		panic(fmt.Sprintf("required env var %s is not set", key))
	}
	return v
}

func envOrDefault(key, def string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return def
}
