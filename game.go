package tictactoe

import "errors"

const (
	//SIZE is the size of the board
	SIZE = 3

	//X represents player X
	X = 'x'

	//O represents player O
	O = 'o'
)

var (
	//ErrGameNotOver error
	ErrGameNotOver = errors.New("Game is not finished")

	//ErrInvalidMove error
	ErrInvalidMove = errors.New("Move is invalid")
)

//Move represents an action made by a given player.
type Move struct {
	y, x int
}

//Game represents the current game state.
type Game struct {
	b     board
	oTurn bool
}

func (g Game) String() string {
	return g.b.String()
}

//IsTerminalState returns whether the game is finished or not.
func (g Game) IsTerminalState() bool {
	_, err := g.Winner()
	return err == nil
}

//Winner returns the winner's ascii value.
//
//Returns ErrGameNotOver if the game is not over.
func (g Game) Winner() (byte, error) {
	if g.b.IsWinner(X) {
		return X, nil
	} else if g.b.IsWinner(O) {
		return O, nil
	}

	return '_', ErrGameNotOver
}

//GetActions returns a list of moves that can be made
//by the current player.
func (g Game) GetActions() (moves []Move) {
	for i := 0; i < SIZE; i++ {
		for j := 0; j < SIZE; j++ {
			if g.b[i][j] == '_' {
				move := Move{
					i, j,
				}
				moves = append(moves, move)
			}
		}
	}
	return
}

//ApplyAction takes a Move and applies the action to the current game state.
//
//Returns the new game state and an error if any occurred.
func (g Game) ApplyAction(m Move) (Game, error) {
	//Is the spot vacant?
	if g.b[m.y][m.x] != '_' {
		return Game{}, ErrInvalidMove
	}

	var playerToMove byte = X
	if g.oTurn {
		playerToMove = O
	}
	g.b[m.y][m.x] = playerToMove
	g.oTurn = !g.oTurn
	return g, nil
}

//Player returns the ascii value of the player
//that is currently deciding a move.
//
//Player returns 'o' if player o is making a move.
//Otherwise, Player return 'x'.
func (g Game) Player() byte {
	if g.oTurn {
		return O
	}
	return X
}

//NewGame returns a valid new game of tic-tac-toe.
func NewGame() Game {
	var g Game
	for i := 0; i < SIZE; i++ {
		for j := 0; j < SIZE; j++ {
			g.b[i][j] = '_'
		}
	}
	return g
}
