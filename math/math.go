package math

func SumSlice(numbers []int) int {
	sum := 0
	for _, num := range numbers {
		sum += num
	}
	return sum
}

func MultiplyMapValues(mapper map[string]int) int {
	total := 1
	for _, val := range mapper {
		total *= val
	}
	return total
}
