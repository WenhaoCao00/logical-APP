package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Question struct to represent a quiz question
type Question struct {
	ID            string   `json:"id"`
	Question      string   `json:"question"`
	Options       []string `json:"options"`
	CorrectAnswer string   `json:"correct_answer"`
	Difficulty    string   `json:"difficulty"`
}

// QuizGenerator struct to hold quiz generation methods
type QuizGenerator struct{}

var _random = rand.New(rand.NewSource(time.Now().UnixNano()))

func (qg QuizGenerator) generateQuiz(numQuestions int, difficulty string) []Question {
	generatedQuestions := make(map[string]struct{})
	questions := []Question{}

	for len(questions) < numQuestions {
		newQuestion := qg.generateRandomQuestion(fmt.Sprintf("%d", len(questions)), difficulty)

		// if the question has not been generated before, add it to the list
		if _, exists := generatedQuestions[newQuestion.Question]; !exists {
			generatedQuestions[newQuestion.Question] = struct{}{}
			questions = append(questions, newQuestion)
		}
	}

	return questions
}

func (qg QuizGenerator) generateRandomQuestion(id string, difficulty string) Question {
	// 随机选择一个问题类型
	questionType := _random.Intn(6)
	switch questionType {
	case 0:
		return qg.generatePropositionQuestion(id, difficulty)
	case 1:
		return qg.generateTruthTableQuestion(id, difficulty)
	case 2:
		return qg.generateLogicalEquivalenceQuestion(id, difficulty)
	case 3:
		return qg.generateDeMorganLawQuestion(id, difficulty)
	case 4:
		return qg.generateConditionalStatementQuestion(id, difficulty)
	case 5:
		return qg.generatePropositionalFunctionQuestion(id, difficulty)
	default:
		return qg.generatePropositionQuestion(id, difficulty)
	}
}

// 命题和逻辑连接词的问题
func (qg QuizGenerator) generatePropositionQuestion(id string, difficulty string) Question {
	symbols := []string{"∧", "∨", "¬", "→", "↔"}
	symbol := symbols[_random.Intn(len(symbols))]
	question := fmt.Sprintf("Which logical connective is represented by the symbol %s?", symbol)
	var options []string
	var correctAnswer string

	switch symbol {
	case "∧":
		options = []string{"AND", "OR", "NOT", "IF-THEN"}
		correctAnswer = "AND"
	case "∨":
		options = []string{"AND", "OR", "NOT", "IF-THEN"}
		correctAnswer = "OR"
	case "¬":
		options = []string{"AND", "OR", "NOT", "IF-THEN"}
		correctAnswer = "NOT"
	case "→":
		options = []string{"AND", "OR", "NOT", "IF-THEN"}
		correctAnswer = "IF-THEN"
	case "↔":
		options = []string{"AND", "NOT", "IF-THEN", "IF AND ONLY IF"}
		correctAnswer = "IF AND ONLY IF"
	default:
		options = []string{"AND", "OR", "NOT", "IF-THEN"}
		correctAnswer = "AND"
	}

	return Question{
		ID:            id,
		Question:      question,
		Options:       options,
		CorrectAnswer: correctAnswer,
		Difficulty:    difficulty,
	}
}

// 真值表的问题
func (qg QuizGenerator) generateTruthTableQuestion(id string, difficulty string) Question {
	p := _random.Intn(2) == 1
	q := _random.Intn(2) == 1
	pStr := "false"
	if p {
		pStr = "true"
	}
	qStr := "false"
	if q {
		qStr = "true"
	}

	var question string
	var options []string
	var correctAnswer string

	if difficulty == "easy" {
		question = fmt.Sprintf("What is the truth value of p∧q when p is %s and q is %s?", pStr, qStr)
		expressionValue := p && q
		options = []string{"True", "False"}
		if expressionValue {
			correctAnswer = "True"
		} else {
			correctAnswer = "False"
		}
	} else if difficulty == "medium" {
		question = fmt.Sprintf("What is the truth value of (p→q) when p is %s and q is %s?", pStr, qStr)
		expressionValue := !p || q
		options = []string{"True", "False"}
		if expressionValue {
			correctAnswer = "True"
		} else {
			correctAnswer = "False"
		}
	} else {
		// hard
		question = fmt.Sprintf("What is the truth value of (p→q)∧(¬q) when p is %s and q is %s?", pStr, qStr)
		pImpliesQ := !p || q
		notQ := !q
		expressionValue := pImpliesQ && notQ
		options = []string{"True", "False"}
		if expressionValue {
			correctAnswer = "True"
		} else {
			correctAnswer = "False"
		}
	}

	return Question{
		ID:            id,
		Question:      question,
		Options:       options,
		CorrectAnswer: correctAnswer,
		Difficulty:    difficulty,
	}
}

// 逻辑等价的问题
func (qg QuizGenerator) generateLogicalEquivalenceQuestion(id string, difficulty string) Question {
	expressions := []string{"¬(p∨q)", "¬(p∧q)", "p→q", "p↔q"}
	expression := expressions[_random.Intn(len(expressions))]
	questionTemplate := "Which of the following is equivalent to %s?"
	question := fmt.Sprintf(questionTemplate, expression)

	var options []string
	var correctAnswer string

	switch expression {
	case "¬(p∨q)":
		options = []string{"¬p∧¬q", "¬p∨¬q", "p∧q", "p∨q"}
		correctAnswer = "¬p∧¬q"
	case "¬(p∧q)":
		options = []string{"¬p∨¬q", "¬p∧¬q", "p∨q", "p∧q"}
		correctAnswer = "¬p∨¬q"
	case "p→q":
		options = []string{"¬p∨q", "p∧¬q", "¬p∧¬q", "p∨¬q"}
		correctAnswer = "¬p∨q"
	case "p↔q":
		options = []string{"(p∧q)∨(¬p∧¬q)", "(p∨q)∧(¬p∨¬q)", "p∧¬q", "¬p∧q"}
		correctAnswer = "(p∧q)∨(¬p∧¬q)"
	default:
		options = []string{"¬p∧¬q", "¬p∨¬q", "p∧q", "p∨q"}
		correctAnswer = "¬p∧¬q"
	}

	return Question{
		ID:            id,
		Question:      question,
		Options:       options,
		CorrectAnswer: correctAnswer,
		Difficulty:    difficulty,
	}
}

// De Morgan 定律的问题
func (qg QuizGenerator) generateDeMorganLawQuestion(id string, difficulty string) Question {
	var expression string
	questionTemplate := "Which of the following is equivalent to %s by De Morgan's Law?"
	var options []string
	var correctAnswer string

	if difficulty == "easy" {
		expression = "¬(p∨q)"
		options = []string{"¬p∧¬q", "¬p∨¬q", "p∧¬q", "p∨¬q"}
		correctAnswer = "¬p∧¬q"
	} else if difficulty == "medium" {
		expression = "¬(p∧q)"
		options = []string{"¬p∨¬q", "¬p∧¬q", "p∨q", "p∧q"}
		correctAnswer = "¬p∨¬q"
	} else {
		// hard
		expression = "¬((p∧q)∨r)"
		options = []string{"¬p∨¬q∧¬r", "¬p∧¬q∧¬r", "¬p∧¬q∨¬r", "¬p∨¬q∨¬r"}
		correctAnswer = "¬p∧¬q∧¬r"
	}

	question := fmt.Sprintf(questionTemplate, expression)

	return Question{
		ID:            id,
		Question:      question,
		Options:       options,
		CorrectAnswer: correctAnswer,
		Difficulty:    difficulty,
	}
}

// 条件和双条件语句的问题
func (qg QuizGenerator) generateConditionalStatementQuestion(id string, difficulty string) Question {
	var expression string
	questionTemplate := "What is the %s of p→q?"
	var options []string
	var correctAnswer string

	if difficulty == "easy" {
		expression = "contrapositive"
		options = []string{"¬q→¬p", "q→p", "¬p→¬q", "p↔q"}
		correctAnswer = "¬q→¬p"
	} else if difficulty == "medium" {
		expression = "inverse"
		options = []string{"¬p→¬q", "¬q→¬p", "q→p", "p↔q"}
		correctAnswer = "¬p→¬q"
	} else {
		// hard
		expression = "converse"
		options = []string{"q→p", "¬q→¬p", "p→q", "q→¬p"}
		correctAnswer = "q→p"
	}

	question := fmt.Sprintf(questionTemplate, expression)

	return Question{
		ID:            id,
		Question:      question,
		Options:       options,
		CorrectAnswer: correctAnswer,
		Difficulty:    difficulty,
	}
}

// 命题函数和量词的问题
func (qg QuizGenerator) generatePropositionalFunctionQuestion(id string, difficulty string) Question {
	var expression string
	questionTemplate := "What is the %s of ∀x P(x)?"
	var options []string
	var correctAnswer string

	if difficulty == "easy" {
		expression = "negation"
		options = []string{"∃x ¬P(x)", "¬∃x P(x)", "¬∀x P(x)", "∃x P(x)"}
		correctAnswer = "∃x ¬P(x)"
	} else if difficulty == "medium" {
		expression = "negation of ∃x P(x)"
		options = []string{"∀x ¬P(x)", "¬∀x P(x)", "¬∃x P(x)", "∃x ¬P(x)"}
		correctAnswer = "∀x ¬P(x)"
	} else {
		// hard
		questionTemplate = "What is the logical form of the statement '%s'?"
		expression = "There is no x such that P(x)"
		options = []string{"∀x ¬P(x)", "¬∃x P(x)", "∃x ¬P(x)", "¬∀x P(x)"}
		correctAnswer = "¬∃x P(x)"
	}

	question := fmt.Sprintf(questionTemplate, expression)

	return Question{
		ID:            id,
		Question:      question,
		Options:       options,
		CorrectAnswer: correctAnswer,
		Difficulty:    difficulty,
	}
}

func generateFixedQuestions() map[string][]Question {
	qg := QuizGenerator{}

	easyQuestions := qg.generateQuiz(5, "easy")
	mediumQuestions := qg.generateQuiz(5, "medium")
	hardQuestions := qg.generateQuiz(5, "hard")

	return map[string][]Question{
		"easy":   easyQuestions,
		"medium": mediumQuestions,
		"hard":   hardQuestions,
	}
}
