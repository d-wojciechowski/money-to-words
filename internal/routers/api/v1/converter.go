package v1

import (
	"ammount-in-words/pkg/converters"
	"ammount-in-words/pkg/web"
	"errors"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

type converterController struct {
	logger *zap.SugaredLogger
}

type ConverterController interface {
	ConvertToPLN(c *gin.Context)
}

func NewConverterController(logger *zap.SugaredLogger) *converterController {
	return &converterController{logger: logger}
}

func (cc *converterController) ConvertToPLN(c *gin.Context) {
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
	cc.logger.Infow("Correct response", "input", money, "output", result)
	c.JSON(http.StatusOK, result)
}
