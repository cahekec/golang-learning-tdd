package main

import "fmt"

func main() {
	result := SumAll([]int{1, 3}, []int{0, 9}, []int{1, 8})
	fmt.Println(result)
}

func Sum(n []int) int {
	r := 0

	for _, num := range n {
		r += num
	}

	return r
}

func SumAll(numbersToSum ...[]int) (sums []int) {
	for _, nums := range numbersToSum {
		sums = append(sums, Sum(nums))
	}

	return sums
}

func SumAllTails(numbersToSum ...[]int) (sums []int) {
	for _, nums := range numbersToSum {
		isEmpty := len(nums) == 0

		if isEmpty {
			sums = append(sums, 0)
			continue
		}

		tail := nums[1:]
		sums = append(sums, Sum(tail))
	}

	return sums
}
