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

func AddTeacher(c *gin.Context, teacher *model.Teacher) error {
	teacherDb := database.Query.Teacher
	err := teacherDb.WithContext(c).Create(teacher)
	if err != nil {
		log.Printf("teacherDb create failed, err:%v\n", err)
		return util.BuildErrorInfo("teacherDb create failed, err:%v", err)
	}
	return nil
}

func EditTeacher(c *gin.Context, teacher *model.Teacher) error {
	teacherDb := database.Query.Teacher
	err := teacherDb.WithContext(c).Save(teacher)
	if err != nil {
		log.Printf("teacherDb create failed, err:%v\n", err)
		return util.BuildErrorInfo("teacherDb create failed, err:%v", err)
	}
	return nil
}

func BatchDelTeacher(c *gin.Context, ids []int64) (int64, error) {
	teacherDb := database.Query.Teacher
	info, err := teacherDb.WithContext(c).Where(teacherDb.ID.In(ids...)).Delete()
	if err != nil {
		log.Printf("teacherDb delete failed, err:%v\n", err)
		return 0, util.BuildErrorInfo("teacherDb delete failed, err:%v", err)
	}
	return info.RowsAffected, nil
}

func GetListTeacher(c *gin.Context, req *model_view.ListTeacherReq) (list []*model.Teacher, err error) {
	teacherDb := database.Query.Teacher
	pageSize := viper.GetInt("page_size")
	if req.Page <= 0 {
		req.Page = 1
	}
	sql := teacherDb.WithContext(c).Offset((req.Page - 1) * pageSize).Limit(pageSize)
	if req.SortType == 1 {
		sql.Order(teacherDb.ID.Desc())
	}
	list, err = sql.Find()
	if err != nil {
		log.Printf("teacherDb query failed, err:%v\n", err)
		return nil, util.BuildErrorInfo("teacherDb query failed, err:%v", err)
	}
	return list, nil
}

func MyTeacherAddLabel(c *gin.Context, teachersLabels []*model.MyteacherLable) error {
	myteacherLableDb := database.Query.MyteacherLable
	//batchSize每次添加的数量, 会分批次添加到数据库
	err := myteacherLableDb.WithContext(c).CreateInBatches(teachersLabels, 10)
	if err != nil {
		log.Printf("teacherDb create failed, err:%v\n", err)
		return util.BuildErrorInfo("teacherDb create failed, err:%v", err)
	}
	return nil
}

// 判断teacher是否存在
func HaveTeacher(c *gin.Context, id int64) bool {
	teacherDb := database.Query.Teacher
	res, _ := teacherDb.WithContext(c).Select(teacherDb.ID).Where(teacherDb.ID.Eq(id)).First()
	if res == nil {
		return false
	} else {
		return true
	}

}
