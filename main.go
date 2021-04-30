package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	args := os.Args[1:]
	int_args := convertToIntegerArray(args)
	result := converToPhonetic(int_args)
	fmt.Println(strings.Join(result, ","))
}

func convertToIntegerArray(args []string) []int {
	result := make([]int, len(args))
	for i := 0; i < len(args); i++ {
		val, err := strconv.Atoi(args[i])
		if err == nil {
			result[i] = val
		}
	}
	return result
}

func converToPhonetic(values []int) []string {
	result := make([]string, len(values))
	conversionFactor := []string{"Zero", "One", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine"}
	for i := 0; i < len(values); i++ {
		currentVal := values[i]
		if currentVal != 0 {
			current := ""
			for currentVal > 0 {
				index := currentVal % 10
				current = fmt.Sprintf("%s%s", conversionFactor[index], current)
				currentVal /= 10
			}
			result[i] = current
		} else {
			result[i] = conversionFactor[currentVal]
		}
	}
	return result
}
