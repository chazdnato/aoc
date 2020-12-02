package day05

import (
	"os"
	"testing"

	"github.com/chazdnato/aoc/sidwtrw"
)

func BenchmarkReactPolymer(b *testing.B) {
	polymer := sidwtrw.SliceOfStrings("input.txt")[0]

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		reactPolymer(polymer)
	}
}

func TestSolutionOne(t *testing.T) {
	_ = os.Chdir("..")
	answer := SolutionOne()
	if answer != 11476 {
		t.Error("Expected: 11476")
	}
	_ = os.Chdir("day05")
}

func TestSolutionTwo(t *testing.T) {
	_ = os.Chdir("..")
	answer := SolutionTwo()
	if answer != 5446 {
		t.Error("Expected: 5446")
	}
	_ = os.Chdir("day05")
}
