package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

type Elves struct {
	number   int
	calories []int
}

type byCalories []Elves

var (
	originalElvesList = []Elves{
		{number: 1, calories: []int{1000, 2000, 3000}},
		{number: 2, calories: []int{4000}},
		{number: 3, calories: []int{5000, 6000}},
		{number: 4, calories: []int{7000, 8000, 9000}},
		{number: 5, calories: []int{10000}},
	}
)

func (e byCalories) Len() int {
	return len(e)
}

func main() {
	input, err := ioutil.ReadFile("input.txt")
	if err != nil {
		fmt.Println(err)
	}

	inputString := string(input)
	inputSplit := strings.Split((strings.Replace(inputString, "\n", ",", -1)), ",,")

	allElves := []Elves{}

	for elves, calories := range inputSplit {
		var totalCal int
		for _, calories := range strings.Split(calories, ",") {
			cal, _ := strconv.Atoi(calories)
			totalCal += cal
		}
		allElves = append(allElves, Elves{number: elves, calories: []int{totalCal}})
	}

	sort.Slice(allElves, func(i, j int) bool {
		return allElves[i].calories[0] > allElves[j].calories[0]
	})

	fmt.Println(allElves[0])
}
