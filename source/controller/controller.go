package controller

import (
	"github.com/gin-gonic/gin"

	authctrl "github.com/Shitcode-Swamp/unix-adm-project/source/controller/auth"
	"github.com/Shitcode-Swamp/unix-adm-project/source/controller/middleware"
	projectsctrl "github.com/Shitcode-Swamp/unix-adm-project/source/controller/projects"
	secretsctrl "github.com/Shitcode-Swamp/unix-adm-project/source/controller/secrets"
)

type AppCtx interface {
	AuthUseCase() authctrl.UseCase
	ProjectsUseCase() projectsctrl.UseCase
	SecretsUseCase() secretsctrl.UseCase
	JWTValidator() middleware.Validator
}

func Setup(r *gin.Engine, ctx AppCtx) {
	api := r.Group("/secrets-service")

	authctrl.New(ctx.AuthUseCase()).Register(api.Group("/auth"))

	protected := api.Group("/")
	protected.Use(middleware.Auth(ctx.JWTValidator()))

	projectsctrl.New(ctx.ProjectsUseCase()).Register(protected.Group("/projects"))
	secretsctrl.New(ctx.SecretsUseCase()).Register(protected.Group("/projects/:name"))
}
