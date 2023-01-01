package routers

import (
	v1 "ammount-in-words/internal/routers/api/v1"
	"ammount-in-words/internal/routers/middleware"
	"ammount-in-words/pkg/logger"
	"github.com/gin-gonic/gin"
)

var Router *gin.Engine

func InitRouter() {
	log := logger.Logger
	log.Infow("Router init procedure: start")
	Router = gin.New()
	Router.Use(middleware.DefaultStructuredLogger())
	Router.Use(gin.Recovery())

	converter := v1.NewConverterController(logger.Logger)

	apiv1 := Router.Group("/api/v1")
	apiv1.Use()
	{
		apiv1.GET("/convert/pln/:money", converter.ConvertToPLN)
	}
	log.Infow("Router init procedure: end")
	log.Sync()
}
