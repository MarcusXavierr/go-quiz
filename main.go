package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/MarcusXavierr/go-quiz/internal/csvparser"
	"github.com/MarcusXavierr/go-quiz/internal/quiz"
)

func main() {
	path := flag.String("csv", "problems.csv", "A csv file in the format of 'question,anwser' (default \"problems.csv\")")
	timeout := flag.Int("timeout", 30, "The timeout value")
	shuffle := flag.Bool("shuffle", false, "Shuffle all the problems")

	flag.Parse()

	problems, err := csvparser.ProblemsFile{Path: *path, FileSystem: os.DirFS(".")}.ParseProblems(*shuffle)
	if err != nil {
		fmt.Printf("Error starting application!\n\n%+v", err)
		return
	}

	game := quiz.Game{Problems: problems, Stdout: os.Stdout, Stdin: os.Stdin}
	game.Play(*timeout)
	game.PrintResult()
}
