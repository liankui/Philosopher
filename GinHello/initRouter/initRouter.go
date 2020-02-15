package initRouter

import (
	"github.com/gin-gonic/gin"
	"github.com/yueekee/Philosopher/GinHello/handler"
	"net/http"
	"strings"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	userRouter := router.Group("/user")
	{
		userRouter.GET("/:name", handler.UserSave)
		userRouter.GET("", handler.UserSaveByQuery)
	}

	indexRouter := router.Group("/")
	{
		indexRouter.Any("", retHelloGinAndMethod)
	}

	return router
}

func retHelloGinAndMethod(context *gin.Context) {
	context.String(http.StatusOK, "hello gin " +  strings.ToLower(context.Request.Method) + " method")
}
