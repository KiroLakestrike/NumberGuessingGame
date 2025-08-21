/*

This is just a small Project to get used to Go.

*/

package main

import (
	"fmt"
)

// Ruleset for the Game
// The game will randomize a number between 1 and X (depending on difficulty)

func main() {

	difficulty := 0
	fmt.Println("********************************************************")
	fmt.Println("* Welcome to my tiny little number guessing game:      *")
	fmt.Println("********************************************************")
	fmt.Println("* If you want some help, you can type HELP at any time *")

	if difficulty == 0 {
		difficulty = loadDifficulty()
	}

	fmt.Println("Your selected difficulty is:", difficulty)

}

// TODO: offload this into its own file. And add logic to prevent the user from selecting wrong numbers. Maybe use a Switch statement for this.
func loadDifficulty() int {
	current := 0

	fmt.Print("Type a number between 1 & 3: ")
	fmt.Scan(&current)

	if current < 1 || current > 3 {
		fmt.Print("Sorry, this is not a valid number 1, 2 or 3 is possible, Try Again: ")
		fmt.Scan(&current)
	}

	return current
}
