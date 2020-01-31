package util

import (
	"github.com/gin-gonic/gin"
	"github.com/siyun7/go-ishare/pkg/enum"
	"net/http"
)

type Response struct {
	Code int
	Msg  string
	Data interface{}
}

func ResponseWithJson(content *gin.Context, r Response) {
	content.JSON(http.StatusOK, gin.H{
		"code": r.Code,
		"msg":  enum.GetMsg(r.Code),
		"data": r.Data,
	})
}
