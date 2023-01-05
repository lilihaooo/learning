package routers

import (
	"github.com/gin-gonic/gin"
	. "learning2.0/api/label"
	. "learning2.0/api/myteacher"
	. "learning2.0/api/teacher"
	. "learning2.0/api/user"
	"learning2.0/middle"
)

func InitRouter(r *gin.Engine) *gin.Engine {
	r.Use(middle.Cors()).Use(middle.JWT())
	//后端系统
	v1Group := r.Group("admin/teacher")
	v2Group := r.Group("admin/label")
	v3Group := r.Group("myteacher")

	{
		// 登录
		// localhost:8888/xxx
		r.POST("/login", LoginApi)
		r.POST("/register", RegisterApi)
		r.GET("/parsetoken", ParseTokenApi)
		//退出登陆
		//r.GET("/parsetoken", ParseTokenApi)

		// localhost:8888/admin/teacher/xxx
		v1Group.POST("/add", AddTeacherApi)
		v1Group.POST("/batchdel", BatchDelTeacherApi)
		v1Group.POST("/edit", EditTeacherApi)
		v1Group.GET("/list", ListTeacherApi)

		// localhost:8888/admin/label/add
		v2Group.POST("/add", AddLabelApi)
		v2Group.POST("/batchdel", BatchDelLabelApi)
		v2Group.POST("/edit", EditLabelApi)
		v2Group.GET("/list", ListLabelApi)

		//我的教师
		// localhost:8888/myteacher/addlabels
		v3Group.POST("/addlabels", MyTeacherAddLabelApi)
		v3Group.GET("/list", MyTeacherListApi)

	}
	return r
}
