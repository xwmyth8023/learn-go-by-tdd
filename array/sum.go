package main

import "fmt"

func Sum(numbers []int) int {
	sum := 0
	// for i :=0; i < 5; i++ {
	// 	sum += numbers[i]
	// }
	for _, number := range numbers {
		sum += number
	}
	return sum
}

//切片有容量，容易超出
// func SumAll(numbersToSum ...[]int)(sums []int){
// 	lengthOfNumbers := len(numbersToSum)
// 	// make 创造切片时可指定长度
// 	sums = make([]int, lengthOfNumbers)
// 	for i, numbers := range numbersToSum{
// 		sums[i] = Sum(numbers)
// 	}
// 	return
// }
func SumAll(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}

	return sums
}

func SumAllTails(numbersToSum ...[]int) []int {
	var sums []int
	for _, numers := range numbersToSum {
		//判断数组或者切片为空
		if len(numers) == 0 {
			sums = append(sums, 0)
		} else {
			tail := numers[1:]
			sums = append(sums, Sum(tail))
		}
	}

	return sums
}

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	numbersArries := [][]int{{1, 2}, {3, 4}}
	stringArries := []string{"a", "b", "c"}
	fmt.Println(Sum(numbers))
	fmt.Println(numbersArries)
	fmt.Println(stringArries)
}
