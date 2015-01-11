package main

import (
	"github.com/nsf/termbox-go"
	"math/rand"
	"time"
)

const (
	boardWidth = 79
	boardHeight = 30
)

type Cell struct {
	alive  bool
}

func (c Cell) Alive() bool{
	return c.alive
}

type Grid struct {
	currentGrid [][]Cell
	previousGrid [][]Cell
}

func(g *Grid) SetCurrentGrid(grid [][]Cell){
	g.currentGrid = grid
}

func(g *Grid) SetPreviousGrid(grid [][]Cell){
	g.previousGrid = grid
}

func (g Grid) CurrentGrid() ([][]Cell) {
	return g.currentGrid
}

func (g Grid) PreviousGrid() ([][]Cell) {
	return g.previousGrid
}

func (g *Grid) init (){
	g.currentGrid = g.initGrid()
}

func (g *Grid) initGrid() ([][]Cell) {
	var newLife = make([][]Cell, boardWidth)
	for i := range newLife {
		newLife[i] = make([]Cell, boardHeight)
	}
	return newLife
}

func (g *Grid) findAliveNeighbors(y int, x int) (int){
	board := g.CurrentGrid()
	aliveNeighbors := 0

	if y >= 1 && x >= 1 && board[y - 1][x - 1].Alive() {
		aliveNeighbors++
	}
	if y >= 1 && x <= boardHeight && board[y - 1][x + 1].Alive() {
		aliveNeighbors++
	}
	if y <= boardWidth && x >= 1 && board[y + 1][x - 1].Alive() {
		aliveNeighbors++
	}
	if y <= boardWidth && x <= boardHeight && board[y + 1][x + 1].Alive() {
		aliveNeighbors++
	}
	if y >= 1 && board[y - 1][x].Alive() {
		aliveNeighbors++
	}
	if y <= boardWidth &&  board[y + 1][x].Alive() {
		aliveNeighbors++
	}
	if x >= 1 && board[y][x-1].Alive() {
		aliveNeighbors++
	}
	if x <= boardHeight && board[y][x+1].Alive() {
		aliveNeighbors++
	}
	
	return aliveNeighbors
}

func (g *Grid) genLife() ([][]Cell){
	newLife := g.initGrid()
	for y:= 1; y < boardWidth; y++ {
		for x := 1; x < boardHeight; x++ {
			currentY := y - 1
			currentX := x - 1
			aliveNeighbors := g.findAliveNeighbors(currentY, currentX)
			currentCell := g.CurrentGrid()[currentY][currentX]

			if currentCell.Alive() && (aliveNeighbors == 2 || aliveNeighbors == 3) {
				cell := Cell{true}
				newLife[y-1][x-1] = cell
			} else if currentCell.Alive() == false && aliveNeighbors == 3 {
				cell := Cell{true}
				newLife[currentY][currentX] = cell
			} else {
				cell := Cell{false}
				newLife[currentY][currentX] = cell
			}
		}
	}
	g.SetPreviousGrid(g.CurrentGrid())
	g.SetCurrentGrid(newLife)
	return newLife
}
func (g *Grid) drawLife() {
	currentLife := g.CurrentGrid()
	// Draws the box
	termbox.SetCell(0, 0, 0x250C, termbox.ColorWhite, termbox.ColorBlack)
	termbox.SetCell(boardWidth + 1, 0, 0x2510, termbox.ColorWhite, termbox.ColorBlack)
	termbox.SetCell(0, boardHeight + 1, 0x2514, termbox.ColorWhite, termbox.ColorBlack)
	termbox.SetCell(boardWidth + 1, boardHeight + 1, 0x2515, termbox.ColorWhite, termbox.ColorBlack)

	for i := 1; i < 80; i++ {
		termbox.SetCell(i, 0, 0x2500, termbox.ColorWhite, termbox.ColorBlack)
		termbox.SetCell(i, 31, 0x2500, termbox.ColorWhite, termbox.ColorBlack)
	}
	for i := 1; i < 31; i++ {
		termbox.SetCell(0, i, 0x2502, termbox.ColorWhite, termbox.ColorBlack)
		termbox.SetCell(80, i, 0x2502, termbox.ColorWhite, termbox.ColorBlack)
	}

	for y:= 0; y < boardWidth; y++ {
		for x := 0; x < 30; x++ {
			currentCell := currentLife[y][x]
			
			if currentCell.Alive() {
				termbox.SetCell(y + 1, x + 1, '*', termbox.ColorWhite, termbox.ColorBlack)
			} else {
				termbox.SetCell(y + 1, x + 1, ' ', termbox.ColorWhite, termbox.ColorBlack)
			}


		}
	}

}

func (g *Grid) seedLife() {
	for i := 0; i < ( boardWidth * boardHeight / 2); i++ {
		g.currentGrid[rand.Intn(boardWidth)][rand.Intn(boardHeight)] = Cell{true}
	}
}

func main() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()

	termbox.SetInputMode(termbox.InputEsc | termbox.InputMouse)
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)

	g := Grid{}
	g.init()
	g.seedLife()
	g.drawLife()
	termbox.Flush()

	for i := 0; i < 500; i++ {
		termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
		g.genLife()
		g.drawLife()
		termbox.Flush()
		time.Sleep(time.Second / 10)
	}

}
