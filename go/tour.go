package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Hello") // seri kullanımda alt alta yazar

	fmt.Println("The time is ", time.Now())

	fmt.Println("random number is :", rand.Intn(100))

	fmt.Printf("sqrt %d :%g \n", 7, math.Sqrt(7)) //printf içine parametre alabilir

	fmt.Println("pi :", math.Pi)

	//Go da dışa aktarilabilir yapılar büyük harflerle başlamak zorundadır .

}
