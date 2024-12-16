package codingquestions

import "fmt"

func main() {
	fmt.Print(romanToInt("III"))
	fmt.Print(romanToInt("LVIII"))
	fmt.Print(romanToInt("MCMXCIV"))
}

func romanToInt(s string) int {
	m := make(map[rune]int)
	m['I'] = 1
	m['V'] = 5
	m['X'] = 10
	m['L'] = 50
	m['C'] = 100
	m['D'] = 500
	m['M'] = 1000
	result := 0
	var prev rune
	for _, l := range s {
		result = result + m[rune(l)]
		if (m[rune(l)] == 5 || m[rune(l)] == 10) && prev == 'I' {
			result = result - m['I']*2
		}
		if (m[rune(l)] == 50 || m[rune(l)] == 100) && prev == 'X' {
			result = result - m['X']*2
		}
		if (m[rune(l)] == 500 || m[rune(l)] == 1000) && prev == 'C' {
			result = result - m['C']*2
		}
		prev = rune(l)
	}
	return result
}
