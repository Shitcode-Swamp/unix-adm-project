package secrets

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type deleteKeysRequest struct {
	Keys []string `json:"keys" binding:"required"`
}

func (c *Controller) patch(ctx *gin.Context) {
	name := ctx.Param("name")
	env := envParam(ctx)
	if !env.Valid() {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "env must be prod or staging"})
		return
	}
	var req deleteKeysRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := c.uc.Delete(ctx.Request.Context(), name, env, req.Keys); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.Status(http.StatusNoContent)
}
