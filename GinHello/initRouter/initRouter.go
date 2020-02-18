package initRouter

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/yueekee/Philosopher/GinHello/handler"
	"github.com/yueekee/Philosopher/GinHello/handler/article"
	"github.com/yueekee/Philosopher/GinHello/middleware"
	"github.com/yueekee/Philosopher/GinHello/utils"
	"net/http"
	"os"
)

func SetupRouter() *gin.Engine {
	router := gin.New()
	router.Use(middleware.Logger(), gin.Recovery())
	//router := gin.Default()

	router.StaticFS("/avatar", http.Dir(utils.RootPath()+"avatar/"))
	router.Static("/statics", "./statics")
	router.StaticFile("/favicon.ico", "./favicon.ico")

	if mode := gin.Mode(); mode == gin.TestMode {
		router.LoadHTMLGlob("./../templates/*")
		//router.LoadHTMLGlob(filepath.Join(os.Getenv("GOPATH"), `src/github.com/yueekee/Philosopher/GinHello/templates/*`))
	} else {
		router.LoadHTMLGlob("templates/*")
	}
	fmt.Println("os.Args:", os.Args)

	articleRouter := router.Group("")
	{
		// 添加一篇文章
		articleRouter.POST("/article", article.Insert)
	}

	userRouter := router.Group("/user")
	{
		userRouter.POST("/register", handler.UserRegister)
		userRouter.POST("/login", handler.UserLogin)
		userRouter.GET("/profile/", middleware.Auth(), handler.UserProfile)
		userRouter.POST("/update", middleware.Auth(), handler.UpdateUserProfile)
	}

	indexRouter := router.Group("/")
	{
		indexRouter.Any("", handler.Index)
	}

	return router
}
