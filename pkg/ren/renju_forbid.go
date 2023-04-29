package ren

import (
	"fmt"
)

// 棋盘大小
const S = 15

// 执黑
const BLACK_COLOR = 1

const WhITE_COLOR = -1

func CheckForbid(board [S][S]int, x, y int) int {
	copyBoard := board[:]
	line4V := newLine4V(copyBoard)

	if copyBoard[x][y] == 1 {
		return line4V.foulr(x, y)
	} else {
		return 0
	}
}

// 白胜的逻辑（五子连珠）
// 黑胜的逻辑（五子连珠且不包含禁手）
func CheckWin(board [S][S]int, x, y int) int {
	copyBoard := board[:]
	line4V := newLine4V(copyBoard)

	// 黑落在禁手处,直接判白胜
	if CheckForbid(board, x, y) != 0 {
		return WHITE_WIN
	}

	if copyBoard[x][y] == 1 && line4V.hasWon(x, y) {
		return BLACK_WIN
	} else if copyBoard[x][y] == -1 && line4V.hasWon(x, y) {
		return WHITE_WIN
	}

	return NO_RESULT
}

// 是否五子连珠（不包含禁手判断）
func isRenju(board [S][S]int, x, y int) int {
	copyBoard := board[:]
	line4V := newLine4V(copyBoard)
	if copyBoard[x][y] == 1 && line4V.hasWon(x, y) {
		return BLACK_WIN
	} else if copyBoard[x][y] == -1 && line4V.hasWon(x, y) {
		return WHITE_WIN
	}

	return NO_RESULT
}

func PrintBoard(board [S][S]int) {
	for _, v1 := range board {
		for k2, v2 := range v1 {
			if v2 == BLACK_COLOR {
				fmt.Print("X ")
			} else if v2 == WHITE_WIN {
				fmt.Print("O ")
			} else {
				fmt.Print(". ")
			}
			if k2 == S-1 {
				fmt.Println()
			}
		}
	}
}
