package codingquestions

import "fmt"

func main() {
	fmt.Println(isPalindrome(121))
}

func isPalindrome(x int) bool {

	val := x
	op := 0
	for x > 0 {
		k := x / 10
		r := x % 10
		x = k
		op = op*10 + r
	}
	fmt.Println(val)
	if op == val {
		return true
	}
	return false
}
