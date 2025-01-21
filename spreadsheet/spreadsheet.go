package spreadsheet

import (
	"strconv"
	"strings"
)

type Cell struct {
	x int
	y int
}

type Excel struct {
	height int
	width  int

	mat     [][]int
	adjlist map[Cell][]Cell
}

func Constructor(height int, width byte) Excel {
	mat := make([][]int, height)
	for i := 0; i < height; i++ {
		mat[i] = make([]int, width-'A'+1)
	}

	adjlist := make(map[Cell][]Cell)

	return Excel{
		height:  height,
		width:   int(width-'A') + 1,
		mat:     mat,
		adjlist: adjlist,
	}
}

func (this *Excel) Set(row int, column byte, val int) {
	target := Cell{x: row - 1, y: int(column - 'A')}

	// remove old formula deps if any
	for src, neighbors := range this.adjlist {
		newNeighbors := []Cell{}
		for _, nbr := range neighbors {
			if nbr != target {
				newNeighbors = append(newNeighbors, nbr)
			}
		}
		this.adjlist[src] = newNeighbors
	}

	// now set the new value
	preval := this.mat[target.x][target.y]
	this.mat[target.x][target.y] = val

	this.propogateDiff(target, target, val-preval)
}

func (this *Excel) propogateDiff(cell Cell, sumcell Cell, diff int) {
	if cell != sumcell {
		this.mat[cell.x][cell.y] += diff
	}

	for _, cellNeigh := range this.adjlist[cell] {
		this.propogateDiff(cellNeigh, sumcell, diff)
	}
}

func (this *Excel) Get(row int, column byte) int {
	col := int(column - 'A')
	return this.mat[row-1][col]
}

func (this *Excel) GetFirstNRows(n int) [][]int {
	firstNRows := [][]int{}
	for i := 0; i < n; i++ {
		row := []int{}
		for j := 0; j < this.width; j++ {
			row = append(row, this.mat[i][j])
		}
		firstNRows = append(firstNRows, row)
	}
	return firstNRows
}

func (this *Excel) Sum(row int, column byte, numbers []string) int {
	// define the sum cell
	// get all the cells in numbers
	// remove all edges to sumcell then add. This should handle update sum formula
	// build graph directed edges from cells to sum cell
	// need to calculate sum over the cells and put that into the sumcell
	// propogate the changes in set method whenever a cell gets updated

	sumcell := Cell{x: row - 1, y: int(column - 'A')}

	for cell, cellNeighbors := range this.adjlist {
		newNeighbors := []Cell{}
		for _, cellneighbor := range cellNeighbors {
			if cellneighbor != sumcell {
				newNeighbors = append(newNeighbors, cellneighbor)
			}
		}
		this.adjlist[cell] = newNeighbors
	}

	cells := []Cell{}
	for _, number := range numbers {
		numberCells := this.getCells(number)
		cells = append(cells, numberCells...)
	}

	sum := 0
	for _, cell := range cells {
		this.adjlist[cell] = append(this.adjlist[cell], sumcell)
		sum += this.mat[cell.x][cell.y]
	}

	this.setValueNoClear(row, column, sum)

	return sum
}

func (this *Excel) setValueNoClear(row int, column byte, val int) {
	target := Cell{x: row - 1, y: int(column - 'A')}

	// now set the new value
	preval := this.mat[target.x][target.y]
	this.mat[target.x][target.y] = val

	this.propogateDiff(target, target, val-preval)
}

func (this *Excel) getCells(number string) []Cell {
	// "ColRow" , "ColRow1:ColRow2"
	strs := strings.Split(number, ":")

	if len(strs) == 1 {
		x, _ := strconv.Atoi(strs[0][1:])
		cell := Cell{
			x: x - 1,
			y: int(strs[0][0] - 'A'),
		}
		return []Cell{cell}
	}

	cells := []Cell{}

	cell1x, _ := strconv.Atoi(strs[0][1:])
	cell1 := Cell{
		x: cell1x - 1,
		y: int(strs[0][0] - 'A'),
	}

	cell2x, _ := strconv.Atoi(strs[1][1:])
	cell2 := Cell{
		x: cell2x - 1,
		y: int(strs[1][0] - 'A'),
	}

	for i := cell1.x; i <= cell2.x; i++ {
		for j := cell1.y; j <= cell2.y; j++ {
			cells = append(cells, Cell{x: i, y: j})
		}
	}

	return cells
}

/**
 * Your Excel object will be instantiated and called as such:
 * obj := Constructor(height, width);
 * obj.Set(row,column,val);
 * param_2 := obj.Get(row,column);
 * param_3 := obj.Sum(row,column,numbers);
 */
