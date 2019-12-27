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

func GetTags(content *gin.Context) {
	name := content.Query("name")

	maps := make(map[string]interface{})
	data := make(map[string]interface{})

	if name != "" {
		maps["name"] = name
	}

	var state int = 1
	if arg := content.Query("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
		maps["state"] = state
	}

	code := enum.SUCCESS

	data["lists"] = models.GetTags(util.GetPage(content), setting.PageSize, maps)
	data["total"] = models.GetTagTotal(maps)

	content.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  enum.GetMsg(code),
		"data": data,
	})

}

func AddTag(content *gin.Context) {
	name := content.Query("name")
	state := com.StrTo(content.DefaultQuery("state", "0")).MustInt()
	createdBy := content.Query("created_by")

	valid := validation.Validation{}
	valid.Required(name, "name").Message("名称不能为空")
	valid.MaxSize(name, 100, "name").Message("名称最长为100字符")

	valid.Required(createdBy, "created_by").Message("创建人不能为空")
	valid.MaxSize(createdBy, 100, "created_by").Message("创建人最长为100字符")
	valid.Range(state, 0, 1, "state").Message("状态只允许0或1")

	code := enum.INVALID_PARAMS

	if ! valid.HasErrors() {
		if !models.ExistTagByName(name) {
			code = enum.SUCCESS
			models.AddTag(name, state, createdBy)
		} else {
			code = enum.ERROR_EXIST_TAG
		}
	}

	content.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg" : enum.GetMsg(code),
		"data" : make(map[string]string),
	})

}

func EditTag(content *gin.Context) {

}

func DeleteTag(content *gin.Context) {

}
