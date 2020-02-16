package initRouter

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/yueekee/Philosopher/GinHello/handler"
	"os"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

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
		userRouter.GET("/:name", handler.UserSave)
		userRouter.GET("", handler.UserSaveByQuery)
		userRouter.POST("/register", handler.UserRegister)
	}

	indexRouter := router.Group("/")
	{
		indexRouter.Any("", handler.Index)
	}

	return router
}
