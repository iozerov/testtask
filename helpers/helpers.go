package helpers

import (
	"io"
	"log"
	"net/http"
	"strings"
)

func ParseData(link string) []string {
	resp, err := http.Get(link)
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

		ipAddresses := strings.Split(bodyString, "\n")
		return ipAddresses
	}
	return nil
}

func FindDifferences(oldSlice []string, newSlice []string) map[string][]string {
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
