package arrays

func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SumAll(numbersToSum ...[]int) (totalSum []int) {
	// sliceLength := len(numbersToSum)
	// totalSum := make([]int, sliceLength)

	// for i, numbers := range numbersToSum {
	// 	totalSum[i] = Sum(numbers)
	// }
	// return totalSum
	for _, numbers := range numbersToSum {
		totalSum = append(totalSum, Sum(numbers))
	}

	return
}

func SumAllTails(numbersToSum ...[]int) (totalSum []int) {
	for _, numbers := range numbersToSum {
		if len(numbers) == 0 {
			totalSum = append(totalSum, 0)
		} else {
			tail := numbers[1:]
			totalSum = append(totalSum, Sum(tail))
		}
	}

	return
}
