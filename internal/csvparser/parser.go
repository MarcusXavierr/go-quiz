package csvparser

import (
	"encoding/csv"
	"fmt"
	"io/fs"
	"strings"

	"math/rand"

	"github.com/MarcusXavierr/go-quiz/internal/quiz"
	"github.com/pkg/errors"
)

type ProblemsFile struct {
	Path       string
	FileSystem fs.FS
}

func (p ProblemsFile) ParseProblems(suffle bool) ([]*quiz.Problem, error) {
	file, err := p.FileSystem.Open(p.Path)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to Open File")
	}

	return parseCsv(file, suffle)
}

func parseCsv(file fs.File, shuffle bool) ([]*quiz.Problem, error) {
	records, err := csv.NewReader(file).ReadAll()
	if err != nil {
		return nil, errors.Wrap(err, "Error reading csv file")
	}

	problems := make([]*quiz.Problem, len(records))

	for k, record := range records {
		if len(record) < 2 {
			return nil, errors.New(fmt.Sprintf("The row %d in the CSV file does not contain all columns\n", k))
		}

		problem := quiz.Problem{
			Question: strings.TrimSpace(record[0]),
			Solution: strings.TrimSpace(record[1]),
		}
		problems[k] = &problem
	}

	if shuffle {
		arr := problems
		rand.Shuffle(
			len(arr),
			func(i, j int) {
				arr[i], arr[j] = arr[j], arr[i]
			},
		)
	}

	return problems, nil
}
