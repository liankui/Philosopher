package handler

import (
	"database/sql"
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

func UserRegister(context *gin.Context) {
	var user model.UserModel
	if err := context.ShouldBind(&user); err != nil {
		context.String(http.StatusBadRequest, "输入的数据不合法")
		log.Panicln("err ->", err.Error())
	}
	passwordAgain := context.PostForm("password-again")
	if passwordAgain != user.Password {
		context.String(http.StatusBadRequest, "密码校验无效，两次密码不一致")
		log.Panicln("密码校验无效，两次密码不一致")
	}
	user.Save()
	context.Redirect(http.StatusMovedPermanently, "/")
}

func UserLogin(context *gin.Context) {
	var user model.UserModel
	if e := context.Bind(&user); e != nil {
		log.Panicln("login 绑定错误", e.Error())
	}

	u := user.QueryByEmail()
	if u.Password == user.Password {
		context.SetCookie("user_cookie", string(u.Id), 1000,
			"/", "localhost", false, true)
		log.Println("登录成功", u.Email)
		context.HTML(http.StatusOK, "index.tmpl", gin.H{
			"email": u.Email,
			"id": u.Id,
		})
	}
}

// 点击右上角email，进入用户的详情页
func UserProfile(context *gin.Context) {
	id := context.Query("id")
	var user model.UserModel
	i, err := strconv.Atoi(id)
	u, e := user.QueryById(i)
	if err != nil || e != nil {
		context.HTML(http.StatusOK, "error.tmpl", gin.H{
			"error": e,
		})
	}
	context.HTML(http.StatusOK, "user_profile.tmpl", gin.H{
		"user": u,
	})
}

func UpdateUserProfile(context *gin.Context) {
	var user model.UserModel
	if err := context.ShouldBind(&user); err != nil {
		context.HTML(http.StatusOK, "error.tmpl", gin.H{
			"error": err,
		})
		log.Panicln("绑定user发生错误", err.Error())
	}
	file, e := context.FormFile("avatar-file")
	if e != nil {
		context.HTML(http.StatusOK, "error.tmpl", gin.H{
			"error": e,
		})
		log.Panicln("绑定avatar-file发生错误", e.Error())
	}
	path := utils.RootPath()
	path = filepath.Join(path, "avatar") // 生成[path]/avatar
	if err := os.MkdirAll(path, os.ModePerm); err != nil {
		context.HTML(http.StatusOK, "error.tmpl", gin.H{
			"error": e,
		})
		log.Panicln("无法创建文件夹", e.Error())
	}
	fileName := strconv.FormatInt(time.Now().Unix(), 10) + file.Filename
	path = filepath.Join(path, fileName)
	if err := context.SaveUploadedFile(file, path); err != nil {
		context.HTML(http.StatusOK, "error.tmpl", gin.H{
			"error": e,
		})
		log.Panicln("无法保存文件", e.Error())
	}

	avatarUrl := "http://localhost:8080/avatar/" + fileName
	user.Avatar = sql.NullString{String:avatarUrl}
	if err := user.Update(user.Id); err != nil {
		context.HTML(http.StatusOK, "error.tmpl", gin.H{
			"error": e,
		})
		log.Panicln("数据无法更新", e.Error())
	}
	context.Redirect(http.StatusMovedPermanently, "/user/profile?id="+strconv.Itoa(user.Id))
}