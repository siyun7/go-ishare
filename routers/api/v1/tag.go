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
	name := content.PostForm("name")
	state := com.StrTo(content.DefaultPostForm("state", "0")).MustInt()
	createdBy := content.PostForm("created_by")

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
	id := com.StrTo(content.Param("id")).MustInt()
	modifiedBy := content.PostForm("modified_by")
	name  := content.PostForm("name ")

	valid := validation.Validation{}

	var state int = 1
	if arg := content.PostForm("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
		valid.Range(state, 0, 1, "state")
	}

	valid.Required(id, "id").Message("ID不能为空")
	valid.Required(modifiedBy, "modified_by").Message("修改人不能为空")
	valid.MaxSize(modifiedBy, 100, "modified_by").Message("修改人最长为100字符")
	valid.MaxSize(name, 100, "modified_by").Message("名称最长为100字符")

	code := enum.INVALID_PARAMS
	msg := ""

	if !valid.HasErrors() {
		code = enum.SUCCESS

		if models.ExistTagByID(id) {
			data := make(map[string]interface{})

			if name != "" {
				data["name"] = name
			}

			data["state"] = state
			data["modified_by"] = modifiedBy

			models.EditTag(id, data)
		} else {
			code = enum.ERROR_NOT_EXIST_TAG
		}

		msg = enum.GetMsg(code)
	} else {
		msg = valid.Errors[0].Message
	}

	content.JSON(http.StatusOK, gin.H{
		"code" : code,
		"msg" : msg,
		"data": make(map[string]interface{}),
	})


}

func DeleteTag(content *gin.Context) {
	id := com.StrTo(content.PostForm("id")).MustInt()

	valid := validation.Validation{}

	valid.Min(id, 1, "id").Message("ID必须大于0")

	code := enum.INVALID_PARAMS

	if !valid.HasErrors() {
		code = enum.SUCCESS

		if models.ExistTagByID(id) {
			models.Deletetag(id)
		} else {
			code = enum.ERROR_NOT_EXIST_TAG
		}
	}

	content.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg": enum.GetMsg(code),
		"data": make(map[string]interface{}),
	})
}
