package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

const (
	inputSize = 1024
)

type input struct {
	arr []byte
	position int
}

func inputInit() (in input){
	info := bufio.NewReader(os.Stdin)
	res, lensSum := make([]*[]byte, 0, 8), 0
	for{
		p := make([]byte, inputSize)
		n, err := info.Read(p)
		if err == io.EOF{
			break
		}
		lensSum += n
		p = p[:n]
		res = append(res, &p)
	}
	result := make([]byte, lensSum)
	for _, a := range res{
		result = append(result, *a...)
	}
	in.arr = result
	return
}

func (in *input) getArrByte() (res []byte){
	res = make([]byte, 0, 64)
	for ; in.position != len(in.arr) && (in.arr[in.position] == ' ' || in.arr[in.position] == '\n' || in.arr[in.position] == 0); in.position++{}
	for ; in.position != len(in.arr) && in.arr[in.position] != ' ' && in.arr[in.position] != '\n'; in.position++{
		res = append(res, in.arr[in.position])
	}
	for ; in.position != len(in.arr) && (in.arr[in.position] == ' ' || in.arr[in.position] == '\n' || in.arr[in.position] == 0); in.position++{}
	return
}

func (in *input) getInt() (res int){
	str := in.getArrByte()
	for index, multiplier := len(str) - 1, 1; index >= 0; index, multiplier = index - 1, multiplier * 10{
		res += int(str[index] - '0') * multiplier
	}
	return
}

func (in *input) getString() string{
	return string(in.getArrByte())
}

type AutomaticMiles struct {
	numberInputAlphabet, numberOutputAlphabet, numberConditions int
	inputAlphabet []string
	outputAlphabet []string
	transitionMatrix [][]int
	outputMatrix [][]string
}

func (auto *AutomaticMiles) input(){
	in := inputInit()
	auto.numberInputAlphabet = in.getInt()
	auto.inputAlphabet = make([]string, auto.numberInputAlphabet)
	for i := 0; i < auto.numberInputAlphabet; i++{
		auto.inputAlphabet[i] = in.getString()
	}
	auto.numberOutputAlphabet = in.getInt()
	auto.outputAlphabet = make([]string, auto.numberOutputAlphabet)
	for i := 0; i < auto.numberOutputAlphabet; i++{
		auto.outputAlphabet[i] = in.getString()
	}
	auto.numberConditions = in.getInt()
	auto.transitionMatrix, auto.outputMatrix = make([][]int, auto.numberConditions), make([][]string, auto.numberConditions)
	for i := 0; i < auto.numberConditions; i++{
		auto.transitionMatrix[i] = make([]int, auto.numberInputAlphabet)
		for j := 0; j < auto.numberInputAlphabet; j++{
			auto.transitionMatrix[i][j] = in.getInt()
		}
	}
	for i := 0; i < auto.numberConditions; i++{
		auto.outputMatrix[i] = make([]string, auto.numberInputAlphabet)
		for j := 0; j < auto.numberInputAlphabet; j++{
			auto.outputMatrix[i][j] = in.getString()
		}
	}
}

type Pair struct {
	a int
	b string
}

func transformation(auto AutomaticMiles){
	vertex := make([]map[string]int, auto.numberConditions)
	vertexs := make([]Pair, 0, auto.numberConditions * auto.numberInputAlphabet)
	input :=  0
	for i := 0; i < auto.numberConditions; i++{
		for j := 0; j < auto.numberInputAlphabet; j++{
			transition, output := auto.transitionMatrix[i][j], auto.outputMatrix[i][j]
			elem := &vertex[transition]
			if len(*elem) == 0{
				*elem = make(map[string]int, auto.numberInputAlphabet)
			}
			_, ok := (*elem)[output]
			if !ok{
				(*elem)[output] = input
				vertexs = append(vertexs, Pair{a: transition, b: output})
				input++
			}
		}
	}
	//fmt.Println(vertexs)
	edges := make([][]Pair, input)
	for i, mapa := range vertex{
		arrTransition, arrOutput := &auto.transitionMatrix[i], &auto.outputMatrix[i]
		for _, value := range mapa{
			numVertex := &edges[value]
			*numVertex = make([]Pair, auto.numberInputAlphabet)
			for j := 0; j < auto.numberInputAlphabet; j++{
				(*numVertex)[j] = Pair{a: vertex[(*arrTransition)[j]][(*arrOutput)[j]], b: auto.inputAlphabet[j]}
			}
		}
	}
	//fmt.Println(edges)
	fmt.Println(`digraph {
	rankdir = LR`)
	for numVertex, vertex := range vertexs{
		fmt.Print("	", numVertex, " [label = \"(", vertex.a, ",", vertex.b, ")\"]\n")
		for _, u := range edges[numVertex]{
			fmt.Print("	", numVertex, " -> ", u.a, " [label = \"", u.b, "\"]\n")
		}
	}
	fmt.Println("}")
}

func main(){
	var mat AutomaticMiles
	mat.input()
	transformation(mat)
	//fmt.Println(mat)
}