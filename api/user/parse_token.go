package user

import (
	"github.com/gin-gonic/gin"
	"learning2.0/common"
	"learning2.0/common/erron"
	"learning2.0/util"
)

func ParseTokenApi(c *gin.Context) {
	claims, _ := c.Get("claims")
	//类型断言
	user, _ := claims.(util.Claims)
	common.SendResponse(c, errno.OK, user)
	return
}
