package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"online_shop/internal/types"
)

func newResponse(c *gin.Context, statusCode int, err error) {
	logrus.Error(err)
	c.AbortWithStatusJSON(statusCode, types.ErrorResponse{Error: err.Error()})
}
