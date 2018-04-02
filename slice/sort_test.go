package slice

import (
	"fmt"
	"testing"
)

func TestSortByKey(t *testing.T) {
	type Student struct {
		Name string
		Age  int
	}

	students := []*Student{
		{
			Name: "lily",
			Age:  10,
		},
		{
			Name: "lucy",
			Age:  20,
		},
	}

	SortAscByKey(students[:], "Name")

	for _, s := range students {
		fmt.Printf("Student = %+v", s)
	}

	SortDescByKey(students[:], "Name")

	for _, s := range students {
		fmt.Printf("Student = %+v", s)
	}
}
