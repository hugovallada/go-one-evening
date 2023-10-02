package main

import (
	"errors"
	"fmt"
)

func main() {
	result, err := Divide(100, 50)
	fmt.Println("Result:", result, "Error:", err)
}

func Divide(y, x float64) (float64, error) {
	if x == 0 {
		return 0, errors.New("cant divide by 0")
	}
	return y/x, nil
}
