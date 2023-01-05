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

func AddLabelApi(c *gin.Context) {
	// 1、 接收参数
	var req model_view.Lable
	err := c.ShouldBind(&req)
	if err != nil {
		log.Printf("ShouldBind Label failed, err:%v\n", err)
		common.SendResponse(c, errno.ErrParams, err.Error())
		return
	}
	//校验参数
	if err := checkAddParam(&req); err != nil {
		log.Printf("checkParam failed, req:%v, err:%v\n", req, err)
		common.SendResponse(c, errno.NoParams, err.Error())
		return
	}

	//////
	label := &model.Lable{
		Title: req.Title,
	}

	err = service.AddLabel(c, label)
	if err != nil {
		common.SendResponse(c, errno.AddErr, err.Error())
		return
	}
	common.SendResponse(c, errno.OK, "标签添加成功")
}

func checkAddParam(req *model_view.Lable) error {
	if req.Title == "" {
		return util.BuildErrorInfo("标签名不能为空")
	}
	return nil
}
