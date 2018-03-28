package slice

import (
	"fmt"
	"testing"

	"github.com/xiagoo/gosort/consts"
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

	SortByKey(students[:], "Name", consts.Asc)

	for _, s := range students {
		fmt.Printf("Student = %+v", s)
	}

}
