package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Auth() gin.HandlerFunc {
	return func(context *gin.Context) {
		cookie, e := context.Request.Cookie("user_cookie")
		if e == nil {
			context.SetCookie("user_cookie", cookie.Value, cookie.MaxAge,
				cookie.Path, cookie.Domain, cookie.Secure, cookie.HttpOnly)
			context.Next()
		} else {
			context.Abort()
			context.HTML(http.StatusUnauthorized, "401.tmpl", nil)
		}
	}
}