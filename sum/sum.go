package sum

import (
	//"fmt"
)

func Sum(numbers []int) int {
	sum := 0

	for _, numbers := range numbers {
		sum += numbers
	}
	return sum
}

func SumAllTails(numbers ...[]int) []int {
	var sums []int
	for _, number := range numbers {
		if len(number) == 0 {
			sums = append(sums, 0)
		} else { 
			tail := number[1:]
			sums = append(sums, Sum(tail))
		}
	}
	return sums
}

/*func SumAllTails(numbers ...[]int) []int {
	return nil
}*/