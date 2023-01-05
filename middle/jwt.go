package middle

import (
	"github.com/gin-gonic/gin"
	"learning2.0/common"
	"learning2.0/common/erron"
	"learning2.0/service"
	"learning2.0/util"
	"log"
)

func JWT() gin.HandlerFunc {
	return func(context *gin.Context) {

		uri := context.Request.RequestURI
		if uri == "/admin/login" || uri == "/admin/register" {
			context.Next()
			return
		}

		tokenStr := context.GetHeader("token")
		if tokenStr == "" {
			common.SendResponse(context, errno.LoginErr, "token为空")
			context.Abort()
			return
		}
		token, claims, err := util.ParseToken(tokenStr)
		if err != nil || !token.Valid {
			common.SendResponse(context, errno.LoginErr, err.Error())
			context.Abort()
			return
		}

		//更具解析出的user_id获取user信息
		user, err := service.GetUserById(context, claims.UserId)
		if err != nil {
			log.Printf("用户查询失败, err:%v\n", err)
			common.SendResponse(context, errno.QueryErr, err.Error())
			return
		}
		if user.ID == 0 {
			common.SendResponse(context, errno.LoginErr, err.Error())
			context.Abort()
			return
		}
		// 将 claims 中的信息存储在 context 中
		context.Set("claims", *claims)
	}
}
