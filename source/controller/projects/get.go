package projects

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (c *Controller) list(ctx *gin.Context) {
	projects, err := c.uc.List(ctx.Request.Context())
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"projects": projects})
}
