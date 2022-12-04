package main

import (
	"fmt"
	"sort"
)

type Student struct {
	Name   string
	Age    int
	Scores float64
}

type StudentSlice []Student

func (stu StudentSlice) Len() int {
	return len(stu)
}

func (stu StudentSlice) Less(i, j int) bool {
	return stu[i].Scores < stu[j].Scores
}

func (stu StudentSlice) Swap(i, j int) {
	stu[i], stu[j] = stu[j], stu[i]
}

func main() {
	var students StudentSlice = []Student{
		{
			Name:   "Vian",
			Age:    17,
			Scores: 98.8,
		},

		{
			Name:   "Niko",
			Age:    18,
			Scores: 87,
		},

		{
			Name:   "Jack",
			Age:    20,
			Scores: 60.5,
		},

		{
			Name:   "Candy",
			Age:    20,
			Scores: 80,
		},

		{
			Name:   "Samity",
			Age:    22,
			Scores: 90,
		},
	}

	for _, v := range students {
		fmt.Println(v)
	}

	fmt.Println()

	sort.Sort(students)
	for _, v := range students {
		fmt.Println(v)
	}
	/*
		{Jack 20 60.5}
		{Candy 20 80}
		{Niko 18 87}
		{Samity 22 90}
		{Vian 17 98.8}
	*/
}
