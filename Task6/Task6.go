package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

const (
	inputSize = 1024 //кол-во байт для ввода
	sizes     = 4    //кол-во смежны вершин в матрице перехода
	stackSize = 64
)

type input struct {
	arr      []byte
	position int
}

func inputInit() (in input) {
	info := bufio.NewReader(os.Stdin)
	res, lensSum := make([]*[]byte, 0, 8), 0
	for {
		p := make([]byte, inputSize)
		n, err := info.Read(p)
		if err == io.EOF {
			break
		}
		lensSum += n
		p = p[:n]
		res = append(res, &p)
	}
	result := make([]byte, lensSum)
	for _, a := range res {
		result = append(result, *a...)
	}
	in.arr = result
	return
}

func (in *input) getArrByte() (res []byte) {
	res = make([]byte, 0, 64)
	for ; in.position != len(in.arr) && (in.arr[in.position] == ' ' || in.arr[in.position] == '\n' || in.arr[in.position] == 0); in.position++ {
	}
	for ; in.position != len(in.arr) && in.arr[in.position] != ' ' && in.arr[in.position] != '\n'; in.position++ {
		res = append(res, in.arr[in.position])
	}
	for ; in.position != len(in.arr) && (in.arr[in.position] == ' ' || in.arr[in.position] == '\n' || in.arr[in.position] == 0); in.position++ {
	}
	return
}

func (in *input) getInt() (res int) {
	str := in.getArrByte()
	for index, multiplier := len(str)-1, 1; index >= 0; index, multiplier = index-1, multiplier*10 {
		res += int(str[index]-'0') * multiplier
	}
	return
}

func (in *input) getString() string {
	return string(in.getArrByte())
}

type Vertex []int //вершина для автомата

//стэк
type Stack []Vertex

//инициализация стэка
func initStack(size int) Stack {
	return make(Stack, 0, size)
}

//добавление элемента в стэк
func (st *Stack) push(value Vertex) {
	*st = append(*st, value)
}

//удаление элемента из стэка
func (st *Stack) pop() (value Vertex) {
	value = (*st)[len(*st)-1]
	*st = (*st)[:len(*st)-1]
	return
}

//пуст ли стэк?
func (st *Stack) isEmpty() bool {
	return len(*st) == 0
}

//недетерминированный распознающий автомат
type recognizingNonDeterministicMachine struct {
	transitionFunction []map[string][]int //матрица переходов
	N, M, q0           int                //кол-во состояний, кол-во переходов, начальное состояние
	Final              []bool             //евляется ли состояние принимающим
	alphabet           []string    //алфавит
}

//инициализация недетерминированный распознающий автомат
func (auto *recognizingNonDeterministicMachine) initRecognizingNonDeterministicMachine(N, M int) {
	auto.transitionFunction, auto.Final, auto.alphabet = make([]map[string][]int, N), make([]bool, N), make([]string, 0, sizes)
	for i := range auto.transitionFunction{
		auto.transitionFunction[i] = make(map[string][]int, sizes)
	}
}

//ввод недетерминированный распознающий автомат
func (auto *recognizingNonDeterministicMachine) inputRecognizingNonDeterministicMachine() {
	in := inputInit() //иницализация ввода
	auto.N, auto.M = in.getInt(), in.getInt()
	auto.initRecognizingNonDeterministicMachine(auto.N, auto.M)
	var a, b int
	var c string
	for i := 0; i < auto.M; i++ {
		a, b, c = in.getInt(), in.getInt(), in.getString()
		if c != "lambda" {
			flag := true
			for _, word := range auto.alphabet{
				if c == word{
					flag = false
					break
				}
			}
			if flag{
				auto.alphabet = append(auto.alphabet, c)
			}
		}
		_, ok := auto.transitionFunction[a][c]
		if !ok {
			auto.transitionFunction[a][c] = make([]int, 0, sizes)
		}
		auto.transitionFunction[a][c] = append(auto.transitionFunction[a][c], b)
	}
	for i := 0; i < auto.N; i++ {
		a = in.getInt()
		if a == 1 {
			auto.Final[i] = true
		}
	}
	auto.q0 = in.getInt()
}

func (auto *recognizingNonDeterministicMachine) closure(z Vertex) (res Vertex) {
	res, boolRes := make(Vertex, 0, sizes), make([]bool, auto.N)
	var dfs func(int)

	dfs = func(q int) {
		if !boolRes[q] {
			boolRes[q] = true
			for _, w := range auto.transitionFunction[q]["lambda"] {
				dfs(w)
			}
		}
	}

	for _, q := range z {
		dfs(q)
	}

	for iter, elem := range boolRes {
		if elem {
			res = append(res, iter)
		}
	}

	return
}

func (auto *recognizingNonDeterministicMachine) det() {
	var search func(Vertex, []Vertex) int
	search = func(vertex Vertex, arr []Vertex) int {
		for i, elem := range arr {
			if len(elem) == len(vertex) {
				l := true
				for i := 0; i < len(elem); i++ {
					if elem[i] != vertex[i] {
						l = false
						break
					}
				}
				if l {
					return i
				}
			}
		}
		return -1
	}

	q0 := auto.closure([]int{auto.q0})
	Q, stack, f := make([]Vertex, 0, stackSize), initStack(stackSize), make([]bool, 0, sizes)
	transitionFunction := make([]map[int][]string, 0, stackSize)
	Q = append(Q, q0)
	transitionFunction = append(transitionFunction, make(map[int][]string, sizes))
	f = append(f, false)
	stack.push(q0)
	for !stack.isEmpty() {
		z := stack.pop()
		index := search(z, Q)
		for _, u := range z {
			if auto.Final[u] {
				f[index] = true
				break
			}
		}
		for _, key := range auto.alphabet {
			mn := make(Vertex, 0, sizes)
			for _, u := range z {
				arr, ok := auto.transitionFunction[u][key]
				if ok {
					mn = append(mn, arr...)
				}
			}
			v := auto.closure(mn)
			if search(v, Q) == -1 {
				Q = append(Q, v)
				transitionFunction = append(transitionFunction, make(map[int][]string, sizes))
				f = append(f, false)
				stack.push(v)
			}
			transitionFunction[index][search(v, Q)] = append(transitionFunction[index][search(v, Q)], key)
		}
	}
	fmt.Println(`digraph {
	rankdir = LR
	dummy [label = "", shape = none]`)
	for i, elem := range Q {
		if f[i] {
			fmt.Print("	", i, " [label = \"", elem, "\", shape = doublecircle]\n")
		} else {
			fmt.Print("	", i, " [label = \"", elem, "\", shape = circle]\n")
		}
	}
	fmt.Println("	dummy -> 0")
	for i, arrs := range transitionFunction {
		for j, arr := range arrs {
			fmt.Print("	", i, " -> ", j, " [label = \"")
			for i, word := range arr {
				if i < len(arr) - 1 {
					fmt.Print(word, ", ")
				} else {
					fmt.Print(word)
				}
			}
			fmt.Print("\"]\n")
		}
	}
	fmt.Print("}\n")
}

func main() {
	var res recognizingNonDeterministicMachine
	res.inputRecognizingNonDeterministicMachine()
	res.det()
	//fmt.Println(res)
	//fmt.Println(res.closure([]int{3}))
}
