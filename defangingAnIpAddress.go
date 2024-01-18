package main

import (
	"fmt"
	"strings"
)

func defangIPaddr(address string) string {
	return strings.ReplaceAll(address, ".","[.]")
}

func main() {
	fmt.Println(defangIPaddr("1.1.1.1"))
}

//Given a valid (IPv4) IP address, return a defanged version of that IP address.
//A defanged IP address replaces every period "." with "[.]".