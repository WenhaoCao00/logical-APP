package main

import (
	"fmt"
	"math/rand"
	"time"
)



func contrapositive(expr string) string {
	return "¬" + expr + "→¬" + expr
}


func inverse(expr string) string {
	return "¬" + expr + "→¬" + expr
}


func converse(expr string) string {
	return expr + "→" + expr
}


func generateRandomExpression(ran *rand.Rand, variables []string) string {
	ops := []string{"∧", "∨", "→"}
	numVars := len(variables)
	expr := variables[ran.Intn(numVars)]
	for i := 0; i < ran.Intn(3)+1; i++ { 
		op := ops[ran.Intn(len(ops))]
		variable := variables[ran.Intn(numVars)]
		expr = fmt.Sprintf("(%s %s %s)", expr, op, variable)
	}
	return expr
}


func generateConditionalExpressionQuestion(ran *rand.Rand) (string, string, []string) {
	variables := []string{"p", "q", "r", "s"}
	expr := generateRandomExpression(ran, variables)


	questionTypes := []struct {
		questionFormat string
		answerFunc     func(string) string
	}{
		{"What is the contrapositive of %s?", contrapositive},
		{"What is the inverse of %s?", inverse},
		{"What is the converse of %s?", converse},
	}

	questionTypeIndex := ran.Intn(len(questionTypes))
	questionType := questionTypes[questionTypeIndex]

	question := fmt.Sprintf(questionType.questionFormat, expr)
	correctAnswer := questionType.answerFunc(expr)


	wrongAnswers := make(map[string]struct{})
	wrongAnswers[correctAnswer] = struct{}{} 

	for len(wrongAnswers) < 4 {
		wrongExpr := generateRandomExpression(ran, variables)
		wrongAnswer := questionType.answerFunc(wrongExpr)
		wrongAnswers[wrongAnswer] = struct{}{}
	}


	delete(wrongAnswers, correctAnswer)
	wrongAnswersList := make([]string, 0, 3)
	for ans := range wrongAnswers {
		wrongAnswersList = append(wrongAnswersList, ans)
	}

	return question, correctAnswer, wrongAnswersList
}

func conditionalStatementGenerate(n int) []map[string]interface{} {
	ran := rand.New(rand.NewSource(time.Now().UnixNano()))
	uniqueQuestions := make(map[string]struct{})
	questions := make([]map[string]interface{}, 0, n)

	for len(uniqueQuestions) < n {
		question, correctAnswer, wrongAnswers := generateConditionalExpressionQuestion(ran)
		if _, exists := uniqueQuestions[question]; !exists {
			uniqueQuestions[question] = struct{}{}
			questions = append(questions, map[string]interface{}{
				"question":       question,
				"correct_answer": correctAnswer,
				"wrong_answers":  wrongAnswers,
			})
		}
	}

	return questions
}
