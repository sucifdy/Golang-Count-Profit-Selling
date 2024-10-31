package main

import "fmt"

func CountProfit(data [][][2]int) []int {

	if len(data) == 0 {
		return []int{}
	}

	numMonths := len(data[0])
	profits := make([]int, numMonths)

	for _, branch := range data {

		for month := 0; month < numMonths; month++ {
			sales := branch[month][0]
			expenses := branch[month][1]
			profit := sales - expenses
			profits[month] += profit
		}
	}

	return profits
}

// Test the function with provided test cases
func main() {
	testCases := [][][][2]int{
		{{{1000, 800}, {700, 500}}, {{1000, 800}, {900, 200}}},
		{{{1000, 500}, {500, 150}, {600, 100}, {800, 750}}},
		{{{1000, 200}}, {{500, 100}}, {{600, 100}}, {{450, 150}}, {{100, 50}}},
		{},
	}

	for _, testCase := range testCases {
		result := CountProfit(testCase)
		fmt.Println(result)
	}
}
