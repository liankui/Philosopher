package test

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/yueekee/Philosopher/GinHello/initRouter"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strconv"
	"testing"
)

var router *gin.Engine

func init() {
	gin.SetMode(gin.TestMode)
	router = initRouter.SetupRouter()
}

func TestUserSave(t *testing.T) {
	username := "eric"
	router := initRouter.SetupRouter()
	w := httptest.NewRecorder()
	request, _ := http.NewRequest(http.MethodGet, "/user/"+username, nil)
	router.ServeHTTP(w, request)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "用户:"+username+"已经保存", w.Body.String())
}

func TestUserSaveByQuery(t *testing.T) {
	username := "eric"
	age := 27
	router := initRouter.SetupRouter()
	w := httptest.NewRecorder()
	request, _ := http.NewRequest(http.MethodGet, "/user?name="+username+"&age="+strconv.Itoa(age), nil)
	router.ServeHTTP(w, request)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "用户:"+username+"年龄:"+strconv.Itoa(age)+"已经保存", w.Body.String())
}

func TestUserSaveWithoutAge(t *testing.T) {
	username := "eric"
	router := initRouter.SetupRouter()
	w := httptest.NewRecorder()
	request, _ := http.NewRequest(http.MethodGet, "/user?name="+username, nil)
	router.ServeHTTP(w, request)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "用户:"+username+"年龄:27"+"已经保存", w.Body.String())
}

func TestUserRegister(t *testing.T) {
	value := url.Values{}
	value.Add("email", "yuekewin@gmail.com")
	value.Add("password", "123456")
	value.Add("password-again", "123456")
	w := httptest.NewRecorder()
	request, _ := http.NewRequest(http.MethodPost, "/user/register", nil)
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded; param=value")
	router.ServeHTTP(w, request)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestUserRegisterError(t *testing.T) {
	value := url.Values{}
	value.Add("email", "yuekn@gmail.com")
	value.Add("password", "123456")
	value.Add("password-again", "asda")
	w := httptest.NewRecorder()
	request, _ := http.NewRequest(http.MethodPost, "/user/register", bytes.NewBufferString(value.Encode()))
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded; param=value")
	router.ServeHTTP(w, request)
	assert.Equal(t, http.StatusOK, w.Code)
}
