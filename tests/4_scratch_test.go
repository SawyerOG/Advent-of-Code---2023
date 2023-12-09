package tests

import (
	"testing"

	four "github.com/SawyerOG/advent-of-code/4"
)

// func TestScratch(t *testing.T) {

// 	want := 13
// 	got := four.Scratch("../4/data")

// 	if got != want {
// 		t.Errorf("got %d want %d", got, want)
// 	}
// }

func TestGame(t *testing.T) {
	game := four.GetGames("../4/data_ex1")

	// wants := []int{8, 2, 2, 1, 0, 0}
	wants := map[int]int{
		// 1: 8,
		// 2: 2,
		// 3: 2,
		// 4: 1,
		// 5: 0,
		// 6: 0,
		1: 512,
	}

	for k, g := range *game {
		got := four.DetermineWins(&g)

		if got != wants[k] {
			t.Errorf("game #%d >> got %d want %d", k, got, wants[k])
		}
	}
}

// func TestScore(t *testing.T) {
// 	want := 128 * 2 * 2
// 	got := four.Score(10)

// 	if got != want {
// 		t.Errorf("got %d want %d", got, want)
// 	}
// }
