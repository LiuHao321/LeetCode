package main

import (
	"fmt"
	"strconv"
	"strings"
)

func check(s string) bool {
	if len(s) < 7 || len(s) > 15 {
		return false
	}
	ip := strings.Split(s, ".")
	if len(ip) != 4 {
		return false
	}
	for i, v := range ip {
		ip_int, err := strconv.Atoi(v)
		if err != nil {
			return false
		}
		if len(v) != 1 && v[0] == '0' {
			return false
		}
		if i == 0 {
			if ip_int < 1 || ip_int > 255 {
				return false
			}
		} else {
			if ip_int < 0 || ip_int > 255 {
				return false
			}
		}
	}
	return true
}
func main() {
	s := "172.16.254.1"
	fmt.Println(check(s))
}
