package myteacher

import (
	"github.com/gin-gonic/gin"
	"learning2.0/common"
	errno "learning2.0/common/erron"
	"learning2.0/model/model_view"
	"learning2.0/service"
	"learning2.0/util"
	"log"
)

func MyTeacherListApi(c *gin.Context) {
	var req model_view.ListMyteacherReq
	err := c.ShouldBind(&req)
	if err != nil {
		log.Printf("ShouldBind ListMyteacherReq failed, err:%v\n", err)
		common.SendResponse(c, errno.ErrParams, err.Error())
		return
	}
	//校验参数
	if err := checkListParam(&req); err != nil {
		log.Printf("checkParam failed, req:%v, err:%v\n", req, err)
		common.SendResponse(c, errno.NoParams, err.Error())
		return
	}

	list, err := service.GetMyteacherList(c, &req)
	if err != nil {
		common.SendResponse(c, errno.QueryErr, err.Error())
		return
	}
	common.SendResponse(c, errno.OK, list)
	return
}
func checkListParam(res *model_view.ListMyteacherReq) error {
	if res.UserID == 0 {
		return util.BuildErrorInfo("用户id不能为空")
	}
	return nil
}