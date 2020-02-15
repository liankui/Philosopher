package test

import (
	"github.com/stretchr/testify/assert"
	"github.com/yueekee/Philosopher/GinHello/initRouter"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestIndexSetupRouter(t *testing.T) {
	router := initRouter.SetupRouter()
	w := httptest.NewRecorder()
	request, _ := http.NewRequest(http.MethodGet, "/", nil)
	router.ServeHTTP(w, request)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "hello gin get method", w.Body.String())
}
