package routes

import (
	"net/http"
	"server/pkg/reader"

	"github.com/gin-gonic/gin"
)

// GetTodayVerse retrieves the verse for the current day

func GetTodayVerse(c *gin.Context) {
	todayVerse := reader.GetTodayVerse()
	c.JSON(http.StatusOK, gin.H{
		"body": todayVerse,
	})
}
