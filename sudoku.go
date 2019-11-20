package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

func random(min, max int) int {
	z := rand.Intn(max)
	if z < min {
		z = min
	}
	return z
}

func check(n *[]int) {
	nest := *n
	insertion := random(1, 10)
	for i := range nest {
		if nest[i] == insertion {
			return
		}
	}
	nest = append(nest, insertion)
	*n = nest
}

func checkS(s *[][]int, nest []int) {
	slice := *s
	for i := range slice {
		for ia := range slice[i] {
			if nest[ia] == slice[i][ia] {
				return
			}
		}
	}
	slice = append(slice, nest)
	*s = slice
}

func fillNest() []int {
	nest := []int{}
	counter := 0
	for counter != 1 {
		if len(nest) == 9 {
			counter = 1
		}
		check(&nest)
	}
	return nest
}

func fillSlice() [][]int {
	slice := [][]int{}
	counter := 0
	for counter != 1 {
		if len(slice) == 9 {
			counter = 1
		}
		nest := fillNest()
		checkS(&slice, nest)
	}
	return slice
}

func printSlice(slice [][]int) {
	for i := range slice {
		fmt.Println(slice[i])
	}
}

func dup(slice [][]int) [][]int {
	duplicate := make([][]int, len(slice))
	for i := range slice {
		duplicate[i] = make([]int, len(slice[i]))
		copy(duplicate[i], slice[i])
	}
	return duplicate
}

func deduct(s [][]int) {
	slice := s
	counter := 0
	for counter = 0; counter < 60; counter++ {
		slice[random(0, 8)][random(0, 8)] = 0
	}
	printSlice(slice)
}

func main() {
	Scanner := bufio.NewScanner(os.Stdin)
	rand.Seed(time.Now().UTC().UnixNano())
	slice1 := fillSlice()
	sudoku := dup(slice1)
	deduct(sudoku)
	Scanner.Scan()
	fmt.Println("")
	printSlice(slice1)
	Scanner.Scan()
}
