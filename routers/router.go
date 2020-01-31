package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/siyun7/go-ishare/pkg/setting"
	v1 "github.com/siyun7/go-ishare/routers/api/v1"
)

func InitRouter () *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())

	r.Use(gin.Recovery())

	gin.SetMode(setting.RunMode)

	apiv1 := r.Group("/api/v1")

	{
		//获取标签列表
		apiv1.GET("/tags", v1.GetTags)
		//新建标签
		apiv1.POST("/tags", v1.AddTag)
		//更新指定标签
		apiv1.PUT("/tags/:id", v1.EditTag)
		//删除指定标签
		apiv1.DELETE("/tags/:id", v1.DeleteTag)

		apiv1.GET("/test/interface/1", v1.RunDemo1)

		apiv1.GET("/test/interface/2", v1.Demo2)


		//获取用户列表
		apiv1.GET("/users", v1.GetUsers)
		//新建用户
		apiv1.POST("/users", v1.AddUser)
		//更新指定用户
		apiv1.PUT("/users/:id", v1.EditUser)
		//删除指定用户
		apiv1.DELETE("/users/:id", v1.DeleteUser)

	}

	return r
}