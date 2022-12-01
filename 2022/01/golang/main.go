package main

import (
	"fmt"
	"sort"
)

type Elves struct {
	number   int
	calories []int
}

type byCalories []Elves

var originalElvesList = []Elves{
	{number: 1, calories: []int{1000, 2000, 3000}},
	{number: 2, calories: []int{4000}},
	{number: 3, calories: []int{5000, 6000}},
	{number: 4, calories: []int{7000, 8000, 9000}},
	{number: 5, calories: []int{10000}},
}

func (e byCalories) Len() int {
	return len(e)
}

func main() {
	allElves := []Elves{}
	for _, elves := range originalElvesList {
		var totalCal int
		for _, calories := range elves.calories {
			totalCal += calories
		}
		allElves = append(allElves, Elves{number: elves.number, calories: []int{totalCal}})
	}

	sort.Slice(allElves, func(i, j int) bool {
		return allElves[i].calories[0] > allElves[j].calories[0]
	})

	fmt.Println(allElves[0])
}
