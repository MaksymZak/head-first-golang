// guess - A simple number guessing game

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
	count := 10
	rand.NewSource(time.Now().Unix())
	target := rand.Intn(100) + 1
	fmt.Print("Chose a number between 1 and 100: ")

	reader := bufio.NewReader(os.Stdin)

	success := false

	for guesses := 0; guesses < count; guesses++ {
		fmt.Println("You have", count-guesses, "guesses left")
		fmt.Print("Make a guess: ")
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
			success = true
			fmt.Println("Good job! You guessed it!")
			break
		}
	}

	if !success {
		fmt.Println("You lose. The correct number was", target)
	}

}
