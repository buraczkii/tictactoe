package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

const USR = "X"
const CMP = "0"
const SYMX = "X"
const SYMO = "0"



type tboard struct {
	// keeps track of whose turn it is next
	nextPlayer *player
	// keeps track of both players
	players []*player
	// the 3 rows of the board
	row1, row2, row3 row
	// keeps count of # of games, won by user, won by comp, and ties
	// todo: am i keeping count inside the players now? i believe so 
	// todo: could also keep a map of wins, ties, etc.
	totalgames, userct, compct, tiect int
}

type row struct {
	p1, p2, p3 string
}

type player struct {
	name string
	computer bool // if this player is a computer or not
	symbol string // symbol representing this player on the board
	wins int // num wins for this user // todo
	opp *player // ptr to the opponent
}

func newplayer(iscomp bool, sym string, name string) player {
	var p player
	p.computer = iscomp
	p.symbol = sym
	p.name = name
	return p
}

func newrow() row {
	var r row
	r.p1 = " "
	r.p2 = " "
	r.p3 = " "
	return r
}

func newboard() tboard {
	var t tboard
	t.row1 = newrow() 
	t.row2 = newrow() 
	t.row3 = newrow() 
	return t
}

//checks if there is a winner
func (t *tboard) checkIfWin() bool {
	// todo: these checks feel messy, i wonder if there is another way 
	if (t.row1.p1 == t.row1.p2 && t.row1.p2 == t.row1.p3 && t.row1.p1 != " ") || (t.row2.p1 == t.row2.p2 && t.row2.p2 == t.row2.p3 && t.row2.p1 != " ") || (t.row3.p1 == t.row3.p2 && t.row3.p2 == t.row3.p3 && t.row3.p1 != " ") {
		return true
	}
	if (t.row1.p1 == t.row2.p1 && t.row2.p1 == t.row3.p1 && t.row1.p1 != " ") || (t.row1.p2 == t.row2.p2 && t.row2.p2 == t.row3.p2 && t.row1.p2 != " ") || (t.row1.p3 == t.row2.p3 && t.row2.p3 == t.row3.p3 && t.row1.p3 != " ") {
		return true
	}
	if (t.row1.p1 == t.row2.p2 && t.row2.p2 == t.row3.p3) || (t.row1.p3 == t.row2.p2 && t.row2.p2 == t.row3.p1) {
		if t.row2.p2 != " " {
			return true
		}
	}
	return false
}


// returns true if the board is full
func (t *tboard) checkIfFull() bool {
	return t.row1.checkIfFull() && t.row2.checkIfFull() && t.row3.checkIfFull()
}

// returns true if the row is full
func (r *row) checkIfFull() bool {
	return r.p1 != " " && r.p2 != " " && r.p3 != " "
}

// computer randomly chooses an empty spot
func (t *tboard) computerTakesATurn() {
	rand.Seed(time.Now().UnixNano())
	spots := t.getEmptySpots()
	rn := rand.Intn(len(spots))
	spot := spots[rn]
	t.markSpot(spot, CMP)
}

// returns the empty spots available for play on the board
func (t *tboard) getEmptySpots() []string {
	var spots []string
	r1s := t.row1.getEmptySpots()
	for _, c := range r1s {
		spots = append(spots, "1" + c)
	}
	r2s := t.row2.getEmptySpots()
	for _, c := range r2s {
		spots = append(spots, "2" + c)
	}
	r3s := t.row3.getEmptySpots()
	for _, c := range r3s {
		spots = append(spots, "3" + c)
	}
	return spots
}

// returns the empty spots in a given row
func (r *row) getEmptySpots() []string {
	var spots []string
	if r.p1 == " " {
		spots = append(spots, "A")
	}
	if r.p2 == " " {
		spots = append(spots, "B")
	}
	if r.p3 == " " {
		spots = append(spots, "C")
	}
	return spots
}

// checks if a spot on the board is empty
func (t *tboard) checkIfSpotEmpty(loc string) bool {
	s := strings.Split(loc, "")
	r := s[0]
	c := s[1]
	switch r {
		case "1":
			return t.row1.checkIfSpotEmpty(c)
		case "2":
			return t.row2.checkIfSpotEmpty(c)
		case "3":
			return t.row3.checkIfSpotEmpty(c)
	}
	return false
}

// checks if a spot in a row is empty
func (r *row) checkIfSpotEmpty(spot string) bool {
	switch spot {
		case "A":
			return r.p1 == " "
		case "B":
			return r.p2 == " "
		case "C":
			return r.p3 == " "
	}
	return false
}

// ex: 1A, X marks spot 1A with 'X'
// assumes valid input
func (t *tboard) markSpot(spot string, sym string) {
	s := strings.Split(spot, "")
	r , c := s[0], s[1]
	switch r {
		case "1":
			t.row1.markSpot(c, sym)
		case "2":
			t.row2.markSpot(c, sym)
		case "3":
			t.row3.markSpot(c, sym)
	}
}

func (r *row) markSpot(spot string, sym string) {
	switch spot {
		case "A":
			r.p1 = sym
		case "B" :
			r.p2 = sym
		case "C":
			r.p3 = sym
	}
}


// resets the board
func (t *tboard) reset() {
	t.row1.reset()
	t.row2.reset()
	t.row3.reset()
}

// resets a row
func (r *row) reset() {
	r.p1 = " "
	r.p2 = " "
	r.p3 = " "
}

// prints out the board's contents
func (t *tboard) printBoard() {
	fmt.Println("\n\t-----A---B---C--")
	fmt.Println("\t 1 |", t.row1.p1, "|", t.row1.p2, "|", t.row1.p3, "|")
	fmt.Println("\t 2 |", t.row2.p1, "|", t.row2.p2, "|", t.row2.p3, "|")
	fmt.Println("\t 3 |", t.row3.p1, "|", t.row3.p2, "|", t.row3.p3, "|")
	fmt.Println("\t-----------------\n")
}

func printExampleBoard() {
	fmt.Println("\n\t-----A----B----C--")
	fmt.Println("\t 1 | 1A | 1B | 1C |")
	fmt.Println("\t 2 | 2A | 2B | 2C |")
	fmt.Println("\t 3 | 3A | 3B | 3C |")
	fmt.Println("\t-------------------\n")
} 

// flips a coin to decide who goes first
func (t *tboard) randomSetTurn() {
	rand.Seed(time.Now().UnixNano())
	rn := rand.Intn(2)
	t.nextPlayer = t.players[rn]
}

// main program
func main() {
	fmt.Println("\n\nxoxoxoxoxoxoxoxxoxoxoxo Tic Tac Toe oxoxoxooxoxoxoxoxoxoxoxoxo")

	fmt.Println("\nINSTRUCTIONS")
	fmt.Println("\tYou get to decide how many human players there will be: 0, 1, or 2. The remaining players will be computer players.")
	fmt.Println("\tA random coin flip will decide who goes first.")
	fmt.Println("\tWhen asked for input, you will type in the location on the board you wish to play.")
	fmt.Println("\tSee the example board for valid input corresponding to a location:")
	printExampleBoard()
	fmt.Println("\noxoxoxoxoxoxoxoxoxoxoxoxoxoxoxoxoxoxoxoxoxoxoxoxoxooxoxoxoxoxoxo\n")

	t := newboard() 
	var player1, player2 player

	fmt.Println("How many human users are playing? Pick 0, 1, or 2.")
	var input string
	fmt.Scanln(&input)

	nu, err := strconv.Atoi(input)
	if err != nil || nu < 0 || nu > 2 {
		fmt.Println("Error reading input, using default 1 human player vs 1 computer player")
	}
	if nu == 0 {
		player1 = newplayer(true, SYMX, "sus1")
		player2 = newplayer(true, SYMO, "sus2")
	} else if nu == 1 {
		player1 = newplayer(true, SYMX, "sus")
		fmt.Println("What is your name?")
		var name string
		fmt.Scanln(&name)
		player2 = newplayer(false, SYMO, name)
	} else if nu == 2 {
		var name1, name2 string
		fmt.Println("What is the first human's name?")
		fmt.Scanln(&name1)
		fmt.Println("What is the second human's name?")
		fmt.Scanln(&name2)
		player1 = newplayer(false, SYMX, name1)
		player2 = newplayer(false, SYMO, name2)
	}
	player1.opp = &player2
	player2.opp = &player1

	plyrs := []*player{&player1, &player2}
	t.players = plyrs
	
	for _, p := range t.players {
		fmt.Println("Player ", p.name, " will be symbolized by ", p.symbol)
	}

	t.randomSetTurn()
	fmt.Printf("\nFlipped a coin. Player to go first is: %s", t.nextPlayer.name)
	time.Sleep(1 * time.Second)

	for true {
		t.printBoard()	

		if t.checkIfWin() || t.checkIfFull() {
			fmt.Println("GAME OVER")
			t.totalgames++
			// store win or full as a field so dont have to run code again?
			if t.checkIfWin() {
				// todo: pretty print a player's name? helper fn?
				fmt.Println("There is a WINNER! Congrats ", t.nextPlayer.opp.name, "!")
				// todo: create a map of user name to #wins	
			} else if t.checkIfFull() {
				fmt.Println("The board is full. This game is a TIE.")
				t.tiect++
			}

			fmt.Println("Would you like to play again? y/n \n(If you wish to change the number of humans vs computer players, please exit the program and restart.)")
			var again string
			fmt.Scanln(&again)
			if again == "y" || again == "yes" || again == "YES" || again == "Y" {
				// todo: winner goes first, flip coin in case of tie
				fmt.Printf("\n\nPlaying again! Resetting board. Flipping a coin to see who goes first.\n")
				t.reset()
				t.randomSetTurn()
				fmt.Printf("\nPlayer to go first is: %v\n\n", t.nextPlayer)
				continue
			} else {
				break
			}

		}

		// next player takes their turn
		if t.nextPlayer.computer {
			fmt.Println("User [", t.nextPlayer.name, "]'s turn to play. Automated.")
			time.Sleep(1*time.Second)
			rand.Seed(time.Now().UnixNano())
			spots := t.getEmptySpots()
			rn := rand.Intn(len(spots))
			spot := spots[rn]
			t.markSpot(spot, t.nextPlayer.symbol)

		} else {
			fmt.Println("User [", t.nextPlayer.name, "]'s turn to play.")
			var spot string
			fmt.Println("Please enter the location you wish to play:")
			fmt.Scanln(&spot)
			for !t.checkIfSpotEmpty(spot) {
				fmt.Println("That spot is taken, please choose another")
				fmt.Scanln(&spot)
			}
			t.markSpot(spot, t.nextPlayer.symbol)	
		}
		// set the next player's turn
		t.nextPlayer = t.nextPlayer.opp

	}

	// todo: pretty print the results of the game here
	fmt.Println("\n\n Exiting game ...\n\n Thanks for playing! Goodbye!")
	fmt.Println("\noxoxoxoxoxoxoxoxoxoxoxoxoxoxoxoxoxoxoxoxoxoxoxoxoxooxoxoxoxoxoxo\n")
}


