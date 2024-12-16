package codingquestions

import "fmt"

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 19))
}

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for k, v := range nums {
		m[v] = k
	}
	arr := make([]int, 0)
	for i, val := range nums {
		chk := target - val
		if _, ok := m[chk]; ok {
			if i == m[chk] {
				continue
			}
			arr = append(arr, i)
			arr = append(arr, m[chk])
			return arr
		}
	}
	return arr
}
