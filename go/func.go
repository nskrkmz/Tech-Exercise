package main

import (
	"fmt"
)

// ardışık tanımlamalarda ortak type olarak kullanılabilir
// func add(x, y int) int {...}
func add(x int, y int) int { // func foo (bar type , ...) returnType{}
	return x + y
}

func main() {
	tut := add(12, 15)
	fmt.Printf("toplam = %d\n", tut)

}
