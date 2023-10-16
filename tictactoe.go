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
	players []*player // todo: should this just be fields player1 & 2?
	// TODO; need to deprecate myTurn field
	// true if computer's turn, false for user's turn
	//myTurn bool
	// the 3 rows of the board
	row1, row2, row3 row
	// keeps count of # of games, won by user, won by comp, and ties
	// todo: am i keeping count inside the players now? i believe so 
	totalgames, userct, compct, tiect int
	// number spots available
	emptySpots int
}

type row struct {
	// todo: an alternative could be a map, would remove all the switchcase stuff
	p1, p2, p3 string
}

type player struct {
	// what should I keep here? 
	name string
	computer bool // if this player is a computer or not
	symbol string // symbol representing this player on the board
	wins int // num wins for this user
	opp *player // player's opponent
}

func newplayer(iscomp bool, sym string, name string) player {
	var p player
	p.computer = iscomp
	p.symbol = sym
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
		fmt.Println("found a winner, across the row")
		return true
	}
	if (t.row1.p1 == t.row2.p1 && t.row2.p1 == t.row3.p1 && t.row1.p1 != " ") || (t.row1.p2 == t.row2.p2 && t.row2.p2 == t.row3.p2 && t.row1.p2 != " ") || (t.row1.p3 == t.row2.p3 && t.row2.p3 == t.row3.p3 && t.row1.p3 != " ") {
		fmt.Println("found a winner, across a column")
		return true
	}
	if (t.row1.p1 == t.row2.p2 && t.row2.p2 == t.row3.p3) || (t.row1.p3 == t.row2.p2 && t.row2.p2 == t.row3.p1) {
		if t.row2.p2 != " " {
			fmt.Println("found a winner, across a diagonal")
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

// todo; use getEmptySpots to check if spots empty instead of these fns
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

func (t *tboard) markSpot(spot string, sym string) {
	// ex: 1A, X marks spot 1A with 'X'
	// assumes valid input
	s := strings.Split(spot, "")
	r , c := s[0], s[1]
//	fmt.Printf("%+v\n", t)
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
	//fmt.Printf("%+v\n", r)
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
	fmt.Println("\t 1 |", "1A", "|", "1B", "|", "1C", "|")
	fmt.Println("\t 2 |", "2A", "|", "2B", "|", "2C", "|")
	fmt.Println("\t 3 |", "3A", "|", "3B", "|", "3C", "|")
	fmt.Println("\t-------------------\n")
} 


func (t *tboard) randomSetTurn() {
	// todo: take into account 2 players, assign turn randomly
	rand.Seed(time.Now().UnixNano())
	rn := rand.Intn(2)
	//if rn == 1 {
	//	t.myTurn = true
	//} else {
	//	t.myTurn = false
	//}
	// todo: remove myTurn code
	t.nextPlayer = t.players[rn]
}

func main() {
	// TODO: give an option for the user to choose to play the game by themselves or for the computer to play with them
	// choose player 1: user or computer
	// choose player 2: user or computer
	// then let them play against each other. maybe need another struct for player
	fmt.Println("\n\nxoxoxoxoxoxoxoxxoxoxoxo Tic Tac Toe oxoxoxooxoxoxoxoxoxoxoxoxo")

	fmt.Println("\nINSTRUCTIONS")
	fmt.Println("\tAs the user, you will play tic tac toe against the program.")
	fmt.Println("\tA random coin flip will decide who goes first.")
	fmt.Println("\tWhen asked for input, you will type in the location on the board you wish to play.")
	fmt.Println("\tSee the example board for valid input corresponding to a location:")
	printExampleBoard()
	fmt.Println("\noxoxoxoxoxoxoxoxoxoxoxoxoxoxoxoxoxoxoxoxoxoxoxoxoxooxoxoxoxoxoxo\n")

	t := newboard() 
	var player1, player2 player

	// todo: ask user if they want a computer or not
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
		// todo: ask user for their name
		player2 = newplayer(false, SYMO, "user")
	} else if nu == 2 {
		player1 = newplayer(false, SYMX, "user1")
		player2 = newplayer(false, SYMO, "user2")
	}
	player1.opp = &player2
	player2.opp = &player1

	plyrs := []*player{&player1, &player2}
	t.players = plyrs

	t.randomSetTurn()
	//if t.myTurn {
	//	fmt.Println("\tCoin flipped, Computer goes first")
	//} else {
	//	fmt.Println("\tCoin flipped, User goes first.")
	//}
	fmt.Printf("\nPlayer to go first is: %v", &t.nextPlayer)
	fmt.Printf("\nPlayers associated with the board: %v\n", t.players)

	fmt.Println("\tThe user will be symbolized by '", USR, "' and the computer will be symbolized by '", CMP, "'")
	//t.computerTakesATurn()
	//t.printBoard()
	for true {
		//t.computerTakesATurn()
		t.printBoard()	

		if t.checkIfWin() {
			fmt.Println("game over, there is a winner")
			break
			// check who the winner is depending on who went last
			// increment numGames, incrememnt for winner
		} else if t.checkIfFull() {
			fmt.Println("game over, the board is full. Tie.")
			// incrememnet numGames, ties
			break
		}

		// next player takes their turn
		if t.nextPlayer.computer {
			rand.Seed(time.Now().UnixNano())
			spots := t.getEmptySpots()
			rn := rand.Intn(len(spots))
			spot := spots[rn]
			t.markSpot(spot, t.nextPlayer.symbol)

		} else {
			// ask for user input
			var spot string
			fmt.Println("Please enter the location you wish to play")
			fmt.Scanln(&spot)
			// todo: confirm spot is empty, else ask again
			for !t.checkIfSpotEmpty(spot) {
				fmt.Println("That spot is taken, please choose another")
				fmt.Scanln(&spot)
			}
			t.markSpot(spot, t.nextPlayer.symbol)	
		}
		// set the next player's turn
		t.nextPlayer = t.nextPlayer.opp

		// ask if user wants to play again
		// if yes, clear the board and carry on in the loop. might need to do some of the startup again like random coin flip
		// if not play again, exit the for loop
	}

	fmt.Println("\n\n Thanks for playing! Goodbye!")
	fmt.Println("\noxoxoxoxoxoxoxoxoxoxoxoxoxoxoxoxoxoxoxoxoxoxoxoxoxooxoxoxoxoxoxo\n")
}


