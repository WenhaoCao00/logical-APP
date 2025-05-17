package main

import (
	"fmt"
	"math/rand"
	"time"
)

// // implies 函数
// func implies(a, b bool) bool {
// 	return !a || b
// }

// // 将布尔值转换为字符串
// func boolToString(b bool) string {
// 	if b {
// 		return "true"
// 	}
// 	return "false"
// }

// // 生成随机布尔值
// func generateRandomBool(ran *rand.Rand) bool {
// 	return ran.Intn(2) == 1
// }

// 计算逆否命题
func contrapositive(expr string) string {
	return "¬" + expr + "→¬" + expr
}

// 计算逆命题
func inverse(expr string) string {
	return "¬" + expr + "→¬" + expr
}

// 计算否命题
func converse(expr string) string {
	return expr + "→" + expr
}

// 生成随机表达式
func generateRandomExpression(ran *rand.Rand, variables []string) string {
	ops := []string{"∧", "∨", "→"}
	numVars := len(variables)
	expr := variables[ran.Intn(numVars)]
	for i := 0; i < ran.Intn(3)+1; i++ { // 随机选择1到3个运算符
		op := ops[ran.Intn(len(ops))]
		variable := variables[ran.Intn(numVars)]
		expr = fmt.Sprintf("(%s %s %s)", expr, op, variable)
	}
	return expr
}

// 生成条件表达式问题
func generateConditionalExpressionQuestion(ran *rand.Rand) (string, string, []string) {
	variables := []string{"p", "q", "r", "s"}
	expr := generateRandomExpression(ran, variables)

	// 随机选择一个问题类型
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

	// 生成三个不同的错误答案
	wrongAnswers := make(map[string]struct{})
	wrongAnswers[correctAnswer] = struct{}{} // 将正确答案加入集合以避免重复

	for len(wrongAnswers) < 4 {
		wrongExpr := generateRandomExpression(ran, variables)
		wrongAnswer := questionType.answerFunc(wrongExpr)
		wrongAnswers[wrongAnswer] = struct{}{}
	}

	// 从集合中删除正确答案并转换为列表
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
