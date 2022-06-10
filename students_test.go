package coverage

import (
	"os"
	"testing"
	"time"
)

// DO NOT EDIT THIS FUNCTION
func init() {
	content, err := os.ReadFile("students_test.go")
	if err != nil {
		panic(err)
	}
	err = os.WriteFile("autocode/students_test", content, 0644)
	if err != nil {
		panic(err)
	}
}

// WRITE YOUR CODE BELOW

func TestLen(t *testing.T) {
	var people = People{
		Person{
			firstName: "A",
			lastName:  "B",
			birthDay:  getTimeWithoutError("12-01-1998"),
		},
		Person{},
		Person{},
	}
	res := people.Len()

	if res != 3 {
		t.Errorf("Result was incorect, got: %d, want: %d", res, 3)
	}
}

func TestLess(t *testing.T) {
	var people = People{
		Person{
			firstName: "A",
			lastName:  "B",
			birthDay:  getTimeWithoutError("1999-02-09"),
		},
		Person{
			firstName: "C",
			lastName:  "D",
			birthDay:  getTimeWithoutError("1999-02-08"),
		},
	}
	res := people.Less(0, 1)

	if res != true {
		t.Errorf("Result was incorect, got: %t, want: %t", res, true)
	}
}

func TestSwap(t *testing.T) {
	var people = People{
		Person{
			firstName: "A",
			lastName:  "B",
			birthDay:  getTimeWithoutError("1998-02-09"),
		},
		Person{
			firstName: "C",
			lastName:  "D",
			birthDay:  getTimeWithoutError("1999-02-08"),
		},
	}

	res := people.Less(0, 1)

	if res != false {
		t.Errorf("Result was incorect, got: %t, want: %t", res, false)
	}

	people.Swap(0, 1)

	res = people.Less(0, 1)

	if res != true {
		t.Errorf("Result was incorect, got: %t, want: %t", res, true)
	}
}

func TestNew(t *testing.T) {
	var res, _ = New("1 2 3\n 2 3 4")

	if res == nil {
		t.Errorf("Result was incorect, got: %v, want: %v", res, nil)
	}
}

func TestNew2(t *testing.T) {
	var res, err = New("")

	if err == nil {
		t.Errorf("Result was incorect, got: %v, want: %v", res, nil)
	}
}

func TestRows(t *testing.T) {
	var res, _ = New("1 2 3\n 2 3 4")

	if res.Rows()[0][2] != 3 {
		t.Errorf("Result was incorect, got: %d, want: %d", res.Rows()[0][2], 3)
	}
}

func TestCols(t *testing.T) {
	var res, _ = New("1 2 3\n 2 3 4 \n 5 6 7")

	if res.Cols()[0][2] != 5 {
		t.Errorf("Result was incorect, got: %d, want: %d", res.Cols()[0][2], 5)
	}
}

func TestSet1(t *testing.T) {
	var m, _ = New("1 2 3\n 2 3 4 \n 5 6 7")
	var res = m.Set(-1, 1, 4)

	if res != false {
		t.Errorf("Result was incorect, got: %t, want: %t", res, false)
	}
}

func TestSet2(t *testing.T) {
	var m, _ = New("1 2 3\n 2 3 4 \n 5 6 7")
	var res = m.Set(2, -1, 4)

	if res != false {
		t.Errorf("Result was incorect, got: %t, want: %t", res, false)
	}
}

func TestSet3(t *testing.T) {
	var m, _ = New("1 2 3\n 2 3 4")
	var res = m.Set(3, 2, 4)

	if res != false {
		t.Errorf("Result was incorect, got: %t, want: %t", res, false)
	}
}

func TestSet4(t *testing.T) {
	var m, _ = New("1 2 3\n 2 3 4 \n 5 6 7")
	var res = m.Set(1, 4, 4)

	if res != false {
		t.Errorf("Result was incorect, got: %t, want: %t", res, false)
	}
}

func TestSet5(t *testing.T) {
	var m, _ = New("1 2 3\n 2 3 4 \n 5 6 7")
	var res = m.Set(2, 2, 4)

	if res != true {
		t.Errorf("Result was incorect, got: %t, want: %t", res, true)
	}
}

func getTimeWithoutError(value string) time.Time {
	timeVar, _ := time.Parse("2006-01-02", value)
	return timeVar
}
