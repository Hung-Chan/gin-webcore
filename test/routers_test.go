package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gin-webcore/routers"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

var token string

var router = routers.InitRouter()

type Response struct {
	Data struct {
		AccessToken string `json:"accessToken"`
		ExpiresIn   int    `json:" expiresIn"`
		TokenType   string `json:"tokenType"`
	} `json:"data"`
}

// TestLoginRouter 登入 .
func TestLoginRouter(t *testing.T) {

	t.Logf("登入api測試")

	// router := routers.InitRouter()

	w := httptest.NewRecorder()

	api := "/auth/login"

	param := make(map[string]interface{})
	param["account"] = "admin"
	param["password"] = "qaz123"

	bytesData, err := json.Marshal(param)
	if err != nil {
		t.Logf(err.Error())
		return
	}

	reader := bytes.NewReader(bytesData)

	request, err := http.NewRequest("POST", api, reader)
	if err != nil {
		t.Logf(err.Error())
		return
	}
	request.Header.Set("Content-Type", "application/json;charset=UTF-8")

	router.ServeHTTP(w, request)

	assert.Equal(t, http.StatusOK, w.Code)
	// fmt.Println(w.Body.String())
	// t.Logf(w.Body.String())

	respBytes, err := ioutil.ReadAll(w.Body)
	if err != nil {
		t.Logf(err.Error())
		return
	}

	data := Response{}
	dataErr := json.Unmarshal(respBytes, &data)
	if dataErr != nil {
		fmt.Println(dataErr)
		return
	}

	token = data.Data.AccessToken
	// fmt.Println(token)
	// t.Logf(string(respBytes))
	request.Header.Set("Authorization", "bearer "+token)

}

// TestInfoRouter 登入者資訊 .
func TestInfoRouter(t *testing.T) {

	t.Logf("info api測試")

	// router := routers.InitRouter()

	w := httptest.NewRecorder()

	request, _ := http.NewRequest("GET", "/auth/info", nil)
	request.Header.Set("Authorization", "bearer "+token)

	router.ServeHTTP(w, request)

	assert.Equal(t, http.StatusOK, w.Code)
	// fmt.Println(w.Body.String())
	// t.Logf(w.Body.String())
}

// TestSidebarMenuRouter 側邊欄資訊 .
func TestSidebarMenuRouter(t *testing.T) {

	t.Logf("SidebarMenu api測試")

	// router := routers.InitRouter()

	w := httptest.NewRecorder()

	request, _ := http.NewRequest("GET", "/auth/sidebarMenu", nil)
	request.Header.Set("Authorization", "bearer "+token)

	router.ServeHTTP(w, request)

	assert.Equal(t, http.StatusOK, w.Code)
	// fmt.Println(w.Body.String())
	// t.Logf(w.Body.String())
}

// TestLogoutRouter 登出 .
// func TestLogoutRouter(t *testing.T) {
// 	router := routers.InitRouter()

// 	w := httptest.NewRecorder()

// 	request, _ := http.NewRequest("POST", "/auth/logout", nil)
// 	request.Header.Set("Authorization", "bearer "+token)

// 	router.ServeHTTP(w, request)

// 	assert.Equal(t, http.StatusOK, w.Code)
// 	// fmt.Println(w.Body.String())
// 	// t.Logf(w.Body.String())
// }
