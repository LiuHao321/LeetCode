package main

import (
	"bytes"
	"fmt"
)

func f(s string) string {
	n := len(s)
	var bt bytes.Buffer
	for i := 0; i < n-1; i++ {
		if s[i:i+2] == "aa" {
			bt.WriteString("abc")
			i++
		} else {
			bt.WriteByte(s[i])
		}
	}
	bt.WriteByte(s[n-1])
	return bt.String()
}

func main() {
	s := "aabcsaaadwdaafeawfa"
	fmt.Println(f(s))
}
