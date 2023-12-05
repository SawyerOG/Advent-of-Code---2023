package tests

import (
	"testing"

	one "github.com/SawyerOG/advent-of-code/1"
)

func check(got, want bool, t *testing.T) {
	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}
func TestGetInt(t *testing.T) {
	check(one.IsInt("1"), true, t)
	check(one.IsInt("w"), false, t)
	check(one.IsInt("e"), false, t)
	check(one.IsInt("3"), true, t)

	//fail
	// check(one.IsInt("3"), false, t)
}

func TestTrebuchet(t *testing.T) {

	//first test the read file function

	var fileContent = ""
	t.Run("Open the file", func(t *testing.T) {
		fileContent = `2911threeninesdvxvheightwobm
3three16xsxhpnqmzmnine8one
seven5khtwo891hlb
sixthreeqpzjpn195
jrnf3`
		//make the text above have no tabs for this test to pass

		got, err := one.GetFile("../1/data_test.txt")
		if err != nil {
			t.Fatal(err.Error())
		}

		if got != fileContent {
			t.Errorf("got %q want %q", got, fileContent)
			t.FailNow()
		}

		//Parse the file to into lines on \n
		t.Run("Parse into line", func(t *testing.T) {
			//I expect a pointer to a []string
			wantLength := 5
			s, err := one.ParseLines(fileContent)

			if err != nil {
				t.Fatal(err.Error())
			}
			sVal := (*s)

			l := len(sVal)

			if wantLength != l {
				t.Errorf("got length %d want length %d", wantLength, l)
			}

			wantFirstVal := "2911threeninesdvxvheightwobm"
			gotFirstVal := sVal[0]

			if wantFirstVal != gotFirstVal {
				t.Errorf("want %s got %s", wantFirstVal, gotFirstVal)
				t.FailNow()
			}

		})
	})

	t.Run("Find an int in the string", func(t *testing.T) {

		s := "2911threeninesdvxvheightwobm"

		wantF := "2"
		wantB := "2"
		gotF, gotB := one.FindInt(s)

		if wantF != gotF {
			t.Errorf("failed forward want %s got %s", wantF, gotF)
		}

		if wantB != gotB {
			t.Errorf("failed backward want %s got %s", wantB, gotB)
		}

		wantFull := "21"
		gotFull := gotF + gotB

		if wantB != gotB {
			t.Errorf("failed concat want %s got %s", wantFull, gotFull)
		}
	})
	//then test the trebuchet function

	t.Run("Trebuchet", func(t *testing.T) {
		got := one.Trebuchet("../1/data.txt")
		want := 53389

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}

	})

}

func BenchmarkTrebuchet(b *testing.B) {
	one.Trebuchet("../1/data.txt")
}
