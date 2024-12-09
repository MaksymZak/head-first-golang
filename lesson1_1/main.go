package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

func main() {

	fmt.Println("Hello lesson1_1")

	var now time.Time = time.Now()
	fmt.Println("Now is day", now.Day())

	brocken := "G# r#cks!"
	replacer := strings.NewReplacer("#", "o")
	fixed := replacer.Replace(brocken)
	fmt.Println(fixed)

	fmt.Print("Enter a number: ")
	// bufio.NewReader() - создает новый Reader, читающий из os.Stdin
	reader := bufio.NewReader(os.Stdin) // os.Stdin Это стандартный ввод, который обычно представляет собой ввод с клавиатуры в консольном приложении.
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		log.Fatal(err)
	}

	fmt.Println("You entered", input)

}
