package main

import (
	"fmt"
)

const (
	ALIVE  = true
	DEAD   = false
	WIDTH  = 50
	HEIGHT = 50
)

func createWindow(row int, col int) [][]int {

	window := make([][]int, HEIGHT)
	for i := range window {
		window[i] = make([]int, WIDTH)
	}

	return window

}

func initWindow(window [][]int) {
	height := len(window)
	width := len((window)[0])
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			window[i][j] = 0
		}
	}

}

func printArray(window *[][]int) {
	for i := range *window {
		for _, j := range (*window)[i] {
			fmt.Print(j, " ")
		}
		fmt.Println()
	}
}
func main() {

	fmt.Println("**HELLO THIS IS MY PROJECT IN GO FOR GAME OF LIFE")
	window := createWindow(WIDTH, HEIGHT)
	initWindow(window)
	printArray(&window)

}

/*Any live cell with fewer than two live neighbours dies, as if by underpopulation.
Any live cell with two or three live neighbours lives on to the next generation.
Any live cell with more than three live neighbours dies, as if by overpopulation.
Any dead cell with exactly three live neighbours becomes a live cell, as if by reproduction.*/
