package main

import "fmt"

func main() {
	questionsconditionalStatement := conditionalStatementGenerate(100)
	questionsTruthTable := truthTableGenerate(0)
	questionsDeMorgan := deMorganGenerate(0)
	

	for _, q := range questionsconditionalStatement {
		fmt.Printf("Question: %s\n", q["question"])
		fmt.Printf("Correct Answer: %s\n", q["correct_answer"])
		wrongAnswers := q["wrong_answers"].([]string)
		fmt.Printf("Wrong Answers: %s, %s, %s\n\n", wrongAnswers[0], wrongAnswers[1], wrongAnswers[2])
	}

	for _, q := range questionsTruthTable {
		fmt.Printf("Question: %s\n", q["question"])
		fmt.Printf("Correct Answer: %t\n", q["correct_answer"])
		fmt.Printf("Incorrect Answer: %t\n\n", q["incorrect_answer"])
	}

	for _, q := range questionsDeMorgan {
		fmt.Printf("Question: %s\n", q["question"])
		fmt.Printf("Correct Answer: %s\n", q["answer"])
		wrongAnswers := q["wrong_answers"].([]string)
		fmt.Printf("Wrong Answers: %s, %s, %s\n\n", wrongAnswers[0], wrongAnswers[1], wrongAnswers[2])
	}

}