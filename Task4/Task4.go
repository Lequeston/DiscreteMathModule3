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
type Set struct {
	x int
	depth int
	parent *Set
}

func (t *Set) makeSet(x int){
	t.x, t.depth, t.parent = x, 0, t
}

func (t *Set) find() (root *Set){
	if t.parent == t{
		root = t
	} else {
		t.parent = t.parent.find()
		root = t.parent
	}
	return
}

func (t *Set) union(y *Set){
	rootX, rootY := t.find(), y.find()
	if rootX.depth < rootY.depth{
		rootX.parent = rootY
	} else {
		rootY.parent = rootX
		if rootX.depth == rootY.depth && rootX != rootY{
			rootX.depth++
		}
	}
}

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

func (mat *AutomaticMiles) init(n, m, start int){
	mat.n, mat.m, mat.start = n, m, start
	mat.inputMatrix, mat.outputMatrix = make([][]int, n), make([][]byte, n)
	for i := 0; i < n; i++{
		mat.inputMatrix[i], mat.outputMatrix[i] = make([]int, m), make([]byte, m)
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

func (mat *AutomaticMiles) split1() (m int, pi []int){
	m = mat.n
	q, pi := make([]Set, m), make([]int, m)
	for i := 0; i < mat.n; i++{
		q[i].makeSet(i)
	}
	for i := 0; i < mat.n; i++{
		for j := i + 1; j < mat.n; j++{
			if q[i].find() != q[j].find(){
				eq := true
				for p := 0; p < mat.m; p++{
					if mat.outputMatrix[i][p] != mat.outputMatrix[j][p]{
						eq = false
						break
					}
				}
				if eq{
					q[i].union(&q[j])
					m--
				}
			}
		}
	}
	for i := 0; i < mat.n; i++{
		pi[i] = q[i].find().x
	}
	return
}

func (mat *AutomaticMiles) split(pi []int) (int, []int){
	m, q := mat.n, make([]Set, mat.n)
	for i := 0; i < mat.n; i++{
		q[i].makeSet(i)
	}
	for i := 0; i < mat.n; i++ {
		for j := i + 1; j < mat.n; j++ {
			if pi[i] == pi[j] && q[i].find() != q[j].find(){
				eq := true
				for p := 0; p < mat.m; p++{
					if pi[mat.inputMatrix[i][p]] != pi[mat.inputMatrix[j][p]]{
						eq = false
						break
					}
				}
				if eq{
					q[i].union(&q[j])
					m--
				}
			}
		}
	}
	for i := 0; i < mat.n; i++{
		pi[i] = q[i].find().x
	}
	return m, pi
}

func (mat *AutomaticMiles) normalize(){
	var dfs func()

	graphVertex := make([]bool, mat.n)

	dfs = func(){
		color := make([]int, mat.n)

		var visitVertex func(int)

		visitVertex = func(v int){
			color[v] = 1
			for _, u := range mat.inputMatrix[v]{
				if color[u] == 0{
					visitVertex(u)
				}
			}
			graphVertex[v] = true
		}

		visitVertex(mat.start)
	}
	dfs()
	var k int
	kof := make([]int, mat.n)
	//kof[mat.start] = -mat.start
	//kof[0] = mat.start
	//fmt.Println(graphVertex)
	for i, fl := range graphVertex{
		//if i != mat.start {
		kof[i] += k - i
		//}
		if fl{
			mat.inputMatrix[k] = mat.inputMatrix[i]
			mat.outputMatrix[k] = mat.outputMatrix[i]
			if i == mat.start{
				mat.start = mat.start - (i - k)
			}
			k++
		}
	}
	//fmt.Println(kof)
	for i := range mat.inputMatrix{
		for j := range mat.inputMatrix[i]{
			mat.inputMatrix[i][j] += kof[mat.inputMatrix[i][j]]
		}
	}
	//k1 := mat.inputMatrix[mat.start]
	//k2 := mat.outputMatrix[mat.start]
	//mat.inputMatrix[mat.start] = mat.inputMatrix[0]
	//mat.outputMatrix[mat.start] = mat.outputMatrix[0]
	//mat.start = 0
	//mat.inputMatrix[0] = k1
	//mat.outputMatrix[0] = k2
	mat.n = k
}

func (mat *AutomaticMiles) aufenkampHohn() (res AutomaticMiles){

	m, pi := mat.split1()
	//fmt.Println(pi)
	var new int
	for{
		new, pi = mat.split(pi)
		//fmt.Println(pi)
		if new == m{
			break
		}
		m = new
	}
	//new, pi = mat.split(pi)
	//fmt.Println(m)
	//fmt.Println(pi)
	mn := make([]bool, mat.n)
	res.init(m, mat.m, mat.start)
	t := make([]int, mat.n)
	h := 0
	for i := 0; i < mat.n; i++{
		qu := pi[i]
		if !mn[qu] {
			mn[qu] = true
			t[qu] += h
		} else {
			h--
		}
	}
	//fmt.Println(t)
	mn = make([]bool, mat.n)
	for i := 0; i < mat.n; i++{
		qu := pi[i]
		if !mn[qu] {
			mn[qu] = true
			for p := 0; p < mat.m; p++ {
				res.inputMatrix[qu + t[qu]][p] = pi[mat.inputMatrix[i][p]] + t[pi[mat.inputMatrix[i][p]]]
				res.outputMatrix[qu + t[qu]][p] = mat.outputMatrix[i][p]
			}
		}
	}
	//res.inputMatrix[0]
	return
}

func (mat *AutomaticMiles) swapStart(){
	if mat.start != 0{
		for i := range mat.inputMatrix {
			for j := range mat.inputMatrix[i] {
				if mat.inputMatrix[i][j] == mat.start{
					mat.inputMatrix[i][j] = 0
				} else if mat.inputMatrix[i][j] == 0{
					mat.inputMatrix[i][j] = mat.start
				}
			}
		}
		k1 := mat.inputMatrix[mat.start]
		k2 := mat.outputMatrix[mat.start]
		mat.inputMatrix[mat.start] = mat.inputMatrix[0]
		mat.outputMatrix[mat.start] = mat.outputMatrix[0]
		mat.start = 0
		mat.inputMatrix[0] = k1
		mat.outputMatrix[0] = k2
	}
}

func (mat *AutomaticMiles) algorithm() (res AutomaticMiles){
	index, v := 0, mat.start
	color := make([]int, mat.n)
	newIndex := make([]int, mat.n)

	res.inputMatrix, res.outputMatrix = make([][]int, mat.n), make([][]byte, mat.n)

	var visitVertex func(int)

	visitVertex = func(v int){
		color[v] = GREEN
		newIndex[v] = index
		for i := 0; i < mat.m; i++{
			u := mat.inputMatrix[v][i]
			if color[u] == WHITE{
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
	res.n, res.m, res.start = index + 1, mat.m, 0
	return
}

func main(){
	var mat1, mat2 AutomaticMiles
	mat1.input()
	mat2.input()
	//m, pi := mat.split1()
	//fmt.Println(m, pi)
	//m, pi = mat.split(pi)
	//fmt.Println(m, pi)
	//mat.algorithm()
	mat1.swapStart()
	mat2.swapStart()
	k1 := mat1.aufenkampHohn()
	k2 := mat2.aufenkampHohn()
	//k.swapStart()
	//k.normalize()
	k1 = k1.algorithm()
	k2 = k2.algorithm()
	//k = k.algorithm()
	if k1.n == k2.n && k1.m == k2.m{
		flag := true
		for i := range k1.inputMatrix{
			for j := range k1.inputMatrix[i]{
				if k1.inputMatrix[i][j] != k2.inputMatrix[i][j] || k1.outputMatrix[i][j] != k2.outputMatrix[i][j]{
					flag = false
				}
			}
		}
		if flag{
			fmt.Println("EQUAL")
		} else {
			fmt.Println("NOT EQUAL")
		}
	} else {
		fmt.Println("NOT EQUAL")
	}
}
