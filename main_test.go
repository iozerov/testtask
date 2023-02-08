package main

import (
	"bytes"
	"encoding/json"
	"example/web-service-gin/addresses"
	"example/web-service-gin/helpers"
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
	//pass
	body := Body{
		IpAddress: "111.222.333.444",
	}
	trueVal, _ := json.Marshal(body)
	trueReq, _ := http.NewRequest("POST", "/contains", bytes.NewBuffer(trueVal))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, trueReq)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "false", w.Body.String())

	//validation error
	body.IpAddress = "1.2.3"
	wrongVal, _ := json.Marshal(body)
	wrongReq, _ := http.NewRequest("POST", "/contains", bytes.NewBuffer(wrongVal))
	wa := httptest.NewRecorder()
	r.ServeHTTP(wa, wrongReq)
	assert.Equal(t, http.StatusBadRequest, wa.Code)

	parsedData := []string{"127.0.0.1", "178.154.220.80", "43.135.157.133", "189.44.9.51", "102.129.37.140", "154.222.226.110"}
	ipAddresses := addresses.New(parsedData)
	contains := ipAddresses.Contains("43.135.157.133")
	assert.Equal(t, contains, true)
	contains = ipAddresses.Contains("88.55.157.111")
	assert.Equal(t, contains, false)
}

func TestDifference(t *testing.T) {
	//full slices
	old := []string{"124.128.223.82", "178.154.220.80", "77.68.26.238", "189.44.9.51", "102.129.37.140"}
	new := []string{"127.0.0.1", "178.154.220.80", "43.135.157.133", "189.44.9.51", "102.129.37.140", "154.222.226.110"}
	expected := map[string][]string{
		"added":   {"127.0.0.1", "43.135.157.133", "154.222.226.110"},
		"removed": {"124.128.223.82", "77.68.26.238"},
	}
	got := helpers.FindDifferences(old, new)
	assert.Equal(t, expected, got)

	//first empty
	old = []string{}
	new = []string{"127.0.0.1", "43.135.157.133", "154.222.226.110"}
	expected = map[string][]string{
		"added":   {"127.0.0.1", "43.135.157.133", "154.222.226.110"},
		"removed": {},
	}
	got = helpers.FindDifferences(old, new)
	assert.Equal(t, expected, got)

	//second empty
	old = []string{"127.0.0.1", "43.135.157.133", "154.222.226.110"}
	new = []string{}
	expected = map[string][]string{
		"added":   {},
		"removed": {"127.0.0.1", "43.135.157.133", "154.222.226.110"},
	}
	got = helpers.FindDifferences(old, new)
	assert.Equal(t, expected, got)
}

func TestFilter(t *testing.T) {
	parsedData := []string{"127.0.0.1", "178.154.220.80", "43.135.157.133", "189.44.9.51", "102.129.37.140", "154.222.226.110"}
	ipAddresses := addresses.New(parsedData)
	matches := ipAddresses.Filter("15")
	assert.Equal(t, len(matches), 3)
	matches = ipAddresses.Filter("754")
	assert.Equal(t, len(matches), 0)
}

func TestDelete(t *testing.T) {
	parsedData := []string{"127.0.0.1", "178.154.220.80", "43.135.157.133", "189.44.9.51", "102.129.37.140", "154.222.226.110"}
	ipAddresses := addresses.New(parsedData)
	ipAddresses = ipAddresses.Delete()
	assert.Equal(t, len(ipAddresses.List), 0)
}
