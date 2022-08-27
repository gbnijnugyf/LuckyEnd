package middleware

import (
	"net/http"

	"github.com/pkg/errors"

	"github.com/shawu21/test/common"
	"github.com/shawu21/test/helper"
	"github.com/shawu21/test/model"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func AuthMiddleware(c *gin.Context) {
	token := c.Request.Header.Get("token")
	if token == "" {
		c.JSON(http.StatusForbidden, helper.ApiReturn(common.CodeError, "token不存在", nil))
		c.Abort()
		return
	}
	studentNumber, err := helper.VerifyToken(token)
	if err != nil {
		c.JSON(http.StatusForbidden, helper.ApiReturn(common.CodeExpires, "权限不足", nil))
		c.Abort()
		return
	}
	UserID, err := model.GetUserIDByStudentNumber(studentNumber)
	if err != nil {
		log.Errorf("Invalid student number %+v", errors.WithStack(err))
		c.Abort()
		return
	}
	c.Set("studentNumber", studentNumber)
	c.Set("user_id", UserID)
	c.Next()
}
