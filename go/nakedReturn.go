package main

import "fmt"

//naked return uzun işlevlerde okunabilirliği azaltır
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	//a, b := split(41)
	//fmt.Println(a, b)
	fmt.Println(split(41))
}
