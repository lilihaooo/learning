package service

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"learning2.0/database"
	"learning2.0/model"
	"learning2.0/model/model_view"
	"learning2.0/util"
	"log"
)

func AddLabel(c *gin.Context, label *model.Lable) error {
	labelDb := database.Query.Lable
	err := labelDb.WithContext(c).Create(label)
	if err != nil {
		log.Printf("labelDb create failed, err:%v\n", err)
		return util.BuildErrorInfo("labelDb create failed, err:%v", err)
	}
	return nil
}

func EditLabel(c *gin.Context, label *model.Lable) error {
	labelDb := database.Query.Lable
	err := labelDb.WithContext(c).Save(label)
	if err != nil {
		log.Printf("labelDb create failed, err:%v\n", err)
		return util.BuildErrorInfo("labelDb create failed, err:%v", err)
	}
	return nil
}

func BatchDelLabel(c *gin.Context, ids []int64) (int64, error) {
	labelDb := database.Query.Lable
	info, err := labelDb.WithContext(c).Where(labelDb.ID.In(ids...)).Delete()
	if err != nil {
		log.Printf("labelDb delete failed, err:%v\n", err)
		return 0, util.BuildErrorInfo("labelDb delete failed, err:%v", err)
	}
	return info.RowsAffected, nil
}

func GetListLable(c *gin.Context, req *model_view.ListLabelReq) (list []*model.Lable, err error) {
	labelDb := database.Query.Lable
	pageSize := viper.GetInt("page_size")
	if req.Page <= 0 {
		req.Page = 1
	}
	sql := labelDb.WithContext(c).Offset((req.Page - 1) * pageSize).Limit(pageSize)
	if req.SortType == 1 {
		sql.Order(labelDb.ID.Desc())
	}
	list, err = sql.Find()
	if err != nil {
		log.Printf("labelDb query failed, err:%v\n", err)
		return nil, util.BuildErrorInfo("labelDb query failed, err:%v", err)
	}
	return list, nil
}

func FilterLabel(c *gin.Context, req *model_view.MyTeacherAddLabelReq) ([]int64, error) {
	var ids []int64
	labelDb := database.Query.Lable
	//查询一列值 传地址
	err := labelDb.WithContext(c).Pluck(labelDb.ID, &ids)
	///..todo
	list, err := labelDb.WithContext(c).Where(labelDb.ID.In(req.LabelID...)).Find()

	if err != nil {
		log.Printf("labelDb query failed, err:%v\n", err)
		return nil, util.BuildErrorInfo("labelDb query failed, err:%v", err)
	}
	//交集
	inter := util.Intersect(req.LabelID, ids)
	//获得该myteacher已有的标签id
	MyteacherLableDb := database.Query.MyteacherLable
	var MyteacherExistLabelIds []int64
	err = MyteacherLableDb.WithContext(c).Where(MyteacherLableDb.MyteacherID.Eq(req.MyTeacherID)).Pluck(MyteacherLableDb.LabelID, &MyteacherExistLabelIds)
	if err != nil {
		log.Printf("MyteacherLableDb query failed, err:%v\n", err)
		return nil, util.BuildErrorInfo("MyteacherLableDb query failed, err:%v", err)
	}

	diff := util.Difference(inter, MyteacherExistLabelIds)
	//获取已有条数量 如果大于5就不能继续添加了
	myteachersLabelCount := len(MyteacherExistLabelIds)
	if myteachersLabelCount+len(diff) > viper.GetInt("max_myteacher_label") {
		return nil, util.BuildErrorInfo("标签数量超出, 最多还可以添加%v个标签", viper.GetInt("max_myteacher_label")-myteachersLabelCount)
	}
	return diff, nil
}
