package tictactoe

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
