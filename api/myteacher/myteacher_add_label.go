package myteacher

import (
	"github.com/gin-gonic/gin"
	"learning2.0/common"
	"learning2.0/common/erron"
	"learning2.0/model"
	"learning2.0/model/model_view"
	"learning2.0/service"
	"learning2.0/util"
	"log"
)

func MyTeacherAddLabelApi(c *gin.Context) {
	//获取参数
	var req model_view.MyTeacherAddLabelReq
	err := c.ShouldBind(&req)
	if err != nil {
		log.Printf("ShouldBind teacheraddlabelReq failed, err:%v\n", err)
		common.SendResponse(c, errno.ErrParams, err.Error())
		return
	}
	//过滤标签
	filetedLabels, err := service.FilterLabel(c, &req)
	if err != nil {
		log.Printf("FilterLabel failed, err:%v\n", err)
		common.SendResponse(c, errno.ErrParams, err.Error())
		return
	}
	req.LabelID = filetedLabels
	//校验参数
	if err := checkTeacherAddLabelParam(c, &req); err != nil {
		log.Printf("checkParam failed, req:%v, err:%v\n", req, err)
		common.SendResponse(c, errno.NoParams, err.Error())
		return
	}

	//添加
	teachersLabels := make([]*model.MyteacherLable, len(req.LabelID))
	for i, v := range req.LabelID {
		teachersLabels[i] = &model.MyteacherLable{
			MyteacherID: req.MyTeacherID,
			LabelID:     v,
		}
	}
	err = service.MyTeacherAddLabel(c, teachersLabels)
	if err != nil {
		common.SendResponse(c, errno.AddErr, err.Error())
		return
	}
	common.SendResponse(c, errno.OK, "该教师标签添加成功")
	return
}

func checkTeacherAddLabelParam(c *gin.Context, req *model_view.MyTeacherAddLabelReq) error {
	if req.MyTeacherID == 0 {
		return util.BuildErrorInfo("my教师Id不能为空")
	}
	if len(req.LabelID) == 0 {
		return util.BuildErrorInfo("标签Id不能为空")
	}
	//判断教师是否存在
	isExistTeacher := service.HaveMyTeacher(c, req.MyTeacherID)
	if !isExistTeacher {
		return util.BuildErrorInfo("myteacher不存在")
	}
	return nil
}
