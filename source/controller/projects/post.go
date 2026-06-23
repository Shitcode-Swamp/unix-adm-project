package projects

import (
	"errors"
	"net/http"

	"github.com/Shitcode-Swamp/unix-adm-project/source/domain"
	ucprojects "github.com/Shitcode-Swamp/unix-adm-project/source/usecase/projects"
	"github.com/gin-gonic/gin"
)

type createRequest struct {
	Name     string                `json:"name" binding:"required"`
	EnvPaths map[domain.Env]string `json:"env_paths" binding:"required"`
}

func (c *Controller) create(ctx *gin.Context) {
	var req createRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := c.uc.Create(ctx.Request.Context(), req.Name, req.EnvPaths); err != nil {
		if errors.Is(err, ucprojects.ErrInvalidPath) {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"name": req.Name})
}
