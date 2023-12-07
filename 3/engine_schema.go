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

func ConvertToGraph(s []string) Schema {
	/*
		There are 3 options:
			1. It is a symbol (not a number and not a period)
			2. It is a period (void space)
			3. It is a number (should be grouped until a non numerical char is found)
	*/
	schema := make(Schema, 0)
	row := make([]string, 0)

	lastFullNumber := ""

	for i := 0; i < len(s); i++ {

		line := strings.Split(s[i], "")

		for v := 0; v < len(line); v++ {
			val := string(line[v])

			isNum, _ := IsInt(val)
			if isNum {
				lastFullNumber += val
				row = append(row, ".")
			} else {
				if lastFullNumber != "" {

					for n := len(lastFullNumber); n > 0; n-- {
						//put the number in every slot the number took
						row[v-n] = lastFullNumber
					}

					lastFullNumber = ""
				}

				if val == "." {
					row = append(row, val)
				} else {
					//symbol
					row = append(row, "#")
				}

			}
		}
		schema = append(schema, row)
		row = nil
	}

	return schema

}

func LookAround(y, x, count, lenY int, graph Schema) bool {
	foundSymbol := false

	//up
	for i := -1; i <= count+2; i++ {
		xloc := x + i
		if y-1 >= 0 && xloc >= 0 && xloc < lenY {
			if graph[y-1][xloc] == "#" {
				foundSymbol = true
				break
			}
		}
	}

	if !foundSymbol {
		//down
		for i := -1; i <= count+2; i++ {
			xloc := x + i
			if y+1 >= 0 && xloc >= 0 && xloc < lenY {
				if graph[y+1][xloc] == "#" {
					foundSymbol = true
					break
				}
			}
		}
	}

	//left
	if !foundSymbol {
		if x-1 >= 0 && graph[y][x-1] == "#" {
			foundSymbol = true
		}
	}

	//right
	if !foundSymbol {
		if x+1 <= len(graph[y]) && graph[y][x+1] == "#" {
			foundSymbol = true
		}
	}

	return foundSymbol
}

func CheckRow(y int, graph Schema) int {
	sum := 0
	fmt.Println(graph[y])
	lenY := len(graph[y])
	for x := 0; x < lenY-1; {
		isNum, num := IsInt(graph[y][x])
		if isNum {

			fmt.Println("is num >> ", num)

			lengthOfNum := len(graph[y][x])
			//find the symbol
			isPart := LookAround(y, x, lengthOfNum, lenY, graph)

			if isPart {
				fmt.Println("is part!", num)
				sum += num

				//skip ahead past the num
				x = x + lengthOfNum
			} else {
				x = x + 1
			}

		} else {
			x = x + 1
		}
	}

	return sum
}

func EngineSchema(path string) int {

	fileData := GetData(path)
	flatArray := GetFlatArray(fileData)
	graph := ConvertToGraph(flatArray)

	sum := 0
	for y := 0; y < len(graph); y++ {
		s := CheckRow(y, graph)
		sum += s
	}

	fmt.Println(sum)
	return sum
}
