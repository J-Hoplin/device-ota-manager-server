package health

import (
	"github.com/gin-gonic/gin"
	"ota-server/internal/res"
)

func HealthRouter(parentRouter *gin.RouterGroup) {
	router := parentRouter.Group("/health")

	router.GET("/", func(context *gin.Context) {
		ok := HealthHandler()
		res.GlobalResponseHandler(context, res.SUCCESS, ok)
		return
	})
}
