package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/yueekee/Philosopher/GinHello/model"
	"github.com/yueekee/Philosopher/GinHello/utils"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"
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

// 点击右上角email，进入用户的详情页
func UserProfile(ctx *gin.Context) {
	id := ctx.Query("id")
	var user model.UserModel
	i, err := strconv.Atoi(id)
	u, e := user.QueryById(i)
	if err != nil || e != nil {
		ctx.HTML(http.StatusOK, "error.tmpl", gin.H{
			"error": e,
		})
	}
	ctx.HTML(http.StatusOK, "user_profile.tmpl", gin.H{
		"user": u,
	})
}

func UpdateUserProfile(ctx *gin.Context) {
	var user model.UserModel
	if err := ctx.ShouldBind(&user); err != nil {
		ctx.HTML(http.StatusOK, "error.tmpl", gin.H{
			"error": err,
		})
		log.Panicln("绑定发生错误", err.Error())
	}
	file, e := ctx.FormFile("avatar-file")
	if e != nil {
		ctx.HTML(http.StatusOK, "error.tmpl", gin.H{
			"error": e,
		})
		log.Panicln("绑定发生错误", e.Error())
	}
	path := utils.RootPath()
	path = filepath.Join(path, "avatar") // 生成[path]/avatar
	if err := os.MkdirAll(path, os.ModePerm); err != nil {
		ctx.HTML(http.StatusOK, "error.tmpl", gin.H{
			"error": e,
		})
		log.Panicln("无法创建文件夹", e.Error())
	}
	fileName := strconv.FormatInt(time.Now().Unix(), 10) + file.Filename
	if err := ctx.SaveUploadedFile(file, path+fileName); err != nil {
		ctx.HTML(http.StatusOK, "error.tmpl", gin.H{
			"error": e,
		})
		log.Panicln("无法保存文件", e.Error())
	}
}