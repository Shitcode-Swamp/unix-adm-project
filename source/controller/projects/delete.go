package projects

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (c *Controller) delete(ctx *gin.Context) {
	name := ctx.Param("name")
	if err := c.uc.Delete(ctx.Request.Context(), name); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.Status(http.StatusNoContent)
}
