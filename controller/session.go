package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/shereifsrf/savespent-api/dao"
	"github.com/shereifsrf/savespent-api/util"
)

type SessionType struct {
	DeviceID int64 `json:"device_id" binding:"required"`
}

func SessionPersist(c *gin.Context) {
	ctx := c.Request.Context()
	body := &SessionType{}
	status := http.StatusOK

	err := c.ShouldBind(body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Session count: 0",
			"erroe":   err.Error(),
		})
		return
	}

	_, err = setCount(ctx)
	count := 0
	sessionCount := dao.Get(ctx, fmt.Sprintf(SessionKeyFormat, body.DeviceID))
	if sessionCount != "" {
		count = int(util.GetInt64(sessionCount))
	}
	count += 1
	sessionCount = strconv.Itoa(count)

	err = dao.Set(ctx, fmt.Sprintf(SessionKeyFormat, body.DeviceID), sessionCount)
	if err != nil {
		status = http.StatusInternalServerError
	}

	c.JSON(status, gin.H{
		"message": "Session count: " + sessionCount,
		"error":   err,
	})
}
