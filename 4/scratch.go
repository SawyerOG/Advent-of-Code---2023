package four

import (
	"fmt"
	"os"
	"strings"
)

type Games map[int][][]string

func GetGames(path string) *Games {
	file, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	games := make(Games)
	s := string(file)
	lines := strings.Split(s, "\n")
	for k, line := range lines {

		game := make([][]string, 2)

		cols := strings.Split(line, ": ")
		nums := strings.Split(cols[1], "|")

		winningNums := strings.TrimSpace(nums[0])
		ourNums := strings.TrimSpace(nums[1])

		game[0] = strings.Split(ourNums, " ")
		game[1] = strings.Split(winningNums, " ")

		games[k+1] = game
	}

	fmt.Println("games", games)
	return &games
}

var scores = map[int]int{1: 1, 2: 2, 3: 4, 4: 8, 5: 16, 6: 32}

func Score(val int) int {

	if s, ok := scores[val]; ok {
		return s
	}

	sum := 1
	for i := 1; i < val; i++ {
		sum *= 2
	}
	scores[val] = sum
	fmt.Printf("val %d sum %d\n", val, sum)
	return sum
}

func DetermineWins(game *[][]string) int {
	ourNums := make(map[string]bool)
	foundNums := 0
	// fmt.Println("game", game)

	for _, num := range (*game)[0] {
		ourNums[num] = true
	}
	// fmt.Println("ourNums", ourNums)

	for _, num := range (*game)[1] {
		if num != "" {

			if _, ok := ourNums[num]; ok {
				// fmt.Println("found num", num)
				foundNums++

				//probably only want to count the value once. In case the value is in
				//the possible winners > 1 remove it from our pool
				delete(ourNums, num)
			}
		}
	}

	if foundNums == 0 {
		return 0
	}

	fmt.Println("foundNums", foundNums)
	return Score(foundNums)
}

func Scratch(path string) int {
	sum := 0
	game := GetGames(path)
	for _, g := range *game {
		sum += DetermineWins(&g)
	}

	fmt.Println("sum", sum)

	return sum //26346
}
