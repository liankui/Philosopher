package initRouter

import (
	"github.com/gin-gonic/gin"
	"github.com/yueekee/Philosopher/GinHello/handler"
	"net/http"
	"strings"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/user/:name", handler.UserSave)
	router.GET("/user", handler.UserSaveByQuery)

	router.GET("/", retHelloGinAndMethod)
	router.POST("/", retHelloGinAndMethod)
	router.PUT("/", retHelloGinAndMethod)
	router.DELETE("/", retHelloGinAndMethod)
	router.PATCH("/", retHelloGinAndMethod)
	router.HEAD("/", retHelloGinAndMethod)
	router.OPTIONS("/", retHelloGinAndMethod)

	return router
}

func retHelloGinAndMethod(context *gin.Context) {
	context.String(http.StatusOK, "hello gin " +  strings.ToLower(context.Request.Method) + " method")
}
