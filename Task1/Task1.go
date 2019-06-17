package main

import (
	"fmt"
	"github.com/skorobogatov/input"
)

const (
	WHITE = iota
	GREEN
	BLACK
)

type AutomaticMiles struct {
	n, m, start  int
	inputMatrix  [][]int
	outputMatrix [][]string
}

func (mat *AutomaticMiles) input() {
	input.Scanf(" %d %d %d ", &mat.n, &mat.m, &mat.start)
	mat.inputMatrix, mat.outputMatrix = make([][]int, mat.n), make([][]string, mat.n)

	for i := 0; i < mat.n; i++ {
		mat.inputMatrix[i] = make([]int, mat.m)
		for j := 0; j < mat.m; j++ {
			input.Scanf(" %d", &mat.inputMatrix[i][j])
		}
	}
	//input.Scanf("\n")
	for i := 0; i < mat.n; i++ {
		mat.outputMatrix[i] = make([]string, mat.m)
		for j := 0; j < mat.m; j++ {
			input.Scanf(" %s", &mat.outputMatrix[i][j])
		}
	}
}

func (mat *AutomaticMiles) output() {
	fmt.Println(mat.n)
	fmt.Println(mat.m)
	fmt.Println(mat.start)
	for i := 0; i < mat.n; i++ {
		for j := 0; j < mat.m; j++ {
			fmt.Print(mat.inputMatrix[i][j], " ")
		}
		fmt.Println()
	}

	for i := 0; i < mat.n; i++ {
		for j := 0; j < mat.m; j++ {
			fmt.Print(string(mat.outputMatrix[i][j]), " ")
		}
		fmt.Println()
	}
}

func (mat *AutomaticMiles) algorithm() (res AutomaticMiles) {
	index, v := 0, mat.start
	color := make([]int, mat.n)
	newIndex := make([]int, mat.n)

	res.inputMatrix, res.outputMatrix = make([][]int, mat.n), make([][]string, mat.n)

	var visitVertex func(int)

	visitVertex = func(v int) {
		color[v] = GREEN
		newIndex[v] = index
		for i := 0; i < mat.m; i++ {
			u := mat.inputMatrix[v][i]
			if color[u] == WHITE {
				index++
				visitVertex(u)
			}
			newV, newU, sym := newIndex[v], newIndex[u], mat.outputMatrix[v][i]
			//fmt.Println(sym)
			res.inputMatrix[newV], res.outputMatrix[newV] = append(res.inputMatrix[newV], newU), append(res.outputMatrix[newV], sym)
		}
		color[v] = BLACK
	}

	visitVertex(v)
	res.n, res.m, res.start = index+1, mat.m, 0
	return
}

func main() {
	var res AutomaticMiles
	res.input()
	k := res.algorithm()
	//res.output()
	k.output()
	//fmt.Println(res.outputMatrix)
}
