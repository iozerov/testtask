package addresses

import (
	"strings"
)

type Addresses struct {
	List []string
}

func New(dataList []string) Addresses {
	addresses := Addresses{dataList}
	return addresses
}

func (a Addresses) Contains(ip_address string) bool {
	contains := false
	for _, el := range a.List {
		if el == ip_address {
			contains = true
		}
	}
	return contains
}

func (a Addresses) Filter(str string) []string {
	matches := []string{}
	if str != "" {
		for _, ip := range a.List {
			if strings.Contains(ip, str) {
				matches = append(matches, ip)
			}
		}
	}
	return matches
}

func (a Addresses) Delete() Addresses {
	a.List = a.List[:0]
	return a
}
