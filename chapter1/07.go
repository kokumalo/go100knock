package main

import (
	"flag"
	"fmt"
)

var (
	x int
	y string
	z float64
)

func main() {
	flag.IntVar(&x, "x", 0, "時間")
	flag.StringVar(&y, "y", "気温", "文字列")
	flag.Float64Var(&z, "z", 0.0, "数字")
	flag.Parse()

	fmt.Println(x, "時の", y, "は", z)
}
