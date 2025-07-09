package main

import (
	"fmt"
)

func printCelsius() {
	c := (boilingF - 32) * 5 / 9 // 可以访问boilingF
	fmt.Println("Boiling point in Celsius:", c)
}
