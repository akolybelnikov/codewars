package ttt_checker

const (
	X = 1 << iota
	O
)

// winners are decimals of the binaries of all the winning combinations
var winners = []uint16{7, 56, 73, 84, 146, 273, 292, 448}

// complete is a decimal board without 0's, i.e. binary 111111111
var complete uint16 = 511

func IsSolved(board [3][3]int) int {
	players := make(map[int]uint16)
	b := uint16(1)
	for i, row := range board {
		for j, col := range row {
			idx := j + i*3
			players[col] |= b << idx
		}
	}

	for _, winner := range winners {
		if valO, exists := players[O]; exists {
			if valO&winner == winner {
				return O
			}
		}
		if valX, exists := players[X]; exists {
			if valX&winner == winner {
				return X
			}
		}

	}
	if players[O]|players[X] != complete {
		return -1
	}
	return 0
}
