package teacher

import (
	"github.com/gin-gonic/gin"
	"learning2.0/common"
	"learning2.0/common/erron"
	"learning2.0/model"
	"learning2.0/model/model_view"
	"learning2.0/service"
	"learning2.0/util"
	"log"
	"time"
)

func EditTeacherApi(c *gin.Context) {
	// 1、 接收参数
	var req model_view.Teacher
	err := c.ShouldBind(&req)
	if err != nil {
		log.Printf("ShouldBind teacher failed, err:%v\n", err)
		common.SendResponse(c, errno.ErrParams, err.Error())
		return
	}
	//校验参数
	if err := checkEditParam(&req); err != nil {
		log.Printf("checkParam failed, req:%v, err:%v\n", req, err)
		common.SendResponse(c, errno.NoParams, err.Error())
		return
	}
	teacher := &model.Teacher{
		ID:      req.ID,
		Name:    req.Name,
		Source:  req.Source,
		City:    req.City,
		Created: time.Now().Unix(),
	}

	err = service.EditTeacher(c, teacher)
	if err != nil {
		common.SendResponse(c, errno.UpdateErr, err.Error())
		return
	}
	common.SendResponse(c, errno.OK, "教师修改成功")
	return
}

func checkEditParam(res *model_view.Teacher) error {
	if res.Name == "" {
		return util.BuildErrorInfo("教师姓名不能为空")
	}
	if res.ID == 0 {
		return util.BuildErrorInfo("教师Id不能为空")
	}
	return nil
}
