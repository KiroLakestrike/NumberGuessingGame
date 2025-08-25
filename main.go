/*

This is just a small Project to get used to Go.

*/

package main

import (
	"fmt"
	"math/rand/v2"
)

// Ruleset for the Game
// The game will randomize a number between 1 and X (depending on difficulty)

func main() {

	difficulty := 0
	fmt.Println("********************************************************")
	fmt.Println("* Welcome to my tiny little number guessing game:      *")
	fmt.Println("********************************************************")

	if difficulty == 0 {
		difficulty = loadDifficulty()
	}

	fmt.Println("Your selected difficulty is:", difficulty)

	// TODO Set Parameters for the given difficulty
	maxNumber := 0

	switch difficulty {
	case 1:
		maxNumber = 10
	case 2:
		maxNumber = 25
	case 3:
		maxNumber = 50
	default:
		maxNumber = 10
	}

	// Game Loop Logic
	playedRounds := 0

	computerPoints := 0
	playerPoints := 0

	// Loop for the Round
	for playedRounds < 3 {
		playerGuess := 0
		playerLives := 5

		currentRandom := getRandomNumber(maxNumber)
		fmt.Printf("Round %v \n", playedRounds+1)
		fmt.Printf("The Game has chosen its number, Please choose yours: ")
		fmt.Scan(&playerGuess)

		// Loop to evaluate the number if player has lives left and if guess was correct
		for playerLives > 0 {
			fmt.Printf("You have chosen the number: %v\n", playerGuess)

			if playerGuess > currentRandom {
				playerLives--
				fmt.Printf("Sorry, your guess was too high, you have %v lives left\n", playerLives)
			} else if playerGuess < currentRandom {
				playerLives--
				fmt.Printf("Sorry, your guess was too low, you have %v lives left\n", playerLives)
			} else {
				playerPoints++
				fmt.Printf("That was correct! 🎉 You guessed the number!\n")
				fmt.Printf("Score is: Computer: %v Player: %v\n", computerPoints, playerPoints)
				break // ✅ end round immediately if correct
			}

			// ✅ ask for a new guess only if they still have lives
			if playerLives > 0 {
				fmt.Print("Please enter a new guess: ")
				fmt.Scan(&playerGuess)
			}
		}

		// Deduct points for computer if player ran out of lives
		if playerLives == 0 {
			computerPoints++
			fmt.Printf("Sorry, this round goes to the computer.\n")
			fmt.Printf("Score is: Computer: %v Player: %v\n", computerPoints, playerPoints)
		}

		playedRounds++
	}

}

func loadDifficulty() int {
	current := 0

	fmt.Print("Type a number between 1 & 3: ")
	fmt.Scan(&current)

	for current < 1 || current > 3 {
		if current < 1 || current > 3 {
			fmt.Print("Sorry, this is not a valid number 1, 2 or 3 is possible, Try Again: ")
			fmt.Scan(&current)
		}
	}

	return current
}

func getRandomNumber(maxNumber int) int {
	randomNumber := rand.IntN(maxNumber)
	if randomNumber == 0 {
		randomNumber++
	}
	return randomNumber
}
