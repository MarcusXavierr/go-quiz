package quiz

import (
	"bufio"
	"fmt"
	"io"
	"strings"
	"time"
)

type event int

const (
	SUCCESS event = iota
	TIMEOUT
)

type Problem struct {
	Question          string
	Solution          string
	AnsweredCorrectly bool
}

func (p *Problem) Answer(solution string) {
	p.AnsweredCorrectly = strings.TrimSpace(solution) == p.Solution
}

type Game struct {
	Problems []*Problem
	Stdout   io.Writer
	Stdin    io.Reader
}

func (g *Game) Play(timeout int) {
	quit := make(chan event)
	scanner := bufio.NewScanner(g.Stdin)

	go func() {
		for i, problem := range g.Problems {
			whitespace := "  "
			if i+1 >= 10 {
				whitespace = " "
			}
			fmt.Fprintf(g.Stdout, "[%d/%d]%s| %s: ", i+1, len(g.Problems), whitespace, problem.Question)
			scanner.Scan()
			text := scanner.Text()
			problem.Answer(text)
		}

		quit <- SUCCESS
	}()

	go func() {
		time.Sleep(time.Duration(timeout) * time.Second)
		quit <- TIMEOUT
	}()

	select {
	case event := <-quit:
		if event == TIMEOUT {
			fmt.Fprint(g.Stdout, "\nThe time is over!\n")
		}

		return
	}
}

func (g *Game) PrintResult() {
	correctAnswers := 0
	for _, problem := range g.Problems {
		if problem.AnsweredCorrectly {
			correctAnswers++
		}
	}

	fmt.Fprintf(g.Stdout, "You scored %d out of %d\n", correctAnswers, len(g.Problems))
}
