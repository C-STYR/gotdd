package main

func Sum(numbers []int) int {
	sum := 0
	for _, num := range numbers {
		sum += num
	}
	return sum
}

func SumAll(numsToSum ...[]int) []int {
	output := []int{}
	for _, slc := range numsToSum {
		output = append(output, Sum(slc))
	}
	return output
}

func SumAllTails(numsToSum ...[]int) []int {
	output := []int{}
	for _, slc := range numsToSum {
		if len(slc) == 0 {
			output = append(output, 0)
		} else {
			output = append(output, Sum(slc[1:]))
		}
	}
	return output
}
