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

// 判断myteacher是否存在
func HaveMyTeacher(c *gin.Context, id int64) bool {
	myteacherDb := database.Query.Myteacher
	res, _ := myteacherDb.WithContext(c).Select(myteacherDb.ID).Where(myteacherDb.ID.Eq(id)).First()
	if res == nil {
		return false
	} else {
		return true
	}
}

func GetMyteacherList(c *gin.Context, req *model_view.ListMyteacherReq) (*model_view.ListMyteacherResp, error) {
	var teacherIDs []int64
	myteacherDb := database.Query.Myteacher
	sql := myteacherDb.WithContext(c).Where(myteacherDb.UserID.Eq(req.UserID))

	pageSize := viper.GetInt("page_size")
	if req.Page <= 0 {
		req.Page = 1
	}
	sql = sql.Offset((req.Page - 1) * pageSize).Limit(pageSize)
	if req.SortType == 1 {
		sql = sql.Order(myteacherDb.Created.Desc())
	}
	err := sql.Pluck(myteacherDb.TeacherID, &teacherIDs)
	if err != nil {
		log.Printf("myteacherDb query failed, err:%v\n", err)
		return nil, util.BuildErrorInfo("myteacherDb query failed, err:%v", err)
	}
	teacherDb := database.Query.Teacher
	sql1 := teacherDb.WithContext(c)
	sql1 = sql1.Where(teacherDb.ID.In(teacherIDs...)).Order()

	res, err := sql1.Find()

	m := make(map[int64]*model.Teacher)
	for _, item := range res {
		m[item.ID] = item
	}
	res2 := make([]*model.Teacher, len(res))
	for i, item := range teacherIDs {
		res2[i] = m[item]
	}

	count, err := sql1.Count()
	if err != nil {
		log.Printf("teacherDb query failed, err:%v\n", err)
		return nil, util.BuildErrorInfo("teacherDb query failed, err:%v", err)
	}
	if err != nil {
		log.Printf("teacherDb query failed, err:%v\n", err)
		return nil, util.BuildErrorInfo("teacherDb query failed, err:%v", err)
	}
	list := &model_view.ListMyteacherResp{
		TeacherInfo: res2,
		Total:       count,
	}
	return list, nil
}
