package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
)

type questions struct {
	q string
	a string
}

func main() {

	// accepting csv file name from the user. Default is set to questions.csv
	csvFileName := flag.String("csv", "questions.csv", "a CSV file in the format of 'question,answer'")

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

			totalQuestions := len(qa)
			correctAnswers := 0
			fmt.Println("Total Questions:", totalQuestions)
			fmt.Println("Your quiz starts now")
			var inputAnswer string
			currentQuestion := 1
			for q, a := range qa {
				fmt.Println("Q" + strconv.Itoa(currentQuestion) + ":")
				fmt.Println(q)
				fmt.Print("Please enter your answer: ")
				fmt.Scan(&inputAnswer)

				if inputAnswer == a {
					correctAnswers = correctAnswers + 1
				}
				currentQuestion = currentQuestion + 1

			}

			fmt.Printf("Score: %v/%v", correctAnswers, totalQuestions)
		}
	}

}
