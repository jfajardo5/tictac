/*
MIT License
Copyright (c) 2023 Julio "jfajardo5" Fajardo

See end of file for extended license information
*/

package main

import "fmt"

// Runtime
func main() {
	game := New()
	game.initPlayers()
	for !game.Over {
		game.DrawBoard()
		game.PlayRound()
		game.Rounds++
	}
	fmt.Println("\n\nGame Over!")
	fmt.Println(game.Result)
}

// Game container
type TicTacToe struct {
	Board   [9]string
	Rounds  int
	Over    bool
	Player1 string
	Player2 string
	Result  string
}

// TicTacToe init function
func New() *TicTacToe {
	return &TicTacToe{
		Board:  initBoard(),
		Rounds: 0,
		Over:   false,
	}
}

// Checks for determining game status
func (t *TicTacToe) IsGameOver() bool {
	// Player 1 checks
	if t.verticalVictory(t.Player1) || t.horizontalVictory(t.Player1) || t.diagonalVictory(t.Player1) {
		t.Result = "Player 1 wins!"
		t.Over = true
	}

	// Player 2 checks
	if t.verticalVictory(t.Player2) || t.horizontalVictory(t.Player2) || t.diagonalVictory(t.Player2) {
		t.Result = "Player 2 wins!"
		t.Over = true
	}

	if t.noSquaresLeft() {
		t.Result = "No more moves left. This is a draw!"
		t.Over = true
	}

	return t.Over
}

// Checks to see if moves can still be made
func (t *TicTacToe) noSquaresLeft() bool {
	for i := range t.Board {
		if t.Board[i] != t.Player1 && t.Board[i] != t.Player2 {
			return false
		}
	}
	return true
}

// Check for horizontal win condition
func (t *TicTacToe) horizontalVictory(player string) bool {
	for i := range t.Board {
		if (i+1)%3 == 0 {
			if t.Board[i] == player && t.Board[i-1] == player && t.Board[i-2] == player {
				return true
			}
		}
	}
	return false
}

// Check for vertical win condition
func (t *TicTacToe) verticalVictory(player string) bool {
	for i := 0; i < 3; i++ {
		if t.Board[i] == player && t.Board[i+3] == player && t.Board[i+6] == player {
			return true
		}
	}
	return false
}

// Check for diagonal win condition
func (t *TicTacToe) diagonalVictory(player string) bool {
	if t.Board[0] == player && t.Board[4] == player && t.Board[8] == player {
		return true
	}
	if t.Board[2] == player && t.Board[4] == player && t.Board[6] == player {
		return true
	}
	return false
}

// Logic for a game round
// TODO re-factor this. Don't like the repetition implemented here
func (t *TicTacToe) PlayRound() {

	player1Selection := 0
	for !t.ValidateInput(player1Selection) {
		fmt.Println("\nPlayer 1, please select a numbered square:")
		fmt.Scan(&player1Selection)
	}
	t.Board[player1Selection-1] = t.Player1
	t.DrawBoard()
	if t.IsGameOver() {
		return
	}

	player2Selection := 0
	for !t.ValidateInput(player2Selection) {
		fmt.Println("\nPlayer 2, please select a numbered square:")
		fmt.Scan(&player2Selection)
	}
	t.Board[player2Selection-1] = t.Player2
	t.IsGameOver()
}

// Initialize player values
// Gets input from players and assigns their marks
func (t *TicTacToe) initPlayers() {
	var player1 string
	fmt.Println("Player 1: Tic or Tac?")
	fmt.Println("Enter X to use X")
	fmt.Println("Enter anything else to use O")
	fmt.Scanln(&player1)
	t.Player1 = "[O]"
	t.Player2 = "[X]"
	if player1 == "X" || player1 == "x" {
		t.Player1 = "[X]"
		t.Player2 = "[O]"
	}
	fmt.Printf("Player 1 is %s \n", t.Player1)
	fmt.Printf("Player 2 is %s \n", t.Player2)
}

// Print the board
func (t *TicTacToe) DrawBoard() {
	fmt.Printf("\n*-----Round %d-----*\n\n", t.Rounds)
	for i := range t.Board {
		fmt.Printf("%s ", t.Board[i])
		if (i+1)%3 == 0 {
			fmt.Println()
		}
	}
}

// Initialize game board
func initBoard() [9]string {
	return [9]string{"[1]", "[2]", "[3]", "[4]", "[5]", "[6]", "[7]", "[8]", "[9]"}
}

// Validate player input.
// This ensures players enter a numbered field within the valid range
// and that their selected field hasn't already been marked
func (t *TicTacToe) ValidateInput(input int) bool {
	if input < 1 || input > 9 {
		return false
	}
	if t.Board[input-1] == t.Player1 || t.Board[input-1] == t.Player2 {
		return false
	}
	return true
}

/*

MIT License

Copyright (c) 2023 Julio Fajardo

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.

*/
