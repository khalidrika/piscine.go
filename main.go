package main

import (
	"fmt"
	"os"
)

const (
	Size  = 9
	Empty = '.'
)

func main() {
	args := os.Args[1:]// nkhezno linput diel luser bch nkhedmo 3lih
	if len(args) != Size {// nchoufo linput diel luser wch fih 9 diel arguments
		fmt.Println("Error")
		return
	}

	board := make([][]rune, Size)//derna slice 2d diel runes bach ndiro fih linput wnkhedmo 3lih

	for i := 0; i < Size; i++ {//nchoufo koula argument wch fih 9 diel chars
		arg := args[i]
		if len(arg) != Size {
			fmt.Println("Error")
			return
		}
		board[i] = []rune(arg)
	}
	for i := 0; i < Size; i++ {//tchouf lina wch luser kteb chi 7aja mn ghir ar9am w "."
		for _, cell := range board[i] {
			if cell != Empty && (cell < '1' || cell > '9') {
				fmt.Println("Error")
				return
			}
		}
	}

	for i := 0; i < Size; i++ {//nchoufo linput diell luser wch 3tana ar9am m3awdin
		for j := 0; j < Size; j++ {
			if board[i][j] != Empty {
				// Check row
				for k := 0; k < Size; k++ {
					if k != j && board[i][k] == board[i][j] {
						fmt.Println("Error")
						return
					}
				}
				// Check column
				for k := 0; k < Size; k++ {
					if k != i && board[k][j] == board[i][j] {
						fmt.Println("Error")
						return
					}
				}
				// Check 3x3 box
				startRow, startCol := i-i%3, j-j%3
				for k := startRow; k < startRow+3; k++ {
					for l := startCol; l < startCol+3; l++ {
						if k != i && l != j && board[k][l] == board[i][j] {
							fmt.Println("Error")
							return
						}
					}
				}
			}
		}
	}

	if !solveSudoku(board){
		fmt.Println("Error")
		return
	}

	printBoard(board)
}

func solveSudoku(board [][]rune) bool {// li ghadi t7el lina sudoku
	var row, col int
	if !findEmptyCell(board, &row, &col){
		return true // Puzzle solved
	}

	for num := '1'; num <= '9'; num++ {
		if isValid(board, row, col, num) {
			board[row][col] = num
			if solveSudoku(board) {
				return true
			}
			board[row][col] = Empty // Backtrack
		}
	}

	return false // No solution found
}

func findEmptyCell(board [][]rune, row, col *int) bool {// tl9a lina ne9ta li kat3ni blassa khawya 
	for i := 0; i < Size; i++ {
		for j := 0; j < Size; j++ {
			if board[i][j] == Empty {
				*row, *col = i, j
				return true
			}
		}
	}
	return false
}

func isValid(board [][]rune, row, col int, num rune) bool {// tchouf lina wch dak ra9m li 3tinaha valid wlala
	return !usedInRow(board, row, num) &&
		!usedInColumn(board, col, num) &&
		!usedInBox(board, row-row%3, col-col%3, num)
}

func usedInRow(board [][]rune, row int, num rune) bool {// tchouf lina wach dak ra9m li derna kayn frow wlala
	for col := 0; col < Size; col++ {
		if board[row][col] == num {
			return true
		}
	}
	return false
}

func usedInColumn(board [][]rune, col int, num rune) bool {// tchouf lina wch dak ra9m li derna kayn flcol wlala
	for row := 0; row < Size; row++ {
		if board[row][col] == num {
			return true
		}
	}
	return false
}

func usedInBox(board [][]rune, startRow, startCol int, num rune) bool {//tchouf lina wch ra9m li derna kayna flbox wlala
	for row := 0; row < 3; row++ {
		for col := 0; col < 3; col++ {
			if board[row+startRow][col+startCol] == num {
				return true
			}
		}
	}
	return false
}

func printBoard(board [][]rune) {// tba3 lina sudoku mli ysali lkhedma
	for i := 0; i < Size; i++ {
		for j := 0; j < Size; j++ {
			fmt.Printf("%c ", board[i][j])
		}
		fmt.Println()
	}
}
