package secrets

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (c *Controller) getKeys(ctx *gin.Context) {
	name := ctx.Param("name")
	env := envParam(ctx)
	if !env.Valid() {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "env must be prod or staging"})
		return
	}
	keys, err := c.uc.ListKeys(ctx.Request.Context(), name, env)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"keys": keys})
}
