package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

// Generate random no 1 - 100
// Ask user to guess
// if the guess is less than target : Guess is too low else guess is too high
// Give 10 chances
// If successfully guessed - "Good job!"
// else on 10 tries - "Game Over - The target was ..."

func main() {
	// Generate random no 1 - 100
	s1 := rand.NewSource(int64(time.Now().Second()))
	r1 := rand.New(s1)
	target := r1.Intn(100) + 1
	//fmt.Println(target)

	// Ask user to guess
	fmt.Println("Generate the random number (1 to 100).  Can you guess it ?")
	reader := bufio.NewReader(os.Stdin)
	for i := 0; i < 10; i++ {
		fmt.Println("Enter your guestimate: ")
		guess, _ := reader.ReadString('\n')
		guess = strings.TrimSpace(guess)
		guessint, _ := strconv.Atoi(guess)
		if guessint == target {
			fmt.Println("Sahi jawaab")
			break
		} else {
			if i == 9 {
				fmt.Println("Game Over! The target was ", target)
				break
			}
			if guessint < target {
				fmt.Println("Gosh you guess too low! No worries you have ", 10-i, " tries remaining")
			} else {
				fmt.Println("Gosh you guess too high! No worries you have ", 10-i, " tries remaining")
			}

		}
	}

}
