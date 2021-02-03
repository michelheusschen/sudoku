package sudoku

import (
	"testing"
)

type testCase struct {
	Name           string
	Solvable       bool
	Unsolved       [Size][Size]int
	UnsolvedString string
	Solved         [Size][Size]int
	SolvedString   string
}

func TestSolve(t *testing.T) {
	testCases := []testCase{
		{
			Name:     "Easy 9x9 sudoku",
			Solvable: true,
			Unsolved: [Size][Size]int{
				{5, 3, 0, 0, 7, 0, 0, 0, 0},
				{6, 0, 0, 1, 9, 5, 0, 0, 0},
				{0, 9, 8, 0, 0, 0, 0, 6, 0},
				{8, 0, 0, 0, 6, 0, 0, 0, 3},
				{4, 0, 0, 8, 0, 3, 0, 0, 1},
				{7, 0, 0, 0, 2, 0, 0, 0, 6},
				{0, 6, 0, 0, 0, 0, 2, 8, 0},
				{0, 0, 0, 4, 1, 9, 0, 0, 5},
				{0, 0, 0, 0, 8, 0, 0, 7, 9},
			},
			UnsolvedString: `[5, 3, 0, 0, 7, 0, 0, 0, 0]
[6, 0, 0, 1, 9, 5, 0, 0, 0]
[0, 9, 8, 0, 0, 0, 0, 6, 0]
[8, 0, 0, 0, 6, 0, 0, 0, 3]
[4, 0, 0, 8, 0, 3, 0, 0, 1]
[7, 0, 0, 0, 2, 0, 0, 0, 6]
[0, 6, 0, 0, 0, 0, 2, 8, 0]
[0, 0, 0, 4, 1, 9, 0, 0, 5]
[0, 0, 0, 0, 8, 0, 0, 7, 9]
`,
			Solved: [Size][Size]int{
				{5, 3, 4, 6, 7, 8, 9, 1, 2},
				{6, 7, 2, 1, 9, 5, 3, 4, 8},
				{1, 9, 8, 3, 4, 2, 5, 6, 7},
				{8, 5, 9, 7, 6, 1, 4, 2, 3},
				{4, 2, 6, 8, 5, 3, 7, 9, 1},
				{7, 1, 3, 9, 2, 4, 8, 5, 6},
				{9, 6, 1, 5, 3, 7, 2, 8, 4},
				{2, 8, 7, 4, 1, 9, 6, 3, 5},
				{3, 4, 5, 2, 8, 6, 1, 7, 9},
			},
			SolvedString: `[5, 3, 4, 6, 7, 8, 9, 1, 2]
[6, 7, 2, 1, 9, 5, 3, 4, 8]
[1, 9, 8, 3, 4, 2, 5, 6, 7]
[8, 5, 9, 7, 6, 1, 4, 2, 3]
[4, 2, 6, 8, 5, 3, 7, 9, 1]
[7, 1, 3, 9, 2, 4, 8, 5, 6]
[9, 6, 1, 5, 3, 7, 2, 8, 4]
[2, 8, 7, 4, 1, 9, 6, 3, 5]
[3, 4, 5, 2, 8, 6, 1, 7, 9]
`,
		},
		{
			Name:     "Very hard 9x9 sudoku",
			Solvable: true,
			Unsolved: [Size][Size]int{
				{8, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 3, 6, 0, 0, 0, 0, 0},
				{0, 7, 0, 0, 9, 0, 2, 0, 0},
				{0, 5, 0, 0, 0, 7, 0, 0, 0},
				{0, 0, 0, 0, 4, 5, 7, 0, 0},
				{0, 0, 0, 1, 0, 0, 0, 3, 0},
				{0, 0, 1, 0, 0, 0, 0, 6, 8},
				{0, 0, 8, 5, 0, 0, 0, 1, 0},
				{0, 9, 0, 0, 0, 0, 4, 0, 0},
			},
			UnsolvedString: `[8, 0, 0, 0, 0, 0, 0, 0, 0]
[0, 0, 3, 6, 0, 0, 0, 0, 0]
[0, 7, 0, 0, 9, 0, 2, 0, 0]
[0, 5, 0, 0, 0, 7, 0, 0, 0]
[0, 0, 0, 0, 4, 5, 7, 0, 0]
[0, 0, 0, 1, 0, 0, 0, 3, 0]
[0, 0, 1, 0, 0, 0, 0, 6, 8]
[0, 0, 8, 5, 0, 0, 0, 1, 0]
[0, 9, 0, 0, 0, 0, 4, 0, 0]
`,
			Solved: [Size][Size]int{
				{8, 1, 2, 7, 5, 3, 6, 4, 9},
				{9, 4, 3, 6, 8, 2, 1, 7, 5},
				{6, 7, 5, 4, 9, 1, 2, 8, 3},
				{1, 5, 4, 2, 3, 7, 8, 9, 6},
				{3, 6, 9, 8, 4, 5, 7, 2, 1},
				{2, 8, 7, 1, 6, 9, 5, 3, 4},
				{5, 2, 1, 9, 7, 4, 3, 6, 8},
				{4, 3, 8, 5, 2, 6, 9, 1, 7},
				{7, 9, 6, 3, 1, 8, 4, 5, 2},
			},
			SolvedString: `[8, 1, 2, 7, 5, 3, 6, 4, 9]
[9, 4, 3, 6, 8, 2, 1, 7, 5]
[6, 7, 5, 4, 9, 1, 2, 8, 3]
[1, 5, 4, 2, 3, 7, 8, 9, 6]
[3, 6, 9, 8, 4, 5, 7, 2, 1]
[2, 8, 7, 1, 6, 9, 5, 3, 4]
[5, 2, 1, 9, 7, 4, 3, 6, 8]
[4, 3, 8, 5, 2, 6, 9, 1, 7]
[7, 9, 6, 3, 1, 8, 4, 5, 2]
`,
		},
	}

	for _, test := range testCases {
		s := New(test.Unsolved)

		if unsolvedString := s.String(); unsolvedString != test.UnsolvedString {
			t.Errorf("Test %s failed, unsolved string representation doesn't match", test.Name)
			t.Errorf("Expected string:\n%s", test.UnsolvedString)
			t.Errorf("But got string:\n%s", unsolvedString)
		}

		solved := s.Solve()

		if solved != test.Solvable {
			t.Errorf("Test %s failed, expected %t for Solve() but got %t", test.Name, test.Solvable, solved)
		}
		if solved && !gridIsEqual(s.Grid, test.Solved) {
			t.Errorf("Test %s failed, grid doesn't match", test.Name)
			t.Errorf("Expected grid:\n%s", New(test.Solved))
			t.Errorf("But got grid:\n%s", s)

			if solvedString := s.String(); solvedString != test.SolvedString {
				t.Errorf("Test %s failed, solved string representation doesn't match", test.Name)
				t.Errorf("Expected string:\n%s", test.SolvedString)
				t.Errorf("But got string:\n%s", solvedString)
			}
		}
	}
}

func gridIsEqual(a, b [Size][Size]int) bool {
	for i := 0; i < Size; i++ {
		for j := 0; j < Size; j++ {
			if a[i][j] != b[i][j] {
				return false
			}
		}
	}

	return true
}

func BenchmarkSolveEasy(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := New([Size][Size]int{
			{5, 3, 0, 0, 7, 0, 0, 0, 0},
			{6, 0, 0, 1, 9, 5, 0, 0, 0},
			{0, 9, 8, 0, 0, 0, 0, 6, 0},
			{8, 0, 0, 0, 6, 0, 0, 0, 3},
			{4, 0, 0, 8, 0, 3, 0, 0, 1},
			{7, 0, 0, 0, 2, 0, 0, 0, 6},
			{0, 6, 0, 0, 0, 0, 2, 8, 0},
			{0, 0, 0, 4, 1, 9, 0, 0, 5},
			{0, 0, 0, 0, 8, 0, 0, 7, 9},
		})

		_ = s.Solve()
	}
}

func BenchmarkSolveHard(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := New([Size][Size]int{
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

		_ = s.Solve()
	}
}
