package one

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

/*
The newly-improved calibration document consists of lines of text; each line originally contained
a specific calibration value that the Elves now need to recover. On each line, the calibration value can be found by
combining the first digit and the last digit (in that order) to form a single two-digit number.

Each line, so I think that means that this should be read from a text doc and split by each line

2911threeninesdvxvheightwobm >> f: 2, l:1 = 3
*  *
*/

/*read the file at the given path, convert it to a string and hand it back*/
func GetFile(path string) (string, error) {
	file, err := os.ReadFile(path)
	if err != nil {
		return "", err

	}
	return string(file), nil
}

/* Parse the file to into lines on \n */
func ParseLines(content string) (*[]string, error) {
	s := strings.Split(content, "\n")
	return &s, nil
}

var strToNum = map[string]string{
	"zero":  "0",
	"one":   "1",
	"two":   "2",
	"three": "3",
	"four":  "4",
	"five":  "5",
	"six":   "6",
	"seven": "7",
	"eight": "8",
	"nine":  "9",
}

func FindIntInString(s string) (char string, ok bool) {
	/*
		seven5khtwo891hlb
		min: 3
		max: 5
	*/
	// fmt.Println(s)
	// if all 5 letters match it does not matter which dir were heading
	char, ok = strToNum[s]
	if ok {
		// fmt.Println("charLuck >> ", charLuck)
		return char, ok
	}

	char, ok = strToNum[s[0:3]]
	if ok {
		// fmt.Println("char3 >> ", charThree)
		return char, ok
	}

	if len(s) >= 4 {
		char, ok = strToNum[s[0:4]]
		if ok {
			// fmt.Println("char40 >> ", charFour)
			return char, ok
		}
	}

	return "", false
}

/*Find the first and last int in a string and return the concatenated int from those values*/
func FindInt(value string) (f, l string) {

	//assume going from left to right
	for i := 0; i < len(value); i++ {
		v := string(value[i])

		if IsInt(v) {
			if f == "" {
				f = v
				l = v
			} else {
				l = v
			}
			continue
		}

		//else, need to figure out if the number is a string embedded in the line
		var nextFive string
		rem := len(value) - i

		if rem >= 5 {
			nextFive = value[i : i+5]
		} else {
			nextFive = value[i : i+rem]
		}

		if len(nextFive) > 2 {
			//dont check if there is no way to create a word
			digit, ok := FindIntInString(nextFive)
			if ok {
				if f == "" {
					f = digit
					l = digit
				} else {
					l = digit
				}
				continue
			}
		}
	}

	return f, l
}

/*Helper function to determine if the value is an int*/
func IsInt(s string) bool {

	_, err := strconv.Atoi(s)
	if err != nil {
		return false
	}

	return true
}

func Trebuchet(path string) int {
	// data, err := GetFile("./1/data_ex.txt")
	data, err := GetFile(path)
	if err != nil {
		// fmt.Println("GetFile")
		log.Fatal(err)
	}

	lines, err := ParseLines(data)
	if err != nil {
		// fmt.Println("ParseLines")
		log.Fatal(err)
	}

	sum := 0

	for _, v := range *lines {
		f, l := FindInt(v)

		val, err := strconv.Atoi(f + l)
		// fmt.Println(val)
		if err != nil {
			log.Fatal("failed to convert concat int >> " + f + l)
		}

		sum += val
	}

	fmt.Println(sum)
	return sum
}
