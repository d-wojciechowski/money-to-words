package v1

import (
	"ammount-in-words/pkg/converters"
	"ammount-in-words/pkg/logger"
	"ammount-in-words/pkg/web"
	"errors"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

var log *zap.SugaredLogger

func init() {
	log = logger.CreateLogger().Sugar()
}

func ConvertToPLN(c *gin.Context) {
	money := c.Param("money")

	result, err := converters.ConvertToWordRepresentation(money)
	if err != nil {
		c.JSON(http.StatusBadRequest, web.ErrorStruct{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
			Url:     c.Request.URL.Path,
			Method:  c.Request.Method,
		})
		c.Error(errors.New(err.Error()))
	}
	log.Infow("Correct response", "input", money, "output", result)
	defer log.Sync()
	c.JSON(http.StatusOK, result)
}
