package main

import (
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type Body struct {
	IpAddress string `json:"ip_address" binding:"required,min=7,max=15"`
}

var ipAddresses []string
var differences map[string][]string

func main() {
	router := gin.Default()
	ipAddresses = getIpAddresses()
	router.GET("/refresh", refresh)
	router.GET("/last-changes", lastChanges)
	router.GET("/filter", filter)
	router.GET("/count", count)
	router.POST("/contains", contains)
	router.DELETE("/delete", delete)
	router.Run("localhost:8080")
}

func getIpAddresses() []string {
	resp, err := http.Get("https://raw.githubusercontent.com/stamparm/ipsum/master/levels/5.txt")

	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Fatalln(err)
		}
		bodyString := string(bodyBytes)

		// ipAddresses := delete_empty(strings.Split(bodyString, "\n"))
		ipAddresses := strings.Split(bodyString, "\n")
		return ipAddresses
	}
	return nil
}

func refresh(c *gin.Context) {
	newIpAddresses := getIpAddresses()
	c.IndentedJSON(http.StatusOK, "Refreshing data was started.")
	differences = findDifferences(ipAddresses, newIpAddresses)
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
	matches := []string{}
	param := c.Query("query")

	if param != "" {
		for _, ip := range ipAddresses {
			if strings.Contains(ip, param) {
				matches = append(matches, ip)
			}
		}
	}
	c.IndentedJSON(http.StatusOK, matches)
}

func count(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, len(ipAddresses))
}

func contains(c *gin.Context) {
	contains := false
	body := Body{}
	if err := c.ShouldBindJSON(&body); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{"message": "Invalid inputs. Please check your inputs"})
		return
	}

	for _, ip_address := range ipAddresses {
		if ip_address == body.IpAddress {
			contains = true
		}
	}

	c.JSON(http.StatusAccepted, contains)
}

func delete(c *gin.Context) {
	ipAddresses = ipAddresses[:0]
	c.JSON(http.StatusAccepted, ipAddresses)
}

func findDifferences(oldSlice []string, newSlice []string) map[string][]string {
	diffData := make(map[string][]string)
	// Loop two times, first to find oldSlice strings not in newSlice,
	// second loop to find newSlice strings not in oldSlice
	for i := 0; i < 2; i++ {
		for _, s1 := range oldSlice {
			found := false
			for _, s2 := range newSlice {
				if s1 == s2 {
					found = true
					break
				}
			}
			// String not found.
			if !found {
				if i == 0 {
					diffData["removed"] = append(diffData["removed"], s1)
				} else {
					diffData["added"] = append(diffData["added"], s1)
				}
			}
		}
		// Swap the slices, only if it was the first loop
		if i == 0 {
			oldSlice, newSlice = newSlice, oldSlice
		}
	}
	return diffData
}
