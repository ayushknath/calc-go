package compute

import (
	"fmt"
	"math"
	"os"
)

func Add(x, y int) int {
	result := x + y
	return result
}

func AddFloat(x, y float64) float64 {
	result := x + y
	return result
}

func Sub(x, y int) int {
	result := x - y
	return result
}

func SubFloat(x, y float64) float64 {
	result := x - y
	return result
}

func Mul(x, y int) int {
	result := x * y
	return result
}

func MulFloat(x, y float64) float64 {
	result := x * y
	return result
}

func Div(x, y int) float64 {
	if y == 0 {
		fmt.Println("cannot divide by zero")
		os.Exit(0)
	}
	result := float64(x) / float64(y)
	return result
}

func DivFloat(x, y float64) float64 {
	if y == 0.0 {
		fmt.Println("cannot divide by zero")
		os.Exit(0)
	}
	result := x / y
	return result
}

func Exp(x, y float64) float64 {
	result := math.Pow(x, y)
	return result
}
