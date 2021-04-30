package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	args := os.Args[1:]
	int_args := convert_to_int_array(args)
	fmt.Println(int_args)
}

func convert_to_int_array(args []string) []int {
	result := make([]int, len(args))
	for i := 0; i < len(args); i++ {
		val, err := strconv.Atoi(args[i])
		if err == nil {
			result[i] = val
		}
	}
	return result
}

func array_phonetic(values []int) []string {
	result := make([]string, len(values))

	return result
}
