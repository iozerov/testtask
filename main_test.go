package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func SetUpRouter() *gin.Engine {
	router := gin.Default()
	return router
}

func TestContains(t *testing.T) {
	r := SetUpRouter()
	r.POST("/contains", contains)
	//success
	body := Body{
		IpAddress: "111.222.333.444",
	}
	trueVal, _ := json.Marshal(body)
	trueReq, _ := http.NewRequest("POST", "/contains", bytes.NewBuffer(trueVal))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, trueReq)
	assert.Equal(t, http.StatusAccepted, w.Code)

	//validation error
	body.IpAddress = "1.2.3"
	wrongVal, _ := json.Marshal(body)
	wrongReq, _ := http.NewRequest("POST", "/contains", bytes.NewBuffer(wrongVal))
	wa := httptest.NewRecorder()
	r.ServeHTTP(wa, wrongReq)
	assert.Equal(t, http.StatusBadRequest, wa.Code)

}

func TestDifference(t *testing.T) {
	old := []string{"124.128.223.82", "178.154.220.80", "77.68.26.238", "189.44.9.51", "102.129.37.140"}
	new := []string{"127.0.0.1", "178.154.220.80", "43.135.157.133", "189.44.9.51", "102.129.37.140", "154.222.226.110"}
	expected := map[string][]string{
		"added":   {"127.0.0.1", "43.135.157.133", "154.222.226.110"},
		"removed": {"124.128.223.82", "77.68.26.238"},
	}
	got := findDifferences(old, new)
	assert.Equal(t, expected, got)
}
