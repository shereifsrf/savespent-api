package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shereifsrf/savespent-api/dao"
)

func GetCount(ctx *gin.Context) {
	count := dao.Get(ctx.Request.Context(), CountKey)
	if count == "" {
		count = "0"
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "count: " + count,
	})
}
