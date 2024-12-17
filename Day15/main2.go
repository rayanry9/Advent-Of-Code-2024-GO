package main

import (
	"bufio"
	"fmt"
	"os"
	//"time"
)

const SIZE = 50

type MoveType int

const (
	UP MoveType = iota
	DOWN
	LEFT
	RIGHT
)

func moveRe(grid *[SIZE][SIZE * 2]int, to [2]int, move MoveType) bool {

	switch move {
	case UP:

		if (*grid)[to[0]][to[1]] == 0 {
			if (*grid)[to[0]+1][to[1]] == 3 {
				(*grid)[to[0]][to[1]] = 3
				(*grid)[to[0]+1][to[1]] = 0
				return true
			} else {
				if grid[to[0]][to[1]+1] == 0 {
					(*grid)[to[0]][to[1]] = 1
					(*grid)[to[0]][to[1]+1] = 2
					(*grid)[to[0]+1][to[1]] = 0
					(*grid)[to[0]+1][to[1]+1] = 0
					return true
				} else if grid[to[0]][to[1]+1] == 1 {
					if moveRe(grid, [2]int{to[0] - 1, to[1] + 1}, move) {
						(*grid)[to[0]][to[1]] = 1
						(*grid)[to[0]][to[1]+1] = 2
						(*grid)[to[0]+1][to[1]] = 0
						(*grid)[to[0]+1][to[1]+1] = 0
						return true
					} else {
						return false
					}
				} else {
					return false
				}
			}
		} else if grid[to[0]][to[1]] == 1 {
			if moveRe(grid, [2]int{to[0] - 1, to[1]}, move) {
				if grid[to[0]+1][to[1]] == 3 {
					(*grid)[to[0]][to[1]] = 3
					(*grid)[to[0]+1][to[1]] = 0
				} else {
					(*grid)[to[0]][to[1]] = 1
					(*grid)[to[0]][to[1]+1] = 2
					(*grid)[to[0]+1][to[1]] = 0
					(*grid)[to[0]+1][to[1]+1] = 0
				}
				return true
			} else {
				return false
			}
		} else if grid[to[0]][to[1]] == 2 {
			if grid[to[0]+1][to[1]] == 3 {
				if moveRe(grid, [2]int{to[0] - 1, to[1] - 1}, move) {
					grid[to[0]][to[1]] = 3
					grid[to[0]+1][to[1]] = 0
					return true
				} else {
					return false
				}
			} else {
				if grid[to[0]][to[1]+1] == 0 {
					if moveRe(grid, [2]int{to[0] - 1, to[1] - 1}, move) {
						grid[to[0]][to[1]] = 1
						grid[to[0]][to[1]+1] = 2
						grid[to[0]+1][to[1]] = 0
						grid[to[0]+1][to[1]+1] = 0
						return true
					}
				} else if grid[to[0]][to[1]+1] == 1 {
					tmp := *grid
					if moveRe(&tmp, [2]int{to[0] - 1, to[1] - 1}, move) && moveRe(&tmp, [2]int{to[0] - 1, to[1] + 1}, move) {

						moveRe(grid, [2]int{to[0] - 1, to[1] - 1}, move)
						moveRe(grid, [2]int{to[0] - 1, to[1] + 1}, move)

						(*grid)[to[0]][to[1]] = 1
						(*grid)[to[0]][to[1]+1] = 2
						(*grid)[to[0]+1][to[1]] = 0
						(*grid)[to[0]+1][to[1]+1] = 0
						return true
					} else {
						return false
					}
				} else {
					return false
				}
			}
		} else if grid[to[0]][to[1]] == -1 {
			return false
		}

	case DOWN:

		if (*grid)[to[0]][to[1]] == 0 {
			if (*grid)[to[0]-1][to[1]] == 3 {
				(*grid)[to[0]][to[1]] = 3
				(*grid)[to[0]-1][to[1]] = 0
				return true
			} else {
				if grid[to[0]][to[1]+1] == 0 {
					(*grid)[to[0]][to[1]] = 1
					(*grid)[to[0]][to[1]+1] = 2
					(*grid)[to[0]-1][to[1]] = 0
					(*grid)[to[0]-1][to[1]+1] = 0
					return true
				} else if grid[to[0]][to[1]+1] == 1 {
					if moveRe(grid, [2]int{to[0] + 1, to[1] + 1}, move) {
						(*grid)[to[0]][to[1]] = 1
						(*grid)[to[0]][to[1]+1] = 2
						(*grid)[to[0]-1][to[1]] = 0
						(*grid)[to[0]-1][to[1]+1] = 0
						return true
					} else {
						return false
					}
				} else {
					return false
				}
			}
		} else if grid[to[0]][to[1]] == 1 {
			if moveRe(grid, [2]int{to[0] + 1, to[1]}, move) {
				if grid[to[0]-1][to[1]] == 3 {
					(*grid)[to[0]][to[1]] = 3
					(*grid)[to[0]-1][to[1]] = 0
				} else {
					(*grid)[to[0]][to[1]] = 1
					(*grid)[to[0]][to[1]+1] = 2
					(*grid)[to[0]-1][to[1]] = 0
					(*grid)[to[0]-1][to[1]+1] = 0
				}
				return true
			} else {
				return false
			}
		} else if grid[to[0]][to[1]] == 2 {
			if grid[to[0]-1][to[1]] == 3 {
				if moveRe(grid, [2]int{to[0] + 1, to[1] - 1}, move) {
					grid[to[0]][to[1]] = 3
					grid[to[0]-1][to[1]] = 0
					return true
				} else {
					return false
				}
			} else {
				if grid[to[0]][to[1]+1] == 0 {
					if moveRe(grid, [2]int{to[0] + 1, to[1] - 1}, move) {
						grid[to[0]][to[1]] = 1
						grid[to[0]][to[1]+1] = 2
						grid[to[0]-1][to[1]] = 0
						grid[to[0]-1][to[1]+1] = 0
						return true
					}
				} else if grid[to[0]][to[1]+1] == 1 {
					tmp := *grid
					if moveRe(&tmp, [2]int{to[0] + 1, to[1] - 1}, move) && moveRe(&tmp, [2]int{to[0] + 1, to[1] + 1}, move) {

						moveRe(grid, [2]int{to[0] + 1, to[1] - 1}, move)
						moveRe(grid, [2]int{to[0] + 1, to[1] + 1}, move)

						(*grid)[to[0]][to[1]] = 1
						(*grid)[to[0]][to[1]+1] = 2
						(*grid)[to[0]-1][to[1]] = 0
						(*grid)[to[0]-1][to[1]+1] = 0
						return true
					} else {
						return false
					}
				} else {
					return false
				}
			}
		} else if grid[to[0]][to[1]] == -1 {
			return false
		}

	case LEFT:

		if (*grid)[to[0]][to[1]] == 0 {
			(*grid)[to[0]][to[1]] = (*grid)[to[0]][to[1]+1]
			(*grid)[to[0]][to[1]+1] = 0
			return true
		} else if (*grid)[to[0]][to[1]] == -1 {
			return false
		} else if (*grid)[to[0]][to[1]] == 1 || grid[to[0]][to[1]] == 2 {
			if moveRe(grid, [2]int{to[0], to[1] - 1}, move) {
				(*grid)[to[0]][to[1]] = (*grid)[to[0]][to[1]+1]
				(*grid)[to[0]][to[1]+1] = 0
				return true
			} else {
				return false
			}
		}

	case RIGHT:

		if (*grid)[to[0]][to[1]] == 0 {
			(*grid)[to[0]][to[1]] = (*grid)[to[0]][to[1]-1]
			(*grid)[to[0]][to[1]-1] = 0
			return true
		} else if (*grid)[to[0]][to[1]] == -1 {
			return false
		} else if (*grid)[to[0]][to[1]] == 1 || grid[to[0]][to[1]] == 2 {
			if moveRe(grid, [2]int{to[0], to[1] + 1}, move) {

				(*grid)[to[0]][to[1]] = (*grid)[to[0]][to[1]-1]
				(*grid)[to[0]][to[1]-1] = 0
				return true
			} else {
				return false
			}
		}
	}
	return false
}

func main() {
	file, _ := os.Open("./input.txt")
	scanner := bufio.NewScanner(file)

	var grid [SIZE][SIZE * 2]int
	var pos [2]int

	for i := range SIZE {
		scanner.Scan()
		line := scanner.Text()
		for j := range SIZE {
			if line[j] == '#' {
				grid[i][j*2] = -1
				grid[i][(j*2)+1] = -1
			} else if line[j] == 'O' {
				grid[i][j*2] = 1
				grid[i][(j*2)+1] = 2
			} else if line[j] == '@' {
				grid[i][j*2] = 3
				pos[0] = i
				pos[1] = j * 2
			}
		}
	}

	scanner.Scan()

	// GRID SHOWCASE
	/*
		for a := range SIZE {
			for b := range SIZE * 2 {
				if grid[a][b] == -1 {
					fmt.Print("#")
				} else if grid[a][b] == 0 {
					fmt.Print(".")
				} else if grid[a][b] == 1 {
					fmt.Print("[")
				} else if grid[a][b] == 2 {
					fmt.Print("]")
				} else if grid[a][b] == 3 {
					fmt.Print("@")
				}
			}
			fmt.Print("\n")
		}
		//*/
	for scanner.Scan() {
		line := scanner.Text()
		for j := range len(line) {
			if line[j] == '<' {
				if moveRe(&grid, [2]int{pos[0], pos[1] - 1}, LEFT) {
					pos[1]--
				}
			} else if line[j] == '>' {
				if moveRe(&grid, [2]int{pos[0], pos[1] + 1}, RIGHT) {
					pos[1]++
				}
			} else if line[j] == '^' {
				if moveRe(&grid, [2]int{pos[0] - 1, pos[1]}, UP) {
					pos[0]--
				}
			} else if line[j] == 'v' {
				if moveRe(&grid, [2]int{pos[0] + 1, pos[1]}, DOWN) {
					pos[0]++
				}
			}

			// GRID SHOWCASE
			/*
				fmt.Printf("%c\n", line[j])
				for a := range SIZE {
					for b := range SIZE * 2 {
						if grid[a][b] == -1 {
							fmt.Print("#")
						} else if grid[a][b] == 0 {
							fmt.Print(".")
						} else if grid[a][b] == 1 {
							fmt.Print("[")
						} else if grid[a][b] == 2 {
							fmt.Print("]")
						} else if grid[a][b] == 3 {
							fmt.Print("@")
						}
					}
					fmt.Print("\n")
				}
				time.Sleep(50 * time.Millisecond)
				//		*/
		}
	}

	var value int
	for i := range SIZE {
		for j := range SIZE * 2 {
			if grid[i][j] == 1 {
				value += (100 * i) + j
			}
		}
	}
	fmt.Println(value)
}
