package initRouter

import (
	"github.com/gin-gonic/gin"
	"github.com/yueekee/Philosopher/GinHello/handler"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.Static("/statics", "./statics")
	router.StaticFile("/favicon.ico", "./favicon.ico")

	if mode := gin.Mode(); mode == gin.TestMode {
		router.LoadHTMLGlob("./../templates/*")
	} else {
		router.LoadHTMLGlob("templates/*")
	}

	userRouter := router.Group("/user")
	{
		userRouter.GET("/:name", handler.UserSave)
		userRouter.GET("", handler.UserSaveByQuery)
	}

	indexRouter := router.Group("/")
	{
		indexRouter.Any("", handler.Index)
	}

	return router
}
