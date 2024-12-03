package main

import (
	"fmt" // для вывода в консоль используется пакет fmt
	"math"
	"reflect"
	"strings"
)

// для работы со строками используется пакет strings

func main() {
	fmt.Println("Hello", 100, 200, 300, "aadwawd")
	fmt.Println(math.Ceil(2.2))
	fmt.Println(strings.Title("hello world"))

	//////////////////////////// 03.12.2024 ////////////////////////////
	////////////////////////////golang types////////////////////////////

	// bool
	// string
	// int  int8  int16  int32  int64
	// uint uint8 uint16 uint32 uint64 uintptr
	// byte // alias for uint8
	// rune // alias for int32
	// float32 float64
	// complex64 complex128

	fmt.Println((reflect.TypeOf(1111111111111111111)))
	fmt.Println((reflect.TypeOf(true)))
	fmt.Println((reflect.TypeOf(11.0)))
	fmt.Println((reflect.TypeOf("hello")))

	////////////////////////////golang variables////////////////////////////
	var a int
	var b int
	b, a = 22, 10

	sum := a + b

	fmt.Println(b, a, "sum is", sum)

	price := 100
	fmt.Println("Price is", price, "USD")

	taxRate := 0.08
	tax := float64(price) * taxRate
	fmt.Println("Tax is", tax, "USD")

	total := price + int(tax)
	fmt.Println("Total cost is", total, "USD")

	availableFunds := 120.95
	fmt.Println(availableFunds, "USD available")
	fmt.Println("Within budget?", total <= int(availableFunds))

}
