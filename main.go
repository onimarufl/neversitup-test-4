package main

import (
	"fmt"
	"strings"
)

func main() {
	slice_of_slices := make([][]string, 3)
	slice_of_slices[0] = []string{":)", ";(", ";}", ":-D"}
	slice_of_slices[1] = []string{";D", ":-(", ":-)", ";~)"}
	slice_of_slices[2] = []string{";]", ":[", ";*", ":$", ";-D"}

	for _, slice := range slice_of_slices {
		fmt.Print(slice)
		count := 0
		for _, str := range slice {

			if strings.Contains(str, ")") || strings.Contains(str, "D") {
				count += 1
			}
		}

		fmt.Println(" / count : ", count)
	}
}
