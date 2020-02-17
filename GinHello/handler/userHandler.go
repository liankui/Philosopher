package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/yueekee/Philosopher/GinHello/model"
	"log"
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
	age := ctx.DefaultQuery("age", "27")
	ctx.String(http.StatusOK, "用户:"+username+"年龄:"+age+"已经保存")
}

func UserRegister(ctx *gin.Context) {
	var user model.UserModel
	if err := ctx.ShouldBind(&user); err != nil {
		ctx.String(http.StatusBadRequest, "输入的数据不合法")
		log.Panicln("err ->", err.Error())
	}
	id := user.Save()
	log.Println("id is", id)
	ctx.Redirect(http.StatusMovedPermanently, "/")
}

func UserLogin(context *gin.Context) {
	var user model.UserModel
	if e := context.Bind(&user); e != nil {
		log.Panicln("login 绑定错误", e.Error())
	}

	u := user.QueryByEmail()
	if u.Password == user.Password {
		log.Println("登录成功", u.Email)
		context.HTML(http.StatusOK, "index.tmpl", gin.H{
			"email": u.Email,
		})
	}
}
