package sudoku

import (
	"fmt"
	"strings"
)

// Sudoku contains the data structure needed to solve the puzzle.
type Sudoku struct {
	// Grid contains the numbers as provided by the user. This grid
	// is manipulated directly when solving the sudoku.
	Grid [9][9]int

	// Box, Row and Column digits are stored in an int for fast
	// changes and checks for a certain digit with bitwise operators.
	boxDigits    [3][3]int
	rowDigits    [9]int
	columnDigits [9]int

	// Amount of steps for solving the sudoku.
	StepCount uint64
}

// New creates a new sudoku.
func New(grid [9][9]int) *Sudoku {
	return &Sudoku{Grid: grid}
}

// backTrack by assigning consecutive digits to empty cells and check
// if that number is allowed. This is done recursively until the sudoku
// is solved or until the sudoku is deemed unsolvable.
func (s *Sudoku) backTrack(row, column int) bool {
	if row == 9 {
		return true
	}
	if column == 9 {
		return s.backTrack(row+1, 0)
	}

	// Start trying numbers when the current cell is empty.
	if s.Grid[row][column] == 0 {
		s.StepCount++

		// Loop through the possible digits.
		for i := 1; i <= 9; i++ {
			bit := 1 << (i - 1)

			// Check with bitwise operators if the digit is still needed
			// in a part of the sudoku.
			boxNeedsDigit := s.boxDigits[row/3][column/3]&bit == 0
			rowNeedsDigit := s.rowDigits[row]&bit == 0
			columnNeedsDigit := s.columnDigits[column]&bit == 0

			// Only set digit when box, row and column needs the digit.
			if boxNeedsDigit && rowNeedsDigit && columnNeedsDigit {
				// Set digit.
				s.boxDigits[row/3][column/3] |= bit
				s.rowDigits[row] |= bit
				s.columnDigits[column] |= bit
				s.Grid[row][column] = i

				// Solve the sudoku further. When this call returns
				// false the digit is incorrect.
				if s.backTrack(row, column+1) {
					return true
				}

				// Digit was incorrect, remove.
				s.boxDigits[row/3][column/3] &^= bit
				s.rowDigits[row] &^= bit
				s.columnDigits[column] &^= bit
				s.Grid[row][column] = 0
			}
		}

		// No possible digit found, so sudoku is impossible to solve.
		return false
	}

	return s.backTrack(row, column+1)
}

// Solve the sudoku. The grid is manipulated directly and
// contains the end result.
func (s *Sudoku) Solve() bool {
	// Initialize box, row and column bits.
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if digit := s.Grid[i][j]; digit > 0 {
				bit := 1 << (digit - 1)

				s.boxDigits[i/3][j/3] |= bit
				s.rowDigits[i] |= bit
				s.columnDigits[j] |= bit
			}
		}
	}

	return s.backTrack(0, 0)
}

// String converts the grid to a string.
func (s *Sudoku) String() (output string) {
	for _, row := range s.Grid {
		output += strings.ReplaceAll(fmt.Sprint(row), " ", ", ") + "\n"
	}

	return output
}
