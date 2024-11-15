package quiz

import (
	"bytes"
	"strings"
	"testing"
)

func TestProblem(t *testing.T) {
	problem := Problem{Question: "blabla", Solution: "abc"}
	t.Run("Answer problem correctly", func(t *testing.T) {
		problem.Answer("abc")
		got := problem.AnsweredCorrectly
		want := true

		if got != want {
			t.Errorf("expected %t but got %t", want, got)
		}
	})

	t.Run("Answer problem correctly", func(t *testing.T) {
		problem.Answer("bac")
		got := problem.AnsweredCorrectly
		want := false

		if got != want {
			t.Errorf("expected %t but got %t", want, got)
		}
	})
}

func TestGame(t *testing.T) {
	t.Run("Play game", func(t *testing.T) {
		stdout := &bytes.Buffer{}
		stdin := strings.NewReader("123\n456")
		game := Game{
			Problems: []*Problem{{Question: "abc", Solution: "123"}, {Question: "def", Solution: "456"}},
			Stdin:    stdin,
			Stdout:   stdout,
		}

		game.Play(10)
		for _, problem := range game.Problems {
			if problem.AnsweredCorrectly != true {
				t.Errorf("Problem %s should be correctly answered with %s", problem.Question, problem.Solution)
			}
		}
	})

	t.Run("Display actual result to user", func(t *testing.T) {
		stdout := &bytes.Buffer{}
		stdin := strings.NewReader("")
		game := Game{
			Problems: []*Problem{{"abc", "123", true}, {"def", "456", false}},
			Stdin:    stdin,
			Stdout:   stdout,
		}

		game.PrintResult()
		got := stdout.String()
		want := "You scored 1 out of 2\n"

		if got != want {
			t.Errorf("expected %s but got %s", want, got)
		}
	})
}
