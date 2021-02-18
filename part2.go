package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

type questions struct {
	q string
	a string
}

var totalQuestions, correctAnswers int

func main() {

	// accepting csv file name from the user. Default is set to questions.csv
	csvFileName := flag.String("csv", "questions.csv", "a CSV file in the format of 'question,answer'")

	fmt.Printf("%T", csvFileName)

	timeLimit := flag.Int("limit", 30, "Time limit for the quiz in seconds")

	flag.Parse()

	_ = csvFileName

	// opening CSV file

	file, err := os.Open(*csvFileName)

	if err != nil {
		log.Fatal("Failed to open CSV file %s", *csvFileName)
	} else {
		data := csv.NewReader(file)
		lines, err := data.ReadAll()
		if err != nil {
			log.Fatal("Failed to read data from CSV file: ", *csvFileName)

		} else {
			qa := map[string]string{}
			for i, val := range lines {
				fmt.Println(i)
				qa[val[0]] = val[1]
			}

			// total number of questions

			totalQuestions = len(qa)
			correctAnswers = 0
			fmt.Println("Total Questions:", totalQuestions)
			fmt.Println("Your quiz starts now")
			var inputAnswer string
			currentQuestion := 1
			timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)

			for q, a := range qa {
				fmt.Println("Q" + strconv.Itoa(currentQuestion) + ":")
				fmt.Println(q)
				answeCh := make(chan string)
				go func() {
					fmt.Scan(&inputAnswer)
					answeCh <- inputAnswer
				}()

				select {
				case <-timer.C:
					fmt.Printf("Score: %v/%v", correctAnswers, totalQuestions)
					return

				case answer := <-answeCh:
					if answer == a {
						correctAnswers = correctAnswers + 1
					}
					currentQuestion = currentQuestion + 1

				}

			}

		}
	}

}
