package v1

import (
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/siyun7/go-ishare/models"
	"github.com/siyun7/go-ishare/pkg/enum"
	"github.com/siyun7/go-ishare/pkg/setting"
	"github.com/siyun7/go-ishare/pkg/util"
	"github.com/unknwon/com"
	"net/http"
)

func AddUser(ctx *gin.Context) {
	name := ctx.PostForm("name")
	password := ctx.PostForm("password")

	valid := validation.Validation{}

	valid.Required(name, "name").Message("请输入用户名")
	valid.AlphaNumeric(name, "name").Message("用户名限定为0-9a-zA-Z")
	valid.MinSize(name, 3, "name").Message("用户名最小长度为3")
	valid.MaxSize(name, 50, "name").Message("用户名最大长度为50")

	valid.Required(password, "password").Message("请输入密码")
	valid.MinSize(password, 3, "password").Message("密码最小长度为3")
	valid.MaxSize(password, 50, "password").Message("密码最大长度为50")

	code := enum.INVALID_PARAMS
	msg := enum.GetMsg(code)

	if !valid.HasErrors() {

		if !models.ExistUserByName(name) {
			password = util.StrMd5(password)
			models.AddUser(name, password)
		} else {
			code = enum.ERROR_EXIST_USER
		}

		msg = enum.GetMsg(code)
	} else {
		msg = valid.Errors[0].Message
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  msg,
		"data": make(map[string]string),
	})
}

func GetUsers(ctx *gin.Context) {
	id := ctx.Query("id")
	status := ctx.Query("status")

	maps := make(map[string]interface{})
	data := make(map[string]interface{})

	if id != "" {
		maps["id"] = com.StrTo(id).MustInt()
	}

	if status != "" {
		maps["status"] = com.StrTo(status).MustInt()
	}

	code := enum.SUCCESS

	data["lists"] = models.GetUsers(util.GetPage(ctx), setting.PageSize, maps)
	data["total"] = 12
	println(code)



}

func GetUser(ctx *gin.Context) {

}

func EditUser(ctx *gin.Context) {

}

func DeleteUser(ctx *gin.Context) {

}
