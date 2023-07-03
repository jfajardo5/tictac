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
		for i := range game.Players {
			game.DrawBoard()
			game.PlayRound(game.Players[i])
			if game.IsGameOver() {
				break
			}
		}
		game.Rounds++
	}
	game.DrawBoard()
	fmt.Println("\n\nGame Over!")
	fmt.Println(game.Result)
}

// Game container
type TicTacToe struct {
	Board   [9]string
	Rounds  int
	Over    bool
	Result  string
	Players []Player
}

// Player object
type Player struct {
	Name   string
	Marker string
}

// TicTacToe init function
func New() *TicTacToe {
	return &TicTacToe{
		Board:  initBoard(),
		Rounds: 1,
		Over:   false,
	}
}

// Checks for determining game status
func (t *TicTacToe) IsGameOver() bool {
	for i := range t.Players {
		if t.verticalVictory(t.Players[i].Marker) || t.horizontalVictory(t.Players[i].Marker) || t.diagonalVictory(t.Players[i].Marker) {
			t.Result = fmt.Sprintf("%s wins!", t.Players[i].Name)
			t.Over = true
		}
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
		if t.Board[i] != "[X]" && t.Board[i] != "[O]" {
			return false
		}
	}
	return true
}

// Check for horizontal win condition
func (t *TicTacToe) horizontalVictory(marker string) bool {
	for i := range t.Board {
		if (i+1)%3 == 0 {
			if t.Board[i] == marker && t.Board[i-1] == marker && t.Board[i-2] == marker {
				return true
			}
		}
	}
	return false
}

// Check for vertical win condition
func (t *TicTacToe) verticalVictory(marker string) bool {
	for i := 0; i < 3; i++ {
		if t.Board[i] == marker && t.Board[i+3] == marker && t.Board[i+6] == marker {
			return true
		}
	}
	return false
}

// Check for diagonal win condition
func (t *TicTacToe) diagonalVictory(marker string) bool {
	if t.Board[0] == marker && t.Board[4] == marker && t.Board[8] == marker {
		return true
	}
	if t.Board[2] == marker && t.Board[4] == marker && t.Board[6] == marker {
		return true
	}
	return false
}

// Logic for a player's turn
func (t *TicTacToe) PlayRound(player Player) {
	selection := 0
	for !t.ValidateInput(selection) {
		fmt.Printf("\n%s, please select a numbered square:\n", player.Name)
		fmt.Scan(&selection)
	}
	t.Board[selection-1] = player.Marker
}

// Initialize player values
// Gets input from players and assigns their marks
func (t *TicTacToe) initPlayers() {
	var name string
	for i := 0; i < 2; i++ {
		fmt.Printf("Player %d: Please enter your name.\n", i+1)
		fmt.Scanln(&name)
		if i == 0 {
			t.Players = append(t.Players, Player{Name: name, Marker: getMarker(name)})
		} else {
			t.Players = append(t.Players, Player{Name: name, Marker: t.getOppositeMarker()})
		}
	}
}

// Returns the marker Player 1 did not pick
func (t *TicTacToe) getOppositeMarker() string {
	for i := range t.Players {
		if t.Players[i].Marker == "[O]" {
			return "[X]"
		}
	}
	return "[O]"
}

// Gets input for marker to use from Player 1
func getMarker(player string) string {
	marker := 0
	fmt.Printf("%s, choose your marker:\n", player)
	for marker != 1 && marker != 2 {
		fmt.Println("Enter 1 to use O")
		fmt.Println("Enter 2 to use X")
		fmt.Scan(&marker)
	}
	if marker == 1 {
		return "[O]"
	}
	return "[X]"
}

// Draw the board
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
	for i := range t.Players {
		if t.Board[input-1] == t.Players[i].Marker {
			return false
		}
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
