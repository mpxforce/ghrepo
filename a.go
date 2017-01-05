package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	i := f()
	for { 
		if i >= len(os.Args) {
			break
		}
		fmt.Printf(" %s\n", os.Args[i])
		i++
	}
	for _, arg := range "abcdefg" {
		fmt.Printf("value %c\n", arg)
	}
	sa := [] string { "jopa", "lepo", "haha"}
	fmt.Println(strings.Join(sa, ":"))
	fmt.Println(sa)
}
func f()(i int) {
	return 0
}

