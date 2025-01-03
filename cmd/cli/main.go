package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/fabiosoliveira/calculator/internal/math"
)

func main() {

	args := os.Args[1:]

	error := "Você precisa passar os argumentos: <method: sum, sub, div, mult > <parametro1> <parametro2>"
	if len(args) != 3 {
		panic(error)
	}

	arg1, err1 := strconv.ParseFloat(args[1], 64)
	arg2, err2 := strconv.ParseFloat(args[2], 64)
	if err1 != nil || err2 != nil {
		panic("Os parâmetros precisam ser números válidos")
	}

	switch args[0] {
	case "sum":
		fmt.Println(math.Add(arg1, arg2))
	case "sub":
		fmt.Println(math.Subtract(arg1, arg2))
	case "div":
		fmt.Println(math.Divide(arg1, arg2))
	case "mult":
		fmt.Println(math.Multiply(arg1, arg2))
	default:
		panic(error)
	}
}
