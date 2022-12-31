package routers

import (
	v1 "ammount-in-words/internal/routers/api/v1"
	"ammount-in-words/internal/routers/middleware"
	"github.com/gin-gonic/gin"
)

var Router *gin.Engine

func InitRouter() {
	Router = gin.New()
	Router.Use(middleware.DefaultStructuredLogger())
	Router.Use(gin.Recovery())

	apiv1 := Router.Group("/api/v1")
	apiv1.Use()
	{
		apiv1.GET("/convert/pln/:money", v1.ConvertToPLN)
	}
}
