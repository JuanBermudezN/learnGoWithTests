package main

func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SumAll(numbersToSum ...[]int) []int {
	lengthOfNumbers := len(numbersToSum)
	sums := make([]int, lengthOfNumbers)

	for i, numbers := range numbersToSum {
		sums[i] = Sum(numbers)
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

// There's a new way to create a slice. make allows you to create a slice with a starting capacity of the len of the numbersToSum we need to work through.

// As mentioned, slices have a capacity. If you have a slice with a capacity of 2 and try to do mySlice[10] = 1 you will get a runtime error.
// However, you can use the append function which takes a slice and a new value, then returns a new slice with all the items in it.

// Slices can be sliced! The syntax is slice[low:high]. If you omit the value on one of the sides of the : it captures everything to that side of it.
// In our case, we are saying "take from 1 to the end" with numbers[1:]. You may wish to spend some time writing other tests around slices and experiment with
// the slice operator to get more familiar with it.
