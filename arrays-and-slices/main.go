package arraysandslices

func Sum(numbers []int) int {
	sum := 0
	for _, n := range numbers {
		sum += n
	}

	return sum
}

func SumAll(slices ...[]int) []int {
	var res []int
	for _, item := range slices {
		res = append(res, Sum(item))
	}

	return res
}

func SumAllTails(slices ...[]int) []int {
	var res []int
	for _, item := range slices {
		if len(item) == 0 {
			res = append(res, 0)
			continue
		}

		res = append(res, Sum(item[1:]))
	}

	return res
}
