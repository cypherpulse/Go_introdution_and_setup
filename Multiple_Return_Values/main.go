package main

import (
	"fmt"
	"strings"
)

func getInitials(n string) (string, string) {
	s := strings.ToUpper(n)
	names := strings.Split(s, " ")

	var initials []string
	for _, v := range names {
		initials = append(initials, string(v[:1]))
	}

	if len(initials) > 1 {
		return initials[0], initials[1]
	}

	return initials[0], "_"
}

func main() {
	fn1, sn1 := getInitials("tifa lockhart")
	fmt.Printf("First Name Initial: %s, Second Name Initial: %s\n", fn1, sn1)

	fn2, sn2 := getInitials("cloud strife")
	fmt.Printf("First Name Initial: %s, Second Name Initial: %s\n", fn2, sn2)

	fn3, sn3 := getInitials("Aerith")
	fmt.Printf("First Name Initial: %s, Second Name Initial: %s\n", fn3, sn3)
}
