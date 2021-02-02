# Sudoku Solver

## Usage

```go
package main

import (
    "fmt"

    "github.com/michelheusschen/sudoku"
)

func main() {
    // Use the New function to create a new unsolved sudoku.
    s := sudoku.New([9][9]int{
        {8, 0, 0, 0, 0, 0, 0, 0, 0},
        {0, 0, 3, 6, 0, 0, 0, 0, 0},
        {0, 7, 0, 0, 9, 0, 2, 0, 0},
        {0, 5, 0, 0, 0, 7, 0, 0, 0},
        {0, 0, 0, 0, 4, 5, 7, 0, 0},
        {0, 0, 0, 1, 0, 0, 0, 3, 0},
        {0, 0, 1, 0, 0, 0, 0, 6, 8},
        {0, 0, 8, 5, 0, 0, 0, 1, 0},
        {0, 9, 0, 0, 0, 0, 4, 0, 0},
    })

    // Solve the sudoku. Returns a boolean to indicate success.
    success := s.Solve()
    if success {
        fmt.Printf("Sudoku was solved in %d steps\n", s.StepCount)
    } else {
        fmt.Println("Sudoku couldn't be solved")
    }

    // Access the grid of the solved sudoku.
    _ = s.Grid

    // Print the result.
    fmt.Println(s)

    // [8, 1, 2, 7, 5, 3, 6, 4, 9]
    // [9, 4, 3, 6, 8, 2, 1, 7, 5]
    // [6, 7, 5, 4, 9, 1, 2, 8, 3]
    // [1, 5, 4, 2, 3, 7, 8, 9, 6]
    // [3, 6, 9, 8, 4, 5, 7, 2, 1]
    // [2, 8, 7, 1, 6, 9, 5, 3, 4]
    // [5, 2, 1, 9, 7, 4, 3, 6, 8]
    // [4, 3, 8, 5, 2, 6, 9, 1, 7]
    // [7, 9, 6, 3, 1, 8, 4, 5, 2]
}
```