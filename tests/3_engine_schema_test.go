package tests

import (
	"fmt"
	"testing"

	three "github.com/SawyerOG/advent-of-code/3"
)

func TestEngineSchema(t *testing.T) {

	var fileData string
	var flatArray []string
	var graph [][]string

	t.Run("Get the file data", func(t *testing.T) {
		fileData = three.GetData("../3/data_ex.txt")

		want := `467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..`

		if fileData != want {
			t.Errorf("got %s want %s", fileData, want)
		}

	})

	t.Run("Convert File Data to Flat Array", func(t *testing.T) {
		flatArray = three.GetFlatArray(fileData)
		wantLength := 10

		fmt.Println("FLAT >> ", flatArray)

		l := len(flatArray)

		if l != wantLength {
			t.Errorf("want length %d got length %d", wantLength, l)
		}
	})

	t.Run("Get Graph", func(t *testing.T) {
		graph = three.ConvertToGraph(flatArray)

		// fmt.Println("GRAPH >> ", got)

		expectLenY := 10

		if len(graph) != expectLenY {
			t.Errorf("Graph len want %d, got len %d", expectLenY, len(graph))
		}

		expectFirstNumIdx := 2
		expectFirstNum := "467"
		expectFirstRowLen := expectLenY

		if len(graph[0]) != expectFirstRowLen {
			t.Errorf("Graph len want %d, got len %d", expectLenY, len(graph))
		}

		if graph[0][expectFirstNumIdx] != expectFirstNum {
			t.Errorf("didnt find the number >> want %s, got %s", expectFirstNum, graph[0][expectFirstNumIdx])
		}

	})

	t.Run("Find the Symbols", func(t *testing.T) {
		found := three.LookAround(0, 0, 3, len(graph), graph)
		want := true

		if found != want {
			t.Errorf("parts not found >> want %v got %v", want, found)
		}

		found = three.LookAround(0, 5, 3, len(graph), graph)
		want = false

		if found != want {
			t.Errorf("parts not found >> want %v got %v", want, found)
		}
	})

	t.Run("Find Engine Parts", func(t *testing.T) {
		got := three.CheckRow(0, graph)
		want := 467

		if got != want {
			t.Errorf("ROW 0 got %d want %d", got, want)
		}

		got = three.CheckRow(1, graph)
		want = 0

		if got != want {
			t.Errorf("ROW 1 got %d want %d", got, want)
		}

		got = three.CheckRow(2, graph)
		want = 35 + 633

		if got != want {
			t.Errorf("ROW 2 got %d want %d", got, want)
		}

		got = three.CheckRow(4, graph)
		want = 617

		if got != want {
			t.Errorf("ROW 4 got %d want %d", got, want)
		}
	})

}

func TestEngineSchemaWhole(t *testing.T) {
	got := three.EngineSchema("../3/data_ex.txt")
	want := 4361

	if got != want {
		t.Errorf("Whole thing got %d want %d", got, want)
	}

}
