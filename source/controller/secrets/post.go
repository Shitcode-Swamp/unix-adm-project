package secrets

import (
	"net/http"

	"github.com/Shitcode-Swamp/unix-adm-project/source/domain"
	"github.com/gin-gonic/gin"
)

type uploadRequest struct {
	Secrets []domain.SecretInput `json:"secrets" binding:"required"`
}

func (c *Controller) post(ctx *gin.Context) {
	name := ctx.Param("name")
	env := envParam(ctx)
	if !env.Valid() {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "env must be prod or staging"})
		return
	}
	var req uploadRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := c.uc.Upload(ctx.Request.Context(), name, env, req.Secrets); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.Status(http.StatusNoContent)
}
