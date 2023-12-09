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

func DetermineWins(game [][]string) int {
	ourNums := make(map[string]bool)
	foundNums := 0
	// fmt.Println("game", game)

	for _, num := range (game)[0] {
		ourNums[num] = true
	}
	// fmt.Println("ourNums", ourNums)

	for _, num := range (game)[1] {
		if num != "" {

			if _, ok := ourNums[num]; ok {
				// fmt.Println("found num", num)
				foundNums++

				//probably only want to count the value once. In case the value is in
				//the possible winners > 1 remove it from our pool
				delete(ourNums, num)
			}
		}

		if len(ourNums) == 0 {
			break
		}
	}

	if foundNums == 0 {
		return 0
	}

	fmt.Println("foundNums", foundNums)
	// return Score(foundNums) //part one scoring
	return foundNums //part two scoring
}

func ScoreTwo(gameNum int, game *[][]string) int {
	//for each win, games[x + 1 + win]) are played again

	fmt.Println("gameNum", gameNum)
	// fmt.Println("game", game)
	return 0
}

func Scratch(path string) int {
	sum := 0
	game := GetGames(path)
	gameRef := *game

	gameScores := make(map[int]int)
	timesGamePlayed := make(map[int]int)

	// for k, g := range *game {
	for k := 1; k <= len(gameRef); k++ {

		wins := DetermineWins(gameRef[k])
		//track the score for this game
		gameScores[k] = wins
		timesGamePlayed[k] = 1
	}

	for k := 1; k <= len(gameRef); k++ {
		// for game, score := range gameScores {
		//if on game 10 and get 6 wins that means we get a copy of game 11 - 17
		//i think as long as those games exist
		game := k
		score := gameScores[k]

		start := game + 1
		end := game + score

		for i := start; i <= end; i++ {
			if _, ok := timesGamePlayed[i]; ok {
				for j := 0; j < timesGamePlayed[game]; j++ {
					timesGamePlayed[i]++
				}
			}
		}

	}

	for _, timesPlayed := range timesGamePlayed {
		sum += timesPlayed
	}

	fmt.Println("sum", sum)
	// fmt.Println("gameScores", gameScores)
	// fmt.Println("timesGamePlayed", timesGamePlayed)

	return sum //26346 //8467762
}
