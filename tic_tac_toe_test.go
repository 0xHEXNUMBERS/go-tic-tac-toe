package tic_tac_toe

import (
	"fmt"
	"reflect"
	"testing"
)

func testNonTerminalAction(g Game, m move) error {
	gameCont, err := g.ApplyAction(m)
	if err != nil {
		return fmt.Errorf("Could not apply non-terminal action: %s", err)
	}

	if gameCont.IsTerminalState() {
		return fmt.Errorf("Non-terminal action resulted in a terminal state")
	}

	_, err = gameCont.WinningPlayers()
	if err == nil {
		return fmt.Errorf("Non-terminal action resulted in a state with winning players: %s", err)
	}
	return nil
}

func testTerminalAction(g Game, m move, winner byte) error {
	gameFinish, err := g.ApplyAction(m)
	if err != nil {
		return fmt.Errorf("Could not apply winning action: %s", err)
	}

	fmt.Println(gameFinish)

	if !gameFinish.IsTerminalState() {
		return fmt.Errorf("Winning action did not result in a terminal state")
	}

	player, err := gameFinish.WinningPlayers()
	if err != nil {
		return fmt.Errorf("Winning action resulted in an invalid state: %s", err)
	}

	if player[0] != winner {
		return fmt.Errorf("Winning action did not result in the correct player winning: want '%c', got '%c'",
			winner,
			player[0],
		)
	}

	return nil
}

func mappifyMoves(actions []move) map[move]bool {
	actionsCollected := make(map[move]bool)
	for _, a := range actions {
		actionsCollected[a] = true
	}
	return actionsCollected
}

func containSameMoves(a, b []move) bool {
	return reflect.DeepEqual(
		mappifyMoves(a), mappifyMoves(b),
	)
}

func XWinGame() Game {
	g := NewGame()
	g.b[0][0] = X
	g.b[0][1] = X
	return g
}

func TestXWinGame(t *testing.T) {
	g := XWinGame()
	actionsGot := g.GetActions()

	actionsWant := []move{
		move{
			0, 2, X,
		},
		move{
			1, 0, X,
		},
		move{
			1, 1, X,
		},
		move{
			1, 2, X,
		},
		move{
			2, 0, X,
		},
		move{
			2, 1, X,
		},
		move{
			2, 2, X,
		},
	}

	if len(actionsGot) != len(actionsWant) {
		t.Errorf("XWinGame: did not receive expected number of moves: got %d, want %d",
			len(actionsGot),
			len(actionsWant),
		)
	}

	if !containSameMoves(actionsGot, actionsWant) {
		t.Error("Actions collected are not the same actions we wanted")
	}

	if err := testTerminalAction(g, actionsWant[0], X); err != nil {
		t.Errorf("XWinGame: %s", err)
	}

	for i := 1; i < len(actionsWant); i++ {
		if err := testNonTerminalAction(g, actionsWant[i]); err != nil {
			t.Errorf("XWinGame: %s", err)
		}
	}
}
