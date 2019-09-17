package main

import (
	"fmt"
	"strings"
)

func main() {
	ass := "Apple Napkin"

	bass := strings.Split(ass, " ")

	for idx, elem := range bass {
		fmt.Println(idx, elem)
	}
}