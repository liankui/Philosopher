package initRouter

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/yueekee/Philosopher/GinHello/handler"
	"github.com/yueekee/Philosopher/GinHello/utils"
	"net/http"
	"os"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

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

	userRouter := router.Group("/user")
	{
		userRouter.POST("/register", handler.UserRegister)
		userRouter.POST("/login", handler.UserLogin)
		userRouter.GET("/profile/", handler.UserProfile)
		userRouter.POST("/update", handler.UpdateUserProfile)
	}

	indexRouter := router.Group("/")
	{
		indexRouter.Any("", handler.Index)
	}

	return router
}
