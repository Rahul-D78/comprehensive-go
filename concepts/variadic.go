package concepts

// Variadic function can accept a varying number of arguments of a specific type.
func Sum(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}
