package tic_tac_toe

import "errors"

const (
	SIZE = 3

	X = 'x'
	O = 'o'
)

var (
	ERR_GAME_NOT_OVER = errors.New("Game is not finished")
)

type board [SIZE][SIZE]byte

func (b board) String() string {
	out := ""
	for i := 0; i < SIZE; i++ {
		out += "|"
		for j := 0; j < SIZE; j++ {
			out += string(b[i][j]) + "|"
		}
		out += "\n"
	}
	return out
}

func (b board) IsWinner(player byte) bool {
	//Horizontals
	for i := 0; i < SIZE; i++ {
		win := true
		for j := 0; j < SIZE; j++ {
			if b[i][j] != player {
				win = false
				break
			}
		}
		if win {
			return true
		}
	}

	//Verticals
	for j := 0; j < SIZE; j++ {
		win := true
		for i := 0; i < SIZE; i++ {
			if b[i][j] != player {
				win = false
				break
			}
		}
		if win {
			return true
		}
	}

	//Down-Right
	drWin := true
	for i := 0; i < SIZE; i++ {
		if b[i][i] != player {
			drWin = false
			break
		}
	}
	if drWin {
		return true
	}

	//Up-Right
	urWin := true
	for i := 0; i < SIZE; i++ {
		if b[i][SIZE-i-1] != player {
			urWin = false
			break
		}
	}
	return urWin
}

type move struct {
	y, x int
}

type Game struct {
	b     board
	oTurn bool
}

func (g Game) String() string {
	return g.b.String()
}

func (g Game) IsTerminalState() bool {
	_, err := g.Winner()
	return err == nil
}

func (g Game) Winner() (byte, error) {
	if g.b.IsWinner(X) {
		return X, nil
	} else if g.b.IsWinner(O) {
		return O, nil
	}

	return '_', ERR_GAME_NOT_OVER
}

func (g Game) GetActions() (moves []move) {
	for i := 0; i < SIZE; i++ {
		for j := 0; j < SIZE; j++ {
			if g.b[i][j] == '_' {
				move := move{
					i, j,
				}
				moves = append(moves, move)
			}
		}
	}
	return
}

func (g Game) ApplyAction(m move) (Game, error) {
	var playerToMove byte = X
	if g.oTurn {
		playerToMove = O
	}
	g.b[m.y][m.x] = playerToMove
	g.oTurn = !g.oTurn
	return g, nil
}

func NewGame() Game {
	var g Game
	for i := 0; i < SIZE; i++ {
		for j := 0; j < SIZE; j++ {
			g.b[i][j] = '_'
		}
	}
	return g
}
