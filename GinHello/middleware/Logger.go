package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

func Logger() gin.HandlerFunc {
	return func(context *gin.Context) {
		host := context.Request.Host
		url := context.Request.URL
		method := context.Request.Method
		fmt.Printf("%s::%s \t %s \t %s ", time.Now().Format(
			"2006-01-02 15:04:05"), host, url, method)
		context.Next()		// next是在请求前执行，next函数后是在请求后执行
		fmt.Println(context.Writer.Status())
	}
}
