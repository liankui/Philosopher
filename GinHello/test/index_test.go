package test

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/yueekee/Philosopher/GinHello/initRouter"
	"net/http"
	"net/http/httptest"
	"testing"
)

var router *gin.Engine

func init() {
	gin.SetMode(gin.TestMode)
	router = initRouter.SetupRouter()
}

func TestIndexSetupRouter(t *testing.T) {
	w := httptest.NewRecorder()
	request, _ := http.NewRequest(http.MethodGet, "/", nil)
	router.ServeHTTP(w, request)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "hello gin get method", w.Body.String())
}

func TestIndexHtml(t *testing.T) {
	w := httptest.NewRecorder()
	request, _ := http.NewRequest(http.MethodGet, "/", nil)
	router.ServeHTTP(w, request)
	assert.Equal(t, http.StatusOK, w.Code)
	//assert.Contains(t,w.Body.String(),"hello gin get method","返回的HTML页面中应该包含 hello gin get method")
}