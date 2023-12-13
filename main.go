package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	csvFilename := flag.String("csv", "problems.csv", "a csv file in the format of 'questions,answer'")
	timeLimit := flag.Int("limit", 3, "time limit for the quiz in seconds")
	flag.Parse()

	file, err := os.Open(*csvFilename)
	if err != nil {
		exit(fmt.Sprintf("failed to open csv file: %s", *csvFilename))
	}
	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		exit("failed to parse the csv file")
	}
	problems := parseLines(lines)

	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)
	// <-timer.C

	correct := 0
problemloop:
	for i, p := range problems {
		fmt.Printf("problem #%d:%s= \n", i+1, p.q)

		answerCh := make(chan string)
		go func() {
			var answer string
			fmt.Scanf("%s\n", &answer)
			answerCh <- answer
		}()

		select {
		case <-timer.C:
			fmt.Println()
			break problemloop
		case answer:= <-answerCh:
			if answer == p.a {
				correct++
			}
		}

	}
	fmt.Printf("\nyou scored %d out of %d.\n", correct, len(problems))
}

type problem struct {
	q string
	a string
}

func parseLines(lines [][]string) []problem {
	ret := make([]problem, len(lines))
	for i, line := range lines {
		ret[i] = problem{
			q: line[0],
			a: strings.TrimSpace(line[1]),
		}
	}
	return ret
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}