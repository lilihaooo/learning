package user

import (
	"github.com/gin-gonic/gin"
	"learning2.0/common"
	"learning2.0/common/erron"
	"learning2.0/model"
	"learning2.0/model/model_view"
	"learning2.0/service"
	"log"
	"time"
)

func RegisterApi(c *gin.Context) {
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
	user := &model.User{
		Username: req.Username,
		Password: EncryptPassword([]byte(req.Password)),
		Created:  time.Now().Unix(),
	}

	err = service.AddUser(c, user)
	if err != nil {
		common.SendResponse(c, errno.AddErr, err.Error())
		return
	}
	common.SendResponse(c, errno.OK, "注册成功")
	return
}
