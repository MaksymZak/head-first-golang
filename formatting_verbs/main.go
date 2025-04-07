package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello formatting verbs")

	width := 1.2
	height := 3.0
	area := width * height
	fmt.Printf("Area is %0.2f square meters\n", area)

	fmt.Printf("The %s cost  %d cent each\n", "apples", 100)
	fmt.Printf("That will be $%0.1f please\n", 0.23*5)

	// main the formatting verbs
	// %f - Число с плавающей точкой
	// %d - Целое число
	// %s - Строка
	// %v - Значение по умолчанию
	// %t - Булевое значение
	// %T - Тип значения
	// %% - Процентный знак

	fmt.Printf("The area is %0.2f square meters\n", area)
	fmt.Printf("The area is %d square meters\n", 15)
	fmt.Printf("The area is %s square meters\n", "15")
	fmt.Printf("The area is %v square meters\n", "any value")
	fmt.Printf("The area is %t square meters\n", true)
	fmt.Printf("The area is %T square meters\n", "this is a string")
	fmt.Printf("The area is %% square meters\n")

	// %#v - Вывод значения в виде кода Go

	fmt.Printf("%12s | %2d\n", "apples", 100)
	fmt.Printf("%12s | %2d\n", "oranges", 200)
	fmt.Printf("%12s | %2d\n", "bananas", 300)

}
