package codingquestions

import (
	"fmt"
	"math/rand"
)

func main() {
	input1 := 100
	input2 := 20
	input3 := 3
	arr := make([]int, 0)
	//var arr []int
	var inc int
	for i := 0; i < input3; i++ {
		if i > 0 {
			inc = generateRandom(input2, input1)
		} else {
			inc = input2
		}
		for k := 0; k <= input3-1; k++ {
			arr = Push(arr, inc)
			inc++
		}
		fmt.Println(arr)
		arr = Pop(arr)
		fmt.Println(arr)
		arr = *new([]int)
	}
}

func Push(arr []int, val int) []int {
	arr = append(arr, val)
	return arr
}

func Pop(arr []int) []int {
	arr = arr[0 : len(arr)-1]
	return arr
}

func generateRandom(min, max int) int {
	for {
		k := rand.Intn(max)
		if k < min {
			continue
		}
		return k
	}
}
