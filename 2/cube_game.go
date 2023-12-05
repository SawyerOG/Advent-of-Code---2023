package two

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

//12 red cubes, 13 green cubes, and 14 blue cubes

type Game map[int]map[string]int

func GenGames(path string) Game {
	file, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err.Error())
	}

	s := strings.Split(string(file), "\n")
	// fmt.Println(s)

	game := Game{}

	for k, v := range s {

		semi := strings.Split(v, ":")[1]
		// fmt.Println("semi >> ", semi)

		list := strings.Split(semi, ";") //[ 3 blue, 4 red  1 red, 2 green, 6 blue  2 green]
		// fmt.Println("LIST >> ", list)

		pairs := make([]string, 0)
		for _, l := range list {
			pair := strings.Split(l, ",")
			for _, ll := range pair {
				pairs = append(pairs, strings.TrimSpace(ll))
			}
		}

		gameID := k + 1
		game[gameID] = make(map[string]int)

		for _, p := range pairs {

			i := strings.Split(p, " ") // [3, blue]

			color := i[1]
			count, err := strconv.Atoi(i[0])
			if err != nil {
				log.Fatal("could not convert to int", i)
			}

			//has this color been seen?
			if colors, ok := game[gameID]; ok {
				// yes - is this value larger
				if count > colors[color] {
					//keep track of this largest value
					colors[color] = count
				}
			} else {
				//no - this is the largest count!
				colors[color] = count
			}
		}

		// fmt.Println(game)
	}

	return game
}

// 12 red cubes, 13 green cubes, and 14 blue
var maxes = map[string]int{
	"red":   12,
	"green": 13,
	"blue":  14,
}

func CubeGame(path string) string {
	sum := 0
	games := GenGames(path)

	for gameID, game := range games {

		winner := true

		for color, count := range game {

			//if the color does not exits it loses
			if max, ok := maxes[color]; !ok {
				//do nothing
				continue
			} else {
				//if the count is less than the max this game is in the running to be a winner
				if count > max {
					winner = false
				}
			}
		}

		if winner {
			sum += gameID
		}
	}

	fmt.Println(sum)

	return "I am Cube"
}
