package main

import (
	"fmt"
	pkgA "github.com/liangtsao/demo-go-multi-modules/moduleA/pkg"
	pkgB "github.com/liangtsao/demo-go-multi-modules/moduleB/pkg"
)

func main() {
	calc := pkgA.NewCalculator()
	fmt.Println("10 + 5 =", calc.Add(5, 1))
	fmt.Println("10 - 5 =", calc.Subtract(5, 1))

	formatter := pkgB.NewStringFormatter()
	fmt.Println(formatter.UpperCase("multi-module sadasdada demo"))
}