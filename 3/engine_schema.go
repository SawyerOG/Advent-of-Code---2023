package three

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Schema [][]string

/*Helper function to determine if the value is an int*/
func IsInt(s string) (bool, int) {

	v, err := strconv.Atoi(s)
	if err != nil {
		return false, 0
	}

	return true, v
}

func GetData(path string) string {
	file, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err.Error())
	}

	s := string(file)
	return s
}

func GetFlatArray(s string) []string {
	return strings.Split(s, "\n")
}
func GetRows(s string) []string {
	return strings.Split(s, " ")
}

func ConvertToGraph(s []string) *Schema {
	/*
		There are 3 options:
			1. It is a symbol (not a number and not a period)
			2. It is a period (void space)
			3. It is a number (should be grouped until a non numerical char is found)
	*/
	schema := make(Schema, 0)

	for _, v := range s {
		y := strings.Split(v, "")
		schema = append(schema, y)
	}

	return &schema
}

var placesLooked = make(map[string]bool)

func parseNumber(s *Schema, y, x int) int {
	g := (*s)
	fullNum := ""

	// this is where the num was found. go left until it is not a num, then right until it is not a num and
	// keep track of the digits
	startX := x - 1
	leftBoundFound := false

	for leftBoundFound == false {
		if _, seen := placesLooked[strconv.Itoa(y)+","+strconv.Itoa(startX)]; seen {
			return 0
		}
		placesLooked[strconv.Itoa(y)+","+strconv.Itoa(startX)] = true

		if startX < 0 {
			//the number is at the start of the line
			leftBoundFound = true
			startX = 0
			break
		}

		isNum, _ := IsInt(g[y][startX])
		if !isNum {
			//the last value was the first digit
			leftBoundFound = true
			startX += 1
			break
		} else {
			startX--
		}
	}

	for i := startX; i < len(g[y]); i++ {
		isNum, _ := IsInt(g[y][i])
		if isNum {
			fullNum += g[y][i]
		} else {
			break
		}
	}

	_, num := IsInt(fullNum)
	return num
}

var dirs = []struct{ y, x int }{
	//left
	{0, -1},
	//top left
	{-1, -1},
	//top
	{-1, 0},
	//top right
	{-1, 1},
	//right
	{0, 1},
	//bottom right
	{1, 1},
	//bottom
	{1, 0},
	//bottom left
	{1, -1},
}

func findParts(s *Schema, y, x int) int {
	g := (*s)
	partOne := 0
	partTwo := 0

	for _, dir := range dirs {
		if y+dir.y >= 0 && y+dir.y < len(g) && x+dir.x >= 0 && x+dir.x < len(g[y]) {
			isNum, _ := IsInt(g[y+dir.y][x+dir.x])
			if isNum {
				// fmt.Println("num")
				// fmt.Println(g[y+dir.y][x+dir.x])
				val := parseNumber(s, y+dir.y, x+dir.x)
				// fmt.Println("val: ", val)

				if val > 0 {
					//only want those items with 2 parts
					if partOne != 0 && partTwo != 0 {
						//this gear exceeds the criteria (2 parts)
						partOne = 0
						partTwo = 0
						break
					}
					if partOne == 0 {
						partOne = val
					} else {
						partTwo = val
					}
				}
			}
		}
	}

	// if both parts multiply them and return the valu

	// fmt.Printf("partOne: %d, partTwo: %d\n", partOne, partTwo)
	if partOne > 0 && partTwo > 0 {
		return partOne * partTwo
	}

	// bad gear
	return 0
}

func findSymbols(s *Schema) int {
	g := (*s)
	sum := 0

	for y := 0; y < len(g); y++ {
		for x := 0; x < len(g[y]); x++ {

			val := g[y][x]
			// fmt.Println(val)
			// isNum, _ := IsInt(val)

			// if val != "." && isNum == false {
			if val == "*" {
				// fmt.Println("val: ", val)
				foundParts := findParts(s, y, x)
				// fmt.Println("foundParts: ", foundParts)
				sum += foundParts
			}

		}
	}

	return sum
}

func EngineSchema(path string) int {

	fileData := GetData(path)
	flatArray := GetFlatArray(fileData)
	graph := ConvertToGraph(flatArray)

	sum := findSymbols(graph)

	fmt.Println("total: ", sum) //part two :87449461
	return sum                  //546312
}
