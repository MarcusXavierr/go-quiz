package csvparser

import (
	"reflect"
	"testing"
	"testing/fstest"

	"github.com/MarcusXavierr/go-quiz/internal/quiz"
)

func TestReadCsv(t *testing.T) {
	t.Run("Validate successful file parsing", func(t *testing.T) {
		fs := fstest.MapFS{
			"test.txt": {Data: []byte("1+1,2\n\"What is 2+2, sir?\", 4")},
		}

		problem := ProblemsFile{Path: "test.txt", FileSystem: fs}
		got, err := problem.ParseProblems(false)
		if err != nil {
			t.Errorf("Not expected error:\n%+v", err)
		}

		want := []*quiz.Problem{
			{Question: "1+1", Solution: "2"},
			{Question: "What is 2+2, sir?", Solution: "4"},
		}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("expected %v but got %v", want, got)
		}
	})

	t.Run("Raise an erro on invalid csv file", func(t *testing.T) {
		fs := fstest.MapFS{
			"test.txt": {Data: []byte("invalid csv")},
		}

		problem := ProblemsFile{Path: "test.txt", FileSystem: fs}
		_, err := problem.ParseProblems(false)
		if err == nil {
			t.Error("Expected error here")
		}
	})
}
