package main

import "fmt"
import "regexp"
import "strings"

func Comtain(str string) {
	if strings.Contains(str, "iPhone") {
		fmt.Print("heheeei")
	}
}

var r *regexp.Regexp = regexp.MustCompile(`iPhone`)

func Rege(str string) {
	if r.MatchString(str) {
		fmt.Print("geggggi")
	}
}
