package main

import (
	"fmt"
	"github.com/skorobogatov/input"
)

type AutomaticMiles struct {
	n, m, start int
	inputMatrix [][]int
	outputMatrix [][]byte
}

func (mat *AutomaticMiles) input(){
	input.Scanf(" %d %d %d", &mat.n, &mat.m, &mat.start)
	mat.inputMatrix, mat.outputMatrix = make([][]int, mat.n), make([][]byte, mat.n)

	for i := 0; i < mat.n; i++ {
		mat.inputMatrix[i] = make([]int, mat.m)
		for j := 0; j < mat.m; j++ {
			input.Scanf(" %d", &mat.inputMatrix[i][j])
		}
	}

	for i := 0; i < mat.n; i++ {
		mat.outputMatrix[i] = make([]byte, mat.m)
		for j := 0; j < mat.m; j++ {
			input.Scanf(" %c", &mat.outputMatrix[i][j])
		}
	}
}

func (mat *AutomaticMiles) output() {
	fmt.Println(`digraph {
    rankdir = LR
    dummy [label = "", shape = none]`)
	for i := 0; i < mat.n; i++{
		fmt.Println("   ", i, "[shape = circle]")
	}
	fmt.Println("    dummy ->", mat.start)
	for i := 0; i < mat.n; i++{
		for j := 0; j < mat.m; j++{
			u := mat.inputMatrix[i][j]
			fmt.Print("    ", i, " -> ", u, " [label = \"", string(j + int('a')) + "(" + string(mat.outputMatrix[i][j]) + ")", "\"]")
			fmt.Println()
		}
	}
	fmt.Println("}")
}

func main(){
	var mat AutomaticMiles
	mat.input()
	mat.output()
}
