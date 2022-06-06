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
