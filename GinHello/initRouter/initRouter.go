package initRouter

import (
	"github.com/gin-gonic/gin"
	"github.com/yueekee/Philosopher/GinHello/handler"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.LoadHTMLGlob("templates/*")  // 注意这里templates前面没有/

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
