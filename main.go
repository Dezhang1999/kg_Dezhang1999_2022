package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	args := os.Args[1:]
	int_args := convertToIntegerArray(args)
	fmt.Println(converToPhonetic(int_args))
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
		current := ""
		if currentVal != 0 {
			for currentVal > 0 {
				index := currentVal % 10
				currentCoversion := conversionFactor[index]
				current += fmt.Sprintf("%s%s", currentCoversion, current)
				currentVal /= 10
			}
			result[i] = current
		} else {
			result[i] = conversionFactor[currentVal]
		}
	}
	return result
}
