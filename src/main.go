package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/ayushknath/calc-go/src/compute"
	"github.com/ayushknath/calc-go/src/interactive"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("shortage of arguments")
		os.Exit(0)
	}

	if os.Args[1] == "-i" {
		interactive.InteractiveMode()
		return
	}

	args := os.Args[1:]

	if len(args) < 3 {
		fmt.Println("shortage of arguments")
		os.Exit(0)
	}

	var x, y int
	var xFloat, yFloat float64
	var errX, errY, errXFloat, errYFloat error
	isFloat := strings.Contains(args[1], ".") || strings.Contains(args[2], ".")
	if isFloat {
		xFloat, errXFloat = strconv.ParseFloat(args[1], 64)
		yFloat, errYFloat = strconv.ParseFloat(args[2], 64)
	} else {
		x, errX = strconv.Atoi(args[1])
		y, errY = strconv.Atoi(args[2])
	}
	if errXFloat != nil || errYFloat != nil || errX != nil || errY != nil {
		fmt.Println("invalid number")
		os.Exit(0)
	}

	switch args[0] {
	case "add":
		if isFloat {
			fmt.Printf("%v + %v = %.2f\n", args[1], args[2], compute.AddFloat(xFloat, yFloat))
		} else {
			fmt.Printf("%v + %v = %v\n", args[1], args[2], compute.Add(x, y))
		}
	case "sub":
		if isFloat {
			fmt.Printf("%v - %v = %.2f\n", args[1], args[2], compute.SubFloat(xFloat, yFloat))
		} else {
			fmt.Printf("%v - %v = %v\n", args[1], args[2], compute.Sub(x, y))
		}
	case "mul":
		if isFloat {
			fmt.Printf("%v * %v = %.2f\n", args[1], args[2], compute.MulFloat(xFloat, yFloat))
		} else {
			fmt.Printf("%v * %v = %v\n", args[1], args[2], compute.Mul(x, y))
		}
	case "div":
		if isFloat {
			fmt.Printf("%v / %v = %.2f\n", args[1], args[2], compute.DivFloat(xFloat, yFloat))
		} else {
			fmt.Printf("%v / %v = %.2f\n", args[1], args[2], compute.Div(x, y))
		}
	case "exp":
		if isFloat {
			fmt.Printf("%v ** %v = %.3f\n", args[1], args[2], compute.Exp(xFloat, yFloat))
		} else {
			fmt.Printf("%v ** %v = %v\n", args[1], args[2], compute.Exp(float64(x), float64(y)))
		}
	default:
		fmt.Println("invalid operation, please try again")
	}
}
