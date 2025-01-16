package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"time"
)

type problem struct {
	question string
	answer   string
}

func main() {
	var filename = flag.String("quiz", "problems.csv", "quiz file in format 'question,answer'")
	var timerLimit = flag.Int("timer", 10, "num of seconds to complete the quiz")

	flag.Parse()

	timer := time.NewTimer(time.Duration(*timerLimit) * time.Second)

	lines := getCsv(*filename)
	problems := parseCsv(lines)

	fmt.Println(problems)

	var correct int
	answers := make(chan string)

	for _, line := range lines {
		go func(ch chan string) {
			var answer string
			fmt.Scanf("%s", &answer)
			answers <- answer
		}(answers)

		fmt.Println(line[0], "?")

		select {
		case <-timer.C:
			fmt.Println("You scored %d out of %d", correct, len(problems))
			return
		case answer := <-answers:
			if answer == line[1] {
				correct++
			}
		}
	}

	fmt.Println("You scored %d out of %d", correct, len(problems))
}

func getCsv(filename string) [][]string {
	ioReader, err := os.OpenFile(filename, os.O_RDONLY, 0644)

	if err != nil {
		panic(err)
	}

	defer ioReader.Close()

	csvReader := csv.NewReader(ioReader)
	lines, err := csvReader.ReadAll()

	if err != nil {
		fmt.Println("Error while parsing CSV:", err)
		os.Exit(1)
	}

	return lines
}

func parseCsv(lines [][]string) []problem {
	r := make([]problem, len(lines))
	for i, l := range lines {
		r[i] = problem{l[0], l[1]}
	}
	return r
}
