package teacher

import (
	"github.com/gin-gonic/gin"
	"learning2.0/common"
	"learning2.0/common/erron"
	"learning2.0/model/model_view"
	"learning2.0/service"
	"learning2.0/util"
	"log"
)

func BatchDelTeacherApi(c *gin.Context) {
	// 1、 接收参数
	var req model_view.BatchDelReq
	err := c.ShouldBind(&req)
	if err != nil {
		log.Printf("ShouldBind 参数 failed, err:%v\n", err)
		common.SendResponse(c, errno.ErrParams, err.Error())
		return
	}
	//校验参数
	if err := checkBatchDelParam(&req); err != nil {
		log.Printf("checkParam failed, req:%v, err:%v\n", req, err)
		common.SendResponse(c, errno.NoParams, err.Error())
		return
	}
	//获得要删除的
	ids := req.ID
	rows, err := service.BatchDelTeacher(c, ids)
	if err != nil {
		common.SendResponse(c, errno.DeleteErr, err.Error())
		return
	}
	common.SendResponse(c, errno.OK, rows)
	return

}

func checkBatchDelParam(res *model_view.BatchDelReq) error {
	if len(res.ID) < 1 {
		return util.BuildErrorInfo("id不能为空")
	}
	return nil
}
