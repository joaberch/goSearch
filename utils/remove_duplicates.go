package utils

// RemoveDuplicates returns a new slice containing the first occurrence of each integer from
// lines, preserving the original order. The input slice is not modified and the operation
// runs in O(n) time using O(n) extra space.
func RemoveDuplicates(lines []int) []int {
	seen := make(map[int]bool)
	var result []int
	for _, line := range lines {
		if !seen[line] {
			result = append(result, line)
			seen[line] = true
		}
	}
	return result
}
