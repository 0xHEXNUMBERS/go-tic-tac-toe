package tictactoe

import (
	"fmt"
	"reflect"
	"testing"
)

func testNonTerminalAction(g Game, m Move) error {
	gameCont, err := g.ApplyAction(m)
	if err != nil {
		return fmt.Errorf("Could not apply non-terminal action: %s", err)
	}

	if gameCont.IsTerminalState() {
		return fmt.Errorf("Non-terminal action resulted in a terminal state")
	}

	_, err = gameCont.Winner()
	if err == nil {
		return fmt.Errorf("Non-terminal action resulted in a state with winning players: %s", err)
	}
	return nil
}

func testTerminalAction(g Game, m Move, winner byte) error {
	gameFinish, err := g.ApplyAction(m)
	if err != nil {
		return fmt.Errorf("Could not apply winning action: %s", err)
	}

	fmt.Println(gameFinish)

	if !gameFinish.IsTerminalState() {
		return fmt.Errorf("Winning action did not result in a terminal state")
	}

	player, err := gameFinish.Winner()
	if err != nil {
		return fmt.Errorf("Winning action resulted in an invalid state: %s", err)
	}

	if player != winner {
		return fmt.Errorf("Winning action did not result in the correct player winning: want '%c', got '%c'",
			winner,
			player,
		)
	}

	return nil
}

func mappifyMoves(actions []Move) map[Move]bool {
	actionsCollected := make(map[Move]bool)
	for _, a := range actions {
		actionsCollected[a] = true
	}
	return actionsCollected
}

func containSameMoves(a, b []Move) bool {
	return reflect.DeepEqual(
		mappifyMoves(a), mappifyMoves(b),
	)
}

func testHorizontal(player byte) error {
	g := NewGame()
	g.b[0][0] = player
	g.b[0][1] = player

	if player == O {
		g.oTurn = true
	}

	actionsGot := g.GetActions()

	actionsWant := []Move{
		{0, 2},
		{1, 0},
		{1, 1},
		{1, 2},
		{2, 0},
		{2, 1},
		{2, 2},
	}

	if len(actionsGot) != len(actionsWant) {
		return fmt.Errorf("did not receive expected number of moves: got %d, want %d",
			len(actionsGot),
			len(actionsWant),
		)
	}

	if !containSameMoves(actionsGot, actionsWant) {
		return fmt.Errorf("Actions collected are not the same actions we wanted")
	}

	if err := testTerminalAction(g, actionsWant[0], player); err != nil {
		return fmt.Errorf("%s", err)
	}

	for i := 1; i < len(actionsWant); i++ {
		if err := testNonTerminalAction(g, actionsWant[i]); err != nil {
			return fmt.Errorf("%s", err)
		}
	}

	return nil
}

func TestHorizontal(t *testing.T) {
	if err := testHorizontal(X); err != nil {
		t.Errorf("TestHorizontal: X game failed: %s", err)
	}

	if err := testHorizontal(O); err != nil {
		t.Errorf("TestHorizontal: O game failed: %s", err)
	}
}

func testVertical(player byte) error {
	g := NewGame()
	g.b[0][0] = player
	g.b[1][0] = player

	if player == O {
		g.oTurn = true
	}

	actionsGot := g.GetActions()

	actionsWant := []Move{
		{2, 0},
		{0, 1},
		{1, 1},
		{2, 1},
		{0, 2},
		{1, 2},
		{2, 2},
	}

	if len(actionsGot) != len(actionsWant) {
		return fmt.Errorf("did not receive expected number of moves: got %d, want %d",
			len(actionsGot),
			len(actionsWant),
		)
	}

	if !containSameMoves(actionsGot, actionsWant) {
		return fmt.Errorf("Actions collected are not the same actions we wanted")
	}

	if err := testTerminalAction(g, actionsWant[0], player); err != nil {
		return fmt.Errorf("%s", err)
	}

	for i := 1; i < len(actionsWant); i++ {
		if err := testNonTerminalAction(g, actionsWant[i]); err != nil {
			return fmt.Errorf("%s", err)
		}
	}

	return nil
}

func TestVertical(t *testing.T) {
	if err := testVertical(X); err != nil {
		t.Errorf("TestVertical: X game failed: %s", err)
	}

	if err := testVertical(O); err != nil {
		t.Errorf("TestVertical: O game failed: %s", err)
	}
}

func testDownRight(player byte) error {
	g := NewGame()
	g.b[0][0] = player
	g.b[1][1] = player

	if player == O {
		g.oTurn = true
	}

	actionsGot := g.GetActions()

	actionsWant := []Move{
		{2, 2},
		{0, 1},
		{0, 2},
		{1, 0},
		{1, 2},
		{2, 0},
		{2, 1},
	}

	if len(actionsGot) != len(actionsWant) {
		return fmt.Errorf("did not receive expected number of moves: got %d, want %d",
			len(actionsGot),
			len(actionsWant),
		)
	}

	if !containSameMoves(actionsGot, actionsWant) {
		return fmt.Errorf("Actions collected are not the same actions we wanted")
	}

	if err := testTerminalAction(g, actionsWant[0], player); err != nil {
		return fmt.Errorf("%s", err)
	}

	for i := 1; i < len(actionsWant); i++ {
		if err := testNonTerminalAction(g, actionsWant[i]); err != nil {
			return fmt.Errorf("%s", err)
		}
	}

	return nil
}

func TestDownRight(t *testing.T) {
	if err := testDownRight(X); err != nil {
		t.Errorf("TestDownRight: X game failed: %s", err)
	}

	if err := testDownRight(O); err != nil {
		t.Errorf("TestDownRight: O game failed: %s", err)
	}
}

func testUpRight(player byte) error {
	g := NewGame()
	g.b[2][0] = player
	g.b[1][1] = player

	if player == O {
		g.oTurn = true
	}

	actionsGot := g.GetActions()

	actionsWant := []Move{
		{0, 2},
		{0, 0},
		{0, 1},
		{1, 0},
		{1, 2},
		{2, 1},
		{2, 2},
	}

	if len(actionsGot) != len(actionsWant) {
		return fmt.Errorf("did not receive expected number of moves: got %d, want %d",
			len(actionsGot),
			len(actionsWant),
		)
	}

	if !containSameMoves(actionsGot, actionsWant) {
		return fmt.Errorf("Actions collected are not the same actions we wanted")
	}

	if err := testTerminalAction(g, actionsWant[0], player); err != nil {
		return fmt.Errorf("%s", err)
	}

	for i := 1; i < len(actionsWant); i++ {
		if err := testNonTerminalAction(g, actionsWant[i]); err != nil {
			return fmt.Errorf("%s", err)
		}
	}

	return nil
}

func TestUpRight(t *testing.T) {
	if err := testUpRight(X); err != nil {
		t.Errorf("TestUpRight: X game failed: %s", err)
	}

	if err := testUpRight(O); err != nil {
		t.Errorf("TestUpRight: O game failed: %s", err)
	}
}
