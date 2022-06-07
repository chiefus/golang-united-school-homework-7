package coverage

import (
	//	"fmt"
	"os"
	"reflect"
	"sort"
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

var baseBirthday = time.Date(2022, time.June, 6, 0, 0, 0, 0, time.UTC)

func TestSortPeople(t *testing.T) {
	testCases := map[string]People{
		"Test sort by birthdays": {
			Person{
				firstName: "Luke",
				lastName:  "Skywalker",
				birthDay:  baseBirthday.AddDate(2, 2, 2),
			},
			Person{
				firstName: "Obi-Wan",
				lastName:  "Kenobi",
				birthDay:  baseBirthday,
			},
			Person{
				firstName: "Darth",
				lastName:  "Vader",
				birthDay:  baseBirthday.AddDate(1, 1, 1),
			},
		},
		"Test sort by first names": {
			Person{
				firstName: "Darth",
				lastName:  "Vader",
				birthDay:  baseBirthday,
			},
			Person{
				firstName: "Obi-Wan",
				lastName:  "Kenobi",
				birthDay:  baseBirthday,
			},
			Person{
				firstName: "Luke",
				lastName:  "Skywalker",
				birthDay:  baseBirthday,
			},
		},
		"Test sort by last names": {
			Person{
				firstName: "Darth",
				lastName:  "Vader",
				birthDay:  baseBirthday,
			},
			Person{
				firstName: "Darth",
				lastName:  "Kenobi",
				birthDay:  baseBirthday,
			},
			Person{
				firstName: "Darth",
				lastName:  "Skywalker",
				birthDay:  baseBirthday,
			},
		},
		"Test sort when all items are equal": {
			Person{
				firstName: "Darth",
				lastName:  "Vader",
				birthDay:  baseBirthday,
			},
			Person{
				firstName: "Darth",
				lastName:  "Vader",
				birthDay:  baseBirthday,
			},
			Person{
				firstName: "Darth",
				lastName:  "Vader",
				birthDay:  baseBirthday,
			},
		},
	}

	expectedValues := map[string]People{
		"Test sort by birthdays": {
			Person{
				firstName: "Luke",
				lastName:  "Skywalker",
				birthDay:  baseBirthday.AddDate(2, 2, 2),
			},
			Person{
				firstName: "Darth",
				lastName:  "Vader",
				birthDay:  baseBirthday.AddDate(1, 1, 1),
			},
			Person{
				firstName: "Obi-Wan",
				lastName:  "Kenobi",
				birthDay:  baseBirthday,
			},
		},
		"Test sort by first names": {
			Person{
				firstName: "Darth",
				lastName:  "Vader",
				birthDay:  baseBirthday,
			},
			Person{
				firstName: "Luke",
				lastName:  "Skywalker",
				birthDay:  baseBirthday,
			},
			Person{
				firstName: "Obi-Wan",
				lastName:  "Kenobi",
				birthDay:  baseBirthday,
			},
		},
		"Test sort by last names": {
			Person{
				firstName: "Darth",
				lastName:  "Kenobi",
				birthDay:  baseBirthday,
			},
			Person{
				firstName: "Darth",
				lastName:  "Skywalker",
				birthDay:  baseBirthday,
			},
			Person{
				firstName: "Darth",
				lastName:  "Vader",
				birthDay:  baseBirthday,
			},
		},
		"Test sort when all items are equal": {
			Person{
				firstName: "Darth",
				lastName:  "Vader",
				birthDay:  baseBirthday,
			},
			Person{
				firstName: "Darth",
				lastName:  "Vader",
				birthDay:  baseBirthday,
			},
			Person{
				firstName: "Darth",
				lastName:  "Vader",
				birthDay:  baseBirthday,
			},
		},
	}

	for testName, testData := range testCases {
		t.Run(testName, func(t *testing.T) {
			expectedValue := expectedValues[testName]
			sort.Stable(testData)
			comparisonResult := reflect.DeepEqual(testData, expectedValue)
			if !comparisonResult {
				t.Errorf("Sort() failed. Exp: %v got: %v", expectedValue, testData)
				// fails on "by birthdays", as it should
				//t.Errorf("Sort() failed. Exp: %v got: %v", expected, people)
			}
		})
	}
}

func TestLenPeople(t *testing.T) {
	testCases := map[string]People{
		"Test Zero Items": {},
		"Test One Item": {
			Person{
				firstName: "Darth",
				lastName:  "Vader",
				birthDay:  baseBirthday,
			},
		},
		"Test Two Items": {
			Person{
				firstName: "Darth",
				lastName:  "Vader",
				birthDay:  baseBirthday,
			},
			Person{
				firstName: "Darth",
				lastName:  "Kenobi",
				birthDay:  baseBirthday,
			},
		},
		"Test Three Items": {
			Person{
				firstName: "Darth",
				lastName:  "Vader",
				birthDay:  baseBirthday,
			},
			Person{
				firstName: "Darth",
				lastName:  "Vader",
				birthDay:  baseBirthday,
			},
			Person{
				firstName: "Darth",
				lastName:  "Vader",
				birthDay:  baseBirthday,
			},
		},
	}

	expectedValues := map[string]int{
		"Test Zero Items":  0,
		"Test One Item":    1,
		"Test Two Items":   2,
		"Test Three Items": 3,
	}

	for testName, testData := range testCases {
		t.Run(testName, func(t *testing.T) {
			expectedValue := expectedValues[testName]
			resultedLength := testData.Len()
			if resultedLength != expectedValue {
				t.Errorf("Len() failed. Exp: %v got: %v", expectedValue, resultedLength)
				// fails on "by birthdays", as it should
				//t.Errorf("Sort() failed. Exp: %v got: %v", expected, people)
			}
		})
	}
}

///////////////////////////////////////////////////////////////////////////////////////////////////////

func TestMatrix_New(t *testing.T) {
	testCases := map[string]struct {
		str      string
		expected *Matrix
		hasError bool
	}{
		"Test of 2x2 matrix creation": {
			`0 1
2 3`,
			&Matrix{2, 2, []int{0, 1, 2, 3}},
			false},
		"Test of 1x1 matrix creation": {
			`0`,
			&Matrix{1, 1, []int{0}},
			false},
		"Test of 1x4 matrix creation": {
			`0 1 2 3`,
			&Matrix{1, 4, []int{0, 1, 2, 3}},
			false},
		"Test of 4x1 matrix creation": {
			`0
1
2
3`,
			&Matrix{4, 1, []int{0, 1, 2, 3}},
			false},
		"Test of bad input: not equal row lengths": {
			`0 1 2 3
4 5`,
			nil,
			true},

		"Test of bad input: character instead digit": {
			`0 1
f 3`,
			nil,
			true},
	}

	for testName, testData := range testCases {
		data := testData
		t.Run(testName, func(t *testing.T) {
			matrix, err := New(data.str)
			if data.hasError {
				if err == nil {
					t.Fatal("The error should have been returned as a test result")
				}
				if matrix != nil {
					t.Fatal("The test should have returned nil matrix")
				}
			} else {
				if err != nil {
					t.Fatal("Error during matrix creation:", err)
				}
				if matrix == nil {
					t.Fatal("Empty matrix has been returned")
				}
				if !reflect.DeepEqual(*matrix, *data.expected) {
					t.Fatalf("Matrix was created incorrectly. Expected: %v got: %v", *data.expected, *matrix)
				}
			}
		})
	}
}

func TestMatrix_Rows(t *testing.T) {
	testCases := map[string]struct {
		str      string
		expected [][]int
	}{
		"Test Rows() of 2x2 matrix": {
			`0 1
2 3`,
			[][]int{{0, 1}, {2, 3}},
		},
		"Test of Rows() 1x1 matrix": {
			`0`,
			[][]int{{0}},
		},
		"Test of of Rows() 1x4 matrix": {
			`0 1 2 3`,
			[][]int{{0, 1, 2, 3}},
		},
		"Test of of Rows()4x1 matrix": {
			`0
1
2
3`,
			[][]int{{0}, {1}, {2}, {3}},
		},
	}
	for testName, testData := range testCases {
		data := testData
		t.Run(testName, func(t *testing.T) {
			matrix, err := New(data.str)
			if err != nil {
				t.Fatal("Error during matrix creation:", err)
			}
			if matrix == nil {
				t.Fatal("Nil matrix has been returned")
			}
			rows := matrix.Rows()
			if rows == nil {
				t.Fatal("Rows() returned nil")
			}
			if !reflect.DeepEqual(rows, data.expected) {
				t.Fatalf("Rows() works incorrectly. Expected: %v got: %v", data.expected, rows)
			}
		})
	}
}

func TestMatrix_Cols(t *testing.T) {
	testCases := map[string]struct {
		str      string
		expected [][]int
	}{
		"Test Cols() of 2x2 matrix": {
			`0 1
2 3`,
			[][]int{{0, 2}, {1, 3}},
		},
		"Test of Cols() 1x1 matrix": {
			`0`,
			[][]int{{0}},
		},
		"Test of of Cols() 1x4 matrix": {
			`0 1 2 3`,
			[][]int{{0}, {1}, {2}, {3}},
		},
		"Test of of Cols()4x1 matrix": {
			`0
1
2
3`,
			[][]int{{0, 1, 2, 3}},
		},
	}
	for testName, testData := range testCases {
		data := testData
		t.Run(testName, func(t *testing.T) {
			matrix, err := New(data.str)
			if err != nil {
				t.Fatal("Error during matrix creation:", err)
			}
			if matrix == nil {
				t.Fatal("Nil matrix has been returned")
			}
			cols := matrix.Cols()
			if cols == nil {
				t.Fatal("Cols() returned nil")
			}
			if !reflect.DeepEqual(cols, data.expected) {
				t.Fatalf("Cols() works incorrectly. Expected: %v got: %v", data.expected, cols)
			}
		})
	}
}

func TestMatrix_Set(t *testing.T) {
	testCases := map[string]struct {
		row, col, value int
		expected        *Matrix
		ok              bool
	}{
		"Test Set() 0 0": {
			0, 0, -1,
			&Matrix{2, 2, []int{-1, 1, 2, 3}},
			true,
		},
		"Test Set() 0 1": {
			0, 1, -1,
			&Matrix{2, 2, []int{0, -1, 2, 3}},
			true,
		},
		"Test Set() 1 0": {
			1, 0, -1,
			&Matrix{2, 2, []int{0, 1, -1, 3}},
			true,
		},
		"Test Set() 1 1": {
			1, 1, -1,
			&Matrix{2, 2, []int{0, 1, 2, -1}},
			true,
		},
		"Test Set(), row number is negative": {
			-1, 0, -1,
			&Matrix{2, 2, []int{0, 1, 2, 3}},
			false,
		},
		"Test Set(), row number is out if index range": {
			2, 0, -1,
			&Matrix{2, 2, []int{0, 1, 2, 3}},
			false,
		},
		"Test Set(), col number is negative": {
			0, -1, -1,
			&Matrix{2, 2, []int{0, 1, 2, 3}},
			false,
		},
		"Test Set(), col number is out if index range": {
			0, 2, -1,
			&Matrix{2, 2, []int{0, 1, 2, 3}},
			false,
		},
	}
	for testName, testData := range testCases {
		data := testData

		t.Run(testName, func(t *testing.T) {
			matrix, err := New(`0 1
2 3`)
			if err != nil {
				t.Fatal("Error during matrix creation:", err)
			}
			ok := matrix.Set(data.row, data.col, data.value)
			if ok != data.ok {
				t.Errorf("Set() works with error. Expected: %v, got: %v", data.ok, ok)
			}
			if !reflect.DeepEqual(*matrix, *data.expected) {
				t.Fatalf("Set() works incorrectly. Expected: %v got: %v", *data.expected, *matrix)
			}
		})
	}
}
