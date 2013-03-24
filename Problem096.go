package main

import (
	"./euler"
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

type puzzle struct {
	grid      [][]int
	inference [][][]bool
}

func (puz *puzzle) print() {
	for i, row := range puz.grid {
		if i%3 == 0 && i > 0 {
			fmt.Print("-----------\n")
		}
		for j, value := range row {
			if j%3 == 0 && j > 0 {
				fmt.Print("|")
			}

			if value != 0 {
				fmt.Print(value)

			} else {
				fmt.Print(" ")
			}
		}
		fmt.Print("\n")
	}
	fmt.Print("\n")

}

//set up initial inference grid
func (puz *puzzle) initInfer() {
	puz.inference = make([][][]bool, 9)
	for i := range puz.inference {
		puz.inference[i] = make([][]bool, 9)
		for j := range puz.inference[i] {
			puz.inference[i][j] = make([]bool, 9)
			for k := range puz.inference[i][j] {
				if puz.grid[i][j] == 0 || puz.grid[i][j]-1 == k {
					puz.inference[i][j][k] = true
				} else {
					puz.inference[i][j][k] = false

				}

			}
		}
	}

}

//add information from columns / rows to inference table
func (puz *puzzle) inferStraight() {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			value := puz.grid[i][j]
			if value != 0 {
				for k := 0; k < 9; k++ {
					if puz.inference[i][k][value-1] && k != j {
						puz.inference[i][k][value-1] = false
					}

					if puz.inference[k][j][value-1] && k != i {
						puz.inference[k][j][value-1] = false
					}

				}

			}
		}
	}

}

//fill in squares based on what we've infered
func (puz *puzzle) deduce() {
	//check for only one possibility in a grid spot
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {

			if puz.grid[i][j] == 0 {

				answer := 0
				total := 0

				for k := 0; k < 9; k++ {
					if puz.inference[i][j][k] {
						total++
						answer = k
					}
				}

				if total == 1 {
					puz.grid[i][j] = answer + 1
				}

			}

		}
	}

	//check for only one possibility in a row / column

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {

			total1 := 0
			total2 := 0
			answer1 := 0
			answer2 := 0

			for k := 0; k < 9; k++ {
				if puz.inference[i][k][j] {
					total1++
					answer1 = k
				}
				if puz.inference[k][i][j] {
					total2++
					answer2 = k
				}
			}

			if total1 == 1 {
				puz.grid[i][answer1] = j + 1
			}

			if total2 == 1 {
				puz.grid[answer2][i] = j + 1
			}

		}

	}

	//check for only one possibility in a 3x3 box
	for istart := 0; istart < 9; istart += 3 {
		for jstart := 0; jstart < 9; jstart += 3 {

			for k := 0; k < 9; k++ {
				total := 0
				iplace, jplace := 0, 0

				for i := 0; i < 3; i++ {
					for j := 0; j < 3; j++ {

						if puz.inference[i+istart][j+jstart][k] {
							total++
							iplace, jplace = i+istart, j+jstart
						}

					}
				}

				if total == 1 {
					puz.grid[iplace][jplace] = k + 1
				}
			}
		}
	}

}

func (puz *puzzle) inferSquares() {
	for istart := 0; istart < 9; istart += 3 {
		for jstart := 0; jstart < 9; jstart += 3 {

			for i := 0; i < 3; i++ {
				for j := 0; j < 3; j++ {

					value := puz.grid[i+istart][j+jstart]
					if value != 0 {

						for k := 0; k < 3; k++ {
							for l := 0; l < 3; l++ {
								if k != i || l != j {

									puz.inference[k+istart][l+jstart][value-1] = false

								}

							}
						}

					}

				}
			}
		}
	}
}

func (puz *puzzle) isComplete() bool {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if puz.grid[i][j] == 0 {
				return false
			}
		}
	}
	return true
}

func (puz *puzzle) isBroken() bool {
	rows := make([]bool, 10)
	columns := make([]bool, 10)

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if puz.grid[i][j] != 0 && rows[puz.grid[i][j]] {
				return true

			}
			if puz.grid[j][i] != 0 && columns[puz.grid[j][i]] {
				return true
			}

			rows[puz.grid[i][j]] = true
			columns[puz.grid[j][i]] = true
		}
		rows = make([]bool, 10)
		columns = make([]bool, 10)

	}

	//inconcsistent 3x3 squares
	for istart := 0; istart < 9; istart += 3 {
		for jstart := 0; jstart < 9; jstart += 3 {

			for i := 0; i < 3; i++ {
				for j := 0; j < 3; j++ {

					if puz.grid[i+istart][j+jstart] != 0 && rows[puz.grid[i+istart][j+jstart]] {
						return true
					}

					rows[puz.grid[i+istart][j+jstart]] = true
				}
			}

			rows = make([]bool, 10)

		}
	}

	//check for no possibility in a grid spot
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {

			total := 0

			for k := 0; k < 9; k++ {
				if puz.inference[i][j][k] {
					total++
				}
			}

			if total == 0 {
				return true
			}

		}
	}

	return false

}

func solve(puz *puzzle) bool {
	puz.initInfer()
	stuck := 0
	for !puz.isComplete() && !puz.isBroken() && stuck < 1000 {
		puz.inferStraight()
		puz.inferSquares()

		puz.deduce()
		stuck++
	}

	if puz.isBroken() {
		return false
	}

	if puz.isComplete() {
		return true
	}

	//if we have neither completed nor broken the puzzle, we guess
	clone := puzzle{grid: puz.grid}
	exhaustion := 0
	for !clone.isComplete()) && exhaustion < 200 {

		//find a guess to make
		a := rand.Int() % 9
		b := rand.Int() % 9
		c := rand.Int() % 9
		for clone.grid[a][b] != 0 || !puz.inference[a][b][c] {
			a = rand.Int() % 9
			b = rand.Int() % 9
			c = rand.Int() % 9
		}

		clone.grid[a][b] = c + 1
		fmt.Println("guessing:")
		clone.print()

		if !solve(&clone) || clone.isBroken() {
			clone = puzzle{grid: puz.grid}
			clone.initInfer()
		}

		exhaustion++
	}
	if exhaustion == 200 {
		return false
		fmt.Println("I'm exhausted")
	}

	puz.grid = clone.grid

	return true

}

func main() {
	starttime := time.Now()

	data := euler.Import("problemdata/sudoku.txt")

	for offset := 1; offset < 500; offset += 10 {
		grid := make([][]int, 9)
		for i := 0; i < 9; i++ {
			grid[i] = make([]int, 9)

			for j := 0; j < 9; j++ {
				grid[i][j], _ = strconv.Atoi(data[i+offset][j : j+1])
			}

		}

		puz := puzzle{grid: grid}

		fmt.Println("solving...")
		puz.print()
		if !solve(&puz) {
			panic("BROKEN")
		}

		fmt.Println("Puzzle", ((offset-1)/10)+1)
		puz.print()
		fmt.Println(puz.isBroken())

	}
	fmt.Println("Elapsed time:", time.Since(starttime))

}
