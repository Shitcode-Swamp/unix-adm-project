package appctx

import (
	"context"
	"fmt"

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
	db, err := repo.Connect(Cfg.MongoURI, Cfg.MongoDB)
	if err != nil {
		return nil, fmt.Errorf("mongo connect: %w", err)
	}

	jwtUtil := jwtutil.New(Cfg.JWTSecret, Cfg.JWTTTLHours)
	projectsRepo := repoprojects.New(db)

	return &App{
		authUC:     auth.New(repousers.New(db), hash.New(), jwtUtil),
		projectsUC: ucprojects.New(projectsRepo),
		secretsUC:  ucsecrets.New(projectsRepo, envfile.New()),
		jwt:        jwtUtil,
	}, nil
}

func (a *App) Seed(ctx context.Context) error {
	return a.authUC.SeedAdmin(ctx, Cfg.AdminUsername, Cfg.AdminPassword)
}

func (a *App) AuthUseCase() authctrl.UseCase         { return a.authUC }
func (a *App) ProjectsUseCase() projectsctrl.UseCase { return a.projectsUC }
func (a *App) SecretsUseCase() secretsctrl.UseCase   { return a.secretsUC }
func (a *App) JWTValidator() middleware.Validator     { return a.jwt }
