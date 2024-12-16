package codingquestions

import "fmt"

func main() {
	m1 := make(map[int64]int64)
	m1[0] = 23
	m1[1] = 134
	m2 := make(map[float64]float64)
	m2[0] = 23.55
	m2[1] = 67.4446
	fmt.Println(SumIntsOrFloats(m1))
	fmt.Println(SumIntsOrFloats(m2))

}

func SumIntsOrFloats[K comparable, V int64 | float64](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}
