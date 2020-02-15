package initRouter

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
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
