package main

import (
	"fmt"
	"os"
	"time"
)

var stuck int

type Guess struct {
	value int
	// possibilities []int
	i int
	j int
}

// Cell : each cell in a board
type Cell struct {
	val      int
	possible []int
}

// SudokuBoard : type for the board
type SudokuBoard struct {
	board      [][]*Cell
	solved     bool
	changed    bool
	madeGuess  bool
	guess      []Guess
	guessWrong bool
}

// NewSudokuBoard : initializes a new board
func NewSudokuBoard(vals [][]int) *SudokuBoard {
	b := new(SudokuBoard)

	for i := 0; i < 9; i++ {
		b.board = append(b.board, make([]*Cell, 9))
	}

	for i, val := range vals {
		for j, c := range val {
			if c == 0 {
				(*b).board[i][j] = &Cell{0, []int{}}
			} else {
				(*b).board[i][j] = &Cell{c, nil}
			}
		}
	}
	(*b).solved = false
	return b
}

// Print : print the board in a nice format
func (s SudokuBoard) Print() {
	fmt.Println()
	for i, val := range s.board {
		for j := range val {
			if j%3 == 0 && j != 0 {
				fmt.Printf("|")
			}
			fmt.Printf(" %v ", s.board[i][j].val)
		}
		if i == 2 || i == 5 {
			fmt.Printf("\n-----------------------------")
		}
		fmt.Println()
	}
}

// Solve : solves the board
func (s *SudokuBoard) Solve() {
	s.Print()
	for !s.solved {
		s.changed = false
		for i, row := range s.board {
			for j := range row {
				s.Evaluate(i, j)
				if s.guessWrong {
					break
				}
			}
			if s.guessWrong {
				break
			}
		}

		s.EvaluateRows()
		s.EvaluateColumns()
		s.EvaluateSquares()
		s.Check()

		if !s.changed && !s.solved {
			s.Guess()
			s.Print()
			s.IsValid()
		}

		// if !s.changed && !s.solved {
		// 	s.Print()
		// 	fmt.Println("Stuck")
		// 	for i, row := range s.board {
		// 		for j := range row {
		// 			if s.board[i][j].val == 0 {
		// 				// fmt.Println(i, j, s.board[i][j])
		// 			}
		// 		}
		// 	}
		// 	os.Exit(1)
		// 	stuck++
		// 	break
		// }

		if !s.IsValid() || s.guessWrong {
			fmt.Println("Is invalid, rolling back")
			fmt.Println(s.guess)
			fmt.Println(len(s.guess))
			os.Exit(1)
			// Roll back
		}
	}

}

// Guess : when there's no solid moves anymore, take a guess
func (s *SudokuBoard) Guess() {
	fmt.Println("Making a guess...")
	for i, row := range s.board {
		for j := range row {
			if s.board[i][j].val == 0 {
				if len(s.guess) == 0 {
					c := new(Cell)
					c.val = s.board[i][j].possible[0]
					s.board[i][j] = c
					s.guess = append(s.guess, Guess{c.val, i, j})
				} else if i != s.guess[0].i && j != s.guess[0].j {
					for _, p := range s.board[i][j].possible {
						if p != s.guess[0].value {
							c := new(Cell)
							c.val = p
							s.board[i][j] = c
							s.guess = append(s.guess, Guess{c.val, i, j})
						}
					}
				}
			}
		}
	}
}

// Check : check to see if board is solved
func (s *SudokuBoard) Check() {
	for i, row := range s.board {
		for j := range row {
			if s.board[i][j].val == 0 {
				s.solved = false
				s.IsValid()
				return
			}
		}
	}
	s.solved = true
	// s.Print()
	// fmt.Printf("FINISHED!\n")
}

// Evaluate : finds possibilities for a cell
func (s *SudokuBoard) Evaluate(i, j int) {
	if (*s).board[i][j].val == 0 {
		c := new(Cell)

		// Check rows
		row := s.CheckRow(i)
		col := s.CheckCol(j)
		cell := s.CheckSquare(i, j)

		possibilities := s.Reduce(row, col, cell)

		if len(possibilities) == 0 {
			s.guessWrong = true
		}

		if len(possibilities) == 1 {
			c.val = possibilities[0]
			c.possible = nil
		} else {
			c.val = 0
			c.possible = possibilities
		}

		// if len(possibilities) > 1 {
		// 	fmt.Println(c, i, j)
		// 	// fmt.Println(s.CheckCol(j))
		// }

		if c.val != 0 {
			fmt.Printf("Solved cell (%v, %v): %v\n", i, j, c.val)
			(*s).board[i][j] = c
			// s.Print()
			s.changed = true
		}
		(*s).board[i][j] = c
		s.IsValid()
	}
}

// Reduce :
func (s *SudokuBoard) Reduce(row, col, square []int) []int {
	var ret []int
	// First check if there is only one value in any of these
	count := 0
	for _, v := range row {
		if v != 0 {
			count++
			if count > 1 {
				break
			}
		}
	}
	if count == 1 {
		for _, v := range row {
			if v != 0 {
				return []int{v}
			}
		}
	}

	count = 0
	for _, v := range col {
		if v != 0 {
			count++
			if count > 1 {
				break
			}
		}
	}
	if count == 1 {
		for _, v := range col {
			if v != 0 {
				return []int{v}
			}
		}
	}

	count = 0
	for _, v := range square {
		if v != 0 {
			count++
			if count > 1 {
				break
			}
		}
	}
	if count == 1 {
		for _, v := range square {
			if v != 0 {
				return []int{v}
			}
		}
	}

	for i := 0; i < 9; i++ {
		if row[i] == col[i] && row[i] == square[i] && row[i] != 0 {
			ret = append(ret, row[i])
		}
	}

	return ret
}

// CheckRow :
func (s *SudokuBoard) CheckRow(i int) []int {
	// vals := []int{}
	n := []int{1, 2, 3, 4, 5, 6, 7, 8, 9} //0, 0, 0, 0, 0, 0, 0, 0, 0}
	for j := range s.board[i] {
		if v := s.board[i][j].val; v != 0 {
			n[v-1] = 0
		}
	}
	return n
}

// CheckCol :
func (s *SudokuBoard) CheckCol(j int) []int {
	n := []int{1, 2, 3, 4, 5, 6, 7, 8, 9} //{0, 0, 0, 0, 0, 0, 0, 0, 0}
	for i := range s.board[j] {
		if v := s.board[i][j].val; v != 0 {
			n[v-1] = 0
		}
	}
	return n
}

// CheckSquare :
func (s *SudokuBoard) CheckSquare(i, j int) []int {
	// var n int
	if i < 3 {
		if j < 3 {
			i, j = 0, 0
		} else if j < 6 {
			i, j = 0, 3
		} else {
			i, j = 0, 6
		}
	} else if i < 6 {
		if j < 3 {
			i, j = 3, 0
		} else if j < 6 {
			i, j = 3, 3
		} else {
			i, j = 3, 6
		}
	} else {
		if j < 3 {
			i, j = 6, 0
		} else if j < 6 {
			i, j = 6, 3
		} else {
			i, j = 6, 6
		}
	}

	n := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	for k := 0; k < 3; k++ {
		for l := 0; l < 3; l++ {
			if v := s.board[i+k][j+l].val; v != 0 {
				n[v-1] = 0
			}
		}
	}
	return n
}

// EvaluateRows :
func (s *SudokuBoard) EvaluateRows() {
	for i, row := range s.board {
		n := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
		for _, cell := range row {
			if cell.val != 0 {
				n[cell.val-1] = 0
			}
		}

		double := make([]bool, 9)
		for i := 0; i < 9; i++ {
			double[i] = false
		}
		for _, cell := range row {
			if cell.val == 0 {
				for _, p := range cell.possible {
					if double[p-1] {
						n[p-1] = 0
					}
					double[p-1] = true
				}
			}
		}

		count := 0
		val := 0
		for i := range n {
			if n[i] != 0 {
				count++
				val = n[i]
			}
		}

		if count == 1 {
			for j, cell := range row {
				for _, p := range cell.possible {
					if p == val {
						c := new(Cell)
						c.val = val
						(*s).board[i][j] = c
						// s.Print()
						// os.Exit(1)
					}
				}
			}
		}

	}

	// os.Exit(1)
}

// EvaluateColumns :
func (s *SudokuBoard) EvaluateColumns() {
	for j := range s.board {
		n := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
		for i := range s.board {
			if v := s.board[i][j].val; v != 0 {
				n[v-1] = 0
			}
		}

		double := make([]bool, 9)
		for i := 0; i < 9; i++ {
			double[i] = false
		}
		for i := range s.board {
			if s.board[i][j].val == 0 {
				for _, p := range s.board[i][j].possible {
					if double[p-1] {
						n[p-1] = 0
					}
					double[p-1] = true
				}
			}
		}

		count := 0
		val := 0
		for i := range n {
			if n[i] != 0 {
				count++
				val = n[i]
			}
		}

		if count == 1 {
			for i := range s.board {
				for _, p := range s.board[i][j].possible {
					if p == val {
						c := new(Cell)
						c.val = val
						(*s).board[i][j] = c
						// s.Print()
						// os.Exit(1)
					}
				}
			}
		}
	}
}

// EvaluateCells :
func (s *SudokuBoard) EvaluateSquares() {
	for a := 0; a < 9; a++ {
		i := 3 * (a / 3)
		j := (a * 3) % 9

		n := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

		for k := 0; k < 3; k++ {
			for l := 0; l < 3; l++ {
				if v := s.board[i+k][j+l].val; v != 0 {
					n[v-1] = 0
				}
			}
		}

		double := make([]bool, 9)
		for i := 0; i < 9; i++ {
			double[i] = false
		}

		for k := 0; k < 3; k++ {
			for l := 0; l < 3; l++ {
				for _, p := range s.board[i+k][j+l].possible {
					if double[p-1] {
						n[p-1] = 0
					}
					double[p-1] = true
				}
			}
		}

		count := 0
		val := 0
		for i := range n {
			if n[i] != 0 {
				count++
				val = n[i]
			}
		}

		for k := 0; k < 3; k++ {
			for l := 0; l < 3; l++ {
				for _, p := range s.board[i+k][j+l].possible {
					if p == val {
						c := new(Cell)
						c.val = val
						(*s).board[i+k][j+l] = c
					}
				}
			}
		}

	}

}

// IsValid : checks if solution is valid
func (s *SudokuBoard) IsValid() bool {
	m := make(map[int]bool)
	for i, row := range s.board {
		for j := range row {
			v := s.board[i][j].val
			if v != 0 && m[v] {
				s.guessWrong = true
				return false
			}
			m[v] = true
		}
	}

	for j, row := range s.board {
		col := make(map[int]bool)
		for i := range row {
			v := s.board[i][j].val
			if v != 0 && col[v] {
				s.guessWrong = true
				return false
			}
			col[v] = true
		}
	}

	for a := 0; a < 9; a++ {
		i := 3 * (a / 3)
		j := (a * 3) % 9
		sq := make(map[int]bool)
		for k := 0; k < 3; k++ {
			for l := 0; l < 3; l++ {
				v := s.board[i+k][j+l].val
				if v != 0 && sq[v] {
					s.guessWrong = true
					return false
				}
				sq[v] = true
			}
		}
	}

	s.guessWrong = false
	return true
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

// TopThree : get the value of top left corner
func (s *SudokuBoard) TopThree() int {
	return s.board[0][0].val*100 + s.board[0][1].val*10 + s.board[0][2].val
}

func main() {
	// start := time.Now()
	// var sum int
	// file := "./p096_sudoku.txt"

	// csvFile, err := os.Open(file)
	// check(err)

	// reader := csv.NewReader(bufio.NewReader(csvFile))

	// var rows [][]int
	// for i := 0; i < 9; i++ {
	// 	rows = append(rows, make([]int, 9))
	// }
	// count := 0
	// for {
	// 	line, err := reader.Read()
	// 	if err == io.EOF {
	// 		break
	// 	} else {
	// 		fmt.Println(line)
	// 		row := []int{}
	// 		newRow := true
	// 		for i := 0; i < 9; i++ {
	// 			n, err := strconv.Atoi(string(line[0][i]))
	// 			if err != nil {
	// 				newRow = false
	// 				break
	// 			} else {
	// 				row = append(row, n)
	// 				// fmt.Println(row, count)
	// 			}
	// 		}
	// 		if newRow {
	// 			// fmt.Println(count)
	// 			rows[count] = row
	// 			count++
	// 			// fmt.Println(rows)
	// 			if count == 9 {
	// 				// fmt.Println(rows)
	// 				b := NewSudokuBoard(rows)
	// 				// b.Print()
	// 				b.Solve()
	// 				sum += b.TopThree()
	// 				fmt.Println(sum)
	// 				count = 0
	// 				// os.Exit(1)
	// 			}
	// 		}
	// 	}
	// }

	var rows [][]int
	row1 := []int{1, 0, 0, 9, 2, 0, 0, 0, 0}
	rows = append(rows, row1)
	row2 := []int{5, 2, 4, 0, 1, 0, 0, 0, 0}
	rows = append(rows, row2)
	row3 := []int{0, 0, 0, 0, 0, 0, 0, 7, 0}
	rows = append(rows, row3)
	row4 := []int{0, 5, 0, 0, 0, 8, 1, 0, 2}
	rows = append(rows, row4)
	row5 := []int{0, 0, 0, 0, 0, 0, 0, 0, 0}
	rows = append(rows, row5)
	row6 := []int{4, 0, 2, 7, 0, 0, 0, 9, 0}
	rows = append(rows, row6)
	row7 := []int{0, 6, 0, 0, 0, 0, 0, 0, 0}
	rows = append(rows, row7)
	row8 := []int{0, 0, 0, 0, 3, 0, 9, 4, 5}
	rows = append(rows, row8)
	row9 := []int{0, 0, 0, 0, 7, 1, 0, 0, 6}
	rows = append(rows, row9)

	b := NewSudokuBoard(rows)

	b.Print()

	start := time.Now()

	b.Solve()
	fmt.Println(stuck)
	fmt.Printf("Solved in %s\n", time.Since(start))
	// // b.Print()
}
