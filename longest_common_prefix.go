package codingquestions

import "fmt"

func main() {
	fmt.Println(longestCommonPrefix([]string{"aacc", "aa", "aa", "aa", "aaca"}))
}

func longestCommonPrefix(strs []string) string {
	// Edge case: If the input array is empty, return an empty string
	if len(strs) == 0 {
		return ""
	}

	// Assume the first string is the longest common prefix initially
	prefix := strs[0]
	// Iterate through each string in the array
	for _, str := range strs[1:] {
		if len(str) == 0 {
			prefix = ""
		}
		if len(prefix) > len(str) {
			prefix = prefix[:len(str)]
		}
		// Compare the current string with the prefix
		for i := 0; i < len(prefix) && i < len(str); i++ {
			if prefix[i] != str[i] {
				// Found a mismatch, reduce the prefix length
				prefix = prefix[:i]
				break
			}
		}
		// If prefix becomes empty, we can return early
		if prefix == "" {
			return ""
		}
	}

	return prefix
}
