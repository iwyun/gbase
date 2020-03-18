package utils

import (
	"github.com/DeanThompson/ginpprof"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"net/http"
)

func HealthCheck(r *gin.RouterGroup) {

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, "success")
	})

}

func GinPProf(r *gin.RouterGroup) {
	rr := r.Group("/debug/pprof")
	ginpprof.WrapGroup(rr)
}

func GinSwagger(r *gin.RouterGroup) {
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
