package main

import (
	"example/web-service-gin/addresses"
	"example/web-service-gin/helpers"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Body struct {
	IpAddress string `json:"ip_address" binding:"required,min=7,max=15"`
}

var ipAddresses addresses.Addresses
var differences map[string][]string

func main() {
	router := gin.Default()
	parsedData := helpers.ParseData("https://raw.githubusercontent.com/stamparm/ipsum/master/levels/5.txt")
	ipAddresses = addresses.New(parsedData)
	router.GET("/refresh", refresh)
	router.GET("/last-changes", lastChanges)
	router.GET("/filter", filter)
	router.GET("/count", count)
	router.POST("/contains", contains)
	router.DELETE("/delete", delete)
	router.Run("localhost:8080")
}

func refresh(c *gin.Context) {
	go refreshData()
	c.IndentedJSON(http.StatusOK, "Refreshing data was started.")
}

func refreshData() {
	parsedData := helpers.ParseData("https://raw.githubusercontent.com/stamparm/ipsum/master/levels/5.txt")
	newIpAddresses := addresses.New(parsedData)
	differences = helpers.FindDifferences(ipAddresses.List, newIpAddresses.List)
	ipAddresses = newIpAddresses
}

func lastChanges(c *gin.Context) {
	if len(differences) == 0 {
		differences = map[string][]string{
			"added":   {},
			"removed": {},
		}
	}
	c.IndentedJSON(http.StatusOK, differences)
}

func filter(c *gin.Context) {
	param := c.Query("query")
	matches := ipAddresses.Filter(param)
	c.IndentedJSON(http.StatusOK, matches)
}

func count(c *gin.Context) {
	fmt.Println(ipAddresses.List)
	c.IndentedJSON(http.StatusOK, len(ipAddresses.List))
}

func contains(c *gin.Context) {
	body := Body{}
	if err := c.ShouldBindJSON(&body); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{"message": "Invalid inputs. Please check your inputs"})
		return
	}
	contains := ipAddresses.Contains(body.IpAddress)
	c.JSON(http.StatusOK, contains)
}

func delete(c *gin.Context) {
	ipAddresses = ipAddresses.Delete()
	c.JSON(http.StatusOK, ipAddresses.List)
}
