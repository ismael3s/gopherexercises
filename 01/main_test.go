package main

import (
	"os"
	"reflect"
	"testing"
)

func TestParseLines(t *testing.T) {
	t.Run("should be able to parse lines and create a slice of problems", func(t *testing.T) {
		test := struct {
			input  [][]string
			expect []problem
		}{
			input:  [][]string{{"1+1", "2"}},
			expect: []problem{{question: "1+1", answer: "2"}},
		}
		got := parseLines(test.input)

		if !reflect.DeepEqual(got, test.expect) {
			t.Errorf("Got: %s\nExpect: %s\n", got, test.expect)
		}
	})

}

func TestOpenFile(t *testing.T) {
	t.Run("should be able to read a csv and return their lines", func(t *testing.T) {
		filename := "./assets/teste.csv"
		f, _ := os.Create(filename)

		f.WriteString("1+1,2\n2+2,4\n")

		f, _ = os.Open(filename)
		defer f.Close()
		defer os.Remove(filename)

		expect := [][]string{{"1+1", "2"}, {"2+2", "4"}}

		got := readCsv(f)

		if !reflect.DeepEqual(got, expect) {
			t.Errorf("Got: %s\nExpect: %s\n", got, expect)
		}

	})
}
