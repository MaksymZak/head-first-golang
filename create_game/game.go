package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {

	rand.NewSource(time.Now().Unix())
	target := rand.Intn(100) + 1
	fmt.Print("Chose a number between 1 and 100: ")

	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')

	if err != nil {
		log.Fatal(err)
	}

	input = strings.TrimSpace(input)
	guess, err := strconv.Atoi(input)

	if err != nil {
		log.Fatal(err)
	}

	if guess < target {
		fmt.Println("Oops. Your guess was LOW")
	} else if guess > target {
		fmt.Println("Oops. Your guess was HIGH")
	} else {
		fmt.Println("Your guess is correct")
	}

	for i := 2; i <= 3; i++ {
		fmt.Println("You have", i, "guesses left")
	}
}
