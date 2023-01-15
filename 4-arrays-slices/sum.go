package main

func Sum(numbers []int) int {
	sum := 0
	// range lets you iterate over an array
	// on each iteration, range returns two values: the index and the value
	// we are choosing to ignore the index value by using _ blank identifier
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SumAll(numbersToSum ...[]int) []int {
	// way 1 with array
	// lengthOfNumbers := len(numbersToSum)
	// sums := make([]int, lengthOfNumbers)

	// for i, numbers := range numbersToSum {
	// 	sums[i] = Sum(numbers)
	// }

	// way 2 with slice
	var sums []int
	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}

	return sums
}

func SumAllTails(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			tail := numbers[1:]
			sums = append(sums, Sum(tail))
		}
	}
	return sums
}
