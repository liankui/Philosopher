package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// user相关的转发接口
func UserSave(context *gin.Context) {
	// 通过Param方法获取参数
	username := context.Param("name")
	context.String(http.StatusOK, "用户:"+username+"已经保存")
}

// 通过query方法获取参数
func UserSaveByQuery(ctx *gin.Context) {
	username := ctx.Query("name")
	age := ctx.Query("age")
	ctx.String(http.StatusOK, "用户:"+username+"年龄:"+age+"已经保存")
}