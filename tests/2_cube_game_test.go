package tests

import (
	"testing"

	two "github.com/SawyerOG/advent-of-code/2"
)

func TestCubeGame(t *testing.T) {
	got := two.CubeGame("../2/data.txt")
	want := "I am Cube"

	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}
