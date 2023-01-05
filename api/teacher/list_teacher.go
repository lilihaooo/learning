package teacher

import (
	"github.com/gin-gonic/gin"
	"learning2.0/common"
	"learning2.0/common/erron"
	"learning2.0/model/model_view"
	"learning2.0/service"
	"log"
)

func ListTeacherApi(c *gin.Context) {
	// 1、 接收参数
	var req model_view.ListTeacherReq
	err := c.ShouldBind(&req)
	if err != nil {
		log.Printf("ShouldBind list_teacherReq failed, err:%v\n", err)
		common.SendResponse(c, errno.ErrParams, err.Error())
		return
	}
	list, err := service.GetListTeacher(c, &req)
	if err != nil {
		common.SendResponse(c, errno.QueryErr, err.Error())
		return
	}
	common.SendResponse(c, errno.OK, list)
	return
}
