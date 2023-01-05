package label

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

func EditLabelApi(c *gin.Context) {
	// 1、 接收参数
	var req model_view.Lable
	err := c.ShouldBind(&req)
	if err != nil {
		log.Printf("ShouldBind label failed, err:%v\n", err)
		common.SendResponse(c, errno.ErrParams, err.Error())
		return
	}
	//校验参数
	if err := checkEditParam(&req); err != nil {
		log.Printf("checkParam failed, req:%v, err:%v\n", req, err)
		common.SendResponse(c, errno.NoParams, err.Error())
		return
	}
	lable := &model.Lable{
		ID:    req.ID,
		Title: req.Title,
	}

	err = service.EditLabel(c, lable)
	if err != nil {
		common.SendResponse(c, errno.UpdateErr, err.Error())
		return
	}
	common.SendResponse(c, errno.OK, "标签修改成功")
	return
}

func checkEditParam(res *model_view.Lable) error {
	if res.Title == "" {
		return util.BuildErrorInfo("标签名不能为空")
	}
	if res.ID == 0 {
		return util.BuildErrorInfo("标签Id不能为空")
	}
	return nil
}
