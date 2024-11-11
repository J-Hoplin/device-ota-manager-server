package api

import (
	"github.com/gin-gonic/gin"
	"ota-server/api/health"
)

func APIEntryRouter(app *gin.Engine) {
	router := app.Group("/api")
	health.HealthRouter(router)
}
