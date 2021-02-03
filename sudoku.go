package sudoku

import (
	"fmt"
	"strings"
)

const (
	// Size of the sudoku. The size also affects the range of possible digits.
	// Maximum size is 30 to prevent int32 overflow. The chosen size must also
	// be a perfect square root, due to a limitation in the implementation.
	// This means only the sizes 4, 9, 16 and 25 are supported.
	Size = 9

	// BoxSize is the square root of Size.
	BoxSize = 3
)

// Sudoku contains the data structure needed to solve the puzzle.
type Sudoku struct {
	// Grid contains the numbers as provided by the user. This grid
	// is manipulated directly when solving the sudoku.
	Grid [Size][Size]int

	// Box, Row and Column digits are stored in an int for fast
	// changes and checks for a certain digit with bitwise operators.
	boxDigits    [BoxSize][BoxSize]int
	rowDigits    [Size]int
	columnDigits [Size]int

	// Amount of steps for solving the sudoku.
	StepCount uint64
}

// New creates a new sudoku.
func New(grid [Size][Size]int) *Sudoku {
	return &Sudoku{Grid: grid}
}

// backTrack by assigning consecutive digits to empty cells and check
// if that number is allowed. This is done recursively until the sudoku
// is solved or until the sudoku is deemed unsolvable.
func (s *Sudoku) backTrack(row, column int) bool {
	if row == Size {
		return true
	}
	if column == Size {
		return s.backTrack(row+1, 0)
	}

	// Start trying numbers when the current cell is empty.
	if s.Grid[row][column] == 0 {
		s.StepCount++

		boxCell := s.boxDigits[row/BoxSize][column/BoxSize]
		rowCell := s.rowDigits[row]
		columnCell := s.columnDigits[column]

		// Digits that are already set in either a box, row or column
		// and can't be placed in the current cell.
		setDigits := boxCell | rowCell | columnCell

		// Loop through all the digits.
		for i := 1; i <= Size; i++ {
			bit := 1 << (i - 1)

			// Check if the digit can be placed in the current cell.
			if setDigits&bit == 0 {
				// Set digit.
				s.boxDigits[row/BoxSize][column/BoxSize] |= bit
				s.rowDigits[row] |= bit
				s.columnDigits[column] |= bit
				s.Grid[row][column] = i

				// Solve the sudoku further. When this call returns
				// false the digit is incorrect.
				if s.backTrack(row, column+1) {
					return true
				}

				// Digit was incorrect, remove.
				s.boxDigits[row/BoxSize][column/BoxSize] &^= bit
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
	for i := 0; i < Size; i++ {
		for j := 0; j < Size; j++ {
			if digit := s.Grid[i][j]; digit > 0 {
				bit := 1 << (digit - 1)

				s.boxDigits[i/BoxSize][j/BoxSize] |= bit
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
