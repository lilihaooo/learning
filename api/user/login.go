package user

import (
	"github.com/gin-gonic/gin"
	"learning2.0/common"
	"learning2.0/common/erron"
	"learning2.0/model/model_view"
	"learning2.0/service"
	"learning2.0/util"
	"log"
)

func LoginApi(c *gin.Context) {
	// 1、 接收参数
	var req model_view.UserReq
	err := c.ShouldBind(&req)
	if err != nil {
		log.Printf("ShouldBind exchange_gift failed, err:%v\n", err)
		common.SendResponse(c, errno.ErrParams, err.Error())
		return
	}
	//校验参数
	if err := checkParam(&req); err != nil {
		log.Printf("checkParam failed, req:%v, err:%v\n", req, err)
		common.SendResponse(c, errno.NoParams, err.Error())
		return
	}
	//通过用户名和密码查询出uesr
	password := EncryptPassword([]byte(req.Password))
	user, err := service.GetUserByUsernameAndPassword(c, req.Username, password)
	if err != nil {
		common.SendResponse(c, errno.VerificationErr, err.Error())
		return
	}
	// 3、返回token
	token := util.GenerateToken(user)
	common.SendResponse(c, errno.OK, token)
	return
}
