package codingquestions

import (
	"fmt"
	"rsc.io/quote"
)

type Pair struct {
	num1 int
	num2 int
}

func main() {
	fmt.Println(quote.Opt())
	result := make([]Pair, 0)
	//input := []int{11, 8, 11, 8, 11}
	//10 - 9,9,9
	//9 - 10,10,10
	input := []int{19, 0, 10, 9, 10, 9, 10, 9}
	sum := 19
	m := make(map[int]int)
	for _, i := range input {
		if _, ok := m[i]; !ok {
			m[i] = 1
		} else {
			m[i]++
		}
	}

	for _, val := range input {
		//fmt.Println(val)
		chk := sum - val

		if op, ok := m[chk]; ok {
			//fmt.Println("Pair ", chk, "Sum ", op)
			if op > 0 {
				l := Pair{
					num1: val,
					num2: chk,
				}
				m[val]--
				m[chk]--
				result = append(result, l)
			}
			//fmt.Println(ok)
		}
	}
	fmt.Println(result)
}
