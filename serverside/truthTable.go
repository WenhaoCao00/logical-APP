package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func not(a bool) bool {
	return !a
}

func and(a, b bool) bool {
	return a && b
}

func or(a, b bool) bool {
	return a || b
}

func implies(a, b bool) bool {
	return !a || b
}

func randomBool(ran *rand.Rand) bool {
	return ran.Intn(2) == 0
}

func boolToString(b bool) string {
	if b {
		return "true"
	}
	return "false"
}

func generateExpressionAndAnswersForTruthTable(ran *rand.Rand) (string, bool, bool) {
	// 每次生成新的 p, q, r, s
	p, q, r, s := randomBool(ran), randomBool(ran), randomBool(ran), randomBool(ran)

	expressions := []struct {
		expr     string
		evaluate func() bool
	}{
		{"¬(p∨q)", func() bool { return not(or(p, q)) }},
		{"¬(p∧q)", func() bool { return not(and(p, q)) }},
		{"¬(p∨(q∧r))", func() bool { return not(or(p, and(q, r))) }},
		{"¬(p↔q)", func() bool { return not(p == q) }},
		{"¬((p∨¬q)∧r)", func() bool { return not(and(or(p, not(q)), r)) }},
		{"¬((p∧¬q)∨¬r)", func() bool { return not(or(and(p, not(q)), not(r))) }},
		{"¬(¬p∨q)", func() bool { return not(or(not(p), q)) }},
		{"¬((¬p∧q)∨¬r)", func() bool { return not(or(and(not(p), q), not(r))) }},
		{"¬((p∨q)∧(¬r∨s))", func() bool { return not(and(or(p, q), or(not(r), s))) }},
		{"¬((p→q)∧(r→¬s))", func() bool { return not(and(implies(p, q), implies(r, not(s)))) }},
		{"¬(¬(p∨q)∧r)", func() bool { return not(and(not(or(p, q)), r)) }},
		{"¬((¬p∨q)∧(r∧¬s))", func() bool { return not(and(or(not(p), q), and(r, not(s)))) }},
		{"¬((p∧¬q)∧(r∨s))", func() bool { return not(and(and(p, not(q)), or(r, s))) }},
		{"¬((¬p∨q)∨(¬r∧s))", func() bool { return not(or(or(not(p), q), and(not(r), s))) }},
		{"¬((p∨q)∧(¬r∧¬s))", func() bool { return not(and(and(p, q), and(not(r), not(s)))) }},
		{"¬((p∧¬q)∧(r∨¬s))", func() bool { return not(and(and(p, not(q)), or(r, not(s)))) }},
		{"¬((¬p∧¬q)∧(r∨s))", func() bool { return not(and(and(not(p), not(q)), or(r, s))) }},
		{"¬((p∨¬q)∨(¬r∧s))", func() bool { return not(or(or(p, not(q)), and(not(r), s))) }},
		{"¬((p∧q)∧(¬r∨¬s))", func() bool { return not(and(and(p, q), or(not(r), not(s)))) }},
		{"¬((¬p∧q)∧(r∧¬s))", func() bool { return not(and(and(not(p), q), and(r, not(s)))) }},
		{"¬((p∨¬q)∨(r∧¬s))", func() bool { return not(or(or(p, not(q)), and(r, not(s)))) }},
	}

	// 随机选择一个表达式
	randomIndex := ran.Intn(len(expressions))
	selectedExpression := expressions[randomIndex]

	// 生成正确答案和一个错误答案
	correctAnswer := selectedExpression.evaluate()
	incorrectAnswer := !correctAnswer

	// 构建问题字符串
	var question string
	expr := selectedExpression.expr
	if strings.Contains(expr, "r") && strings.Contains(expr, "s") {
		question = fmt.Sprintf("what is %s when p is %s, q is %s, r is %s, and s is %s",
			expr, boolToString(p), boolToString(q), boolToString(r), boolToString(s))
	} else if strings.Contains(expr, "r") {
		question = fmt.Sprintf("what is %s when p is %s, q is %s, and r is %s",
			expr, boolToString(p), boolToString(q), boolToString(r))
	} else {
		question = fmt.Sprintf("what is %s when p is %s and q is %s",
			expr, boolToString(p), boolToString(q))
	}

	return question, correctAnswer, incorrectAnswer
}

func truthTableGenerate(n int) []map[string]interface{} {
	ran := rand.New(rand.NewSource(time.Now().UnixNano()))
	uniqueQuestions := make(map[string]struct{})
	questions := make([]map[string]interface{}, 0, n)

	for len(uniqueQuestions) < n {
		question, correct, incorrect := generateExpressionAndAnswersForTruthTable(ran)
		if _, exists := uniqueQuestions[question]; !exists {
			uniqueQuestions[question] = struct{}{}
			questions = append(questions, map[string]interface{}{
				"question":       question,
				"correct_answer": strconv.FormatBool(correct),
				"wrong_answers": []string{strconv.FormatBool(incorrect)},
			})
		}
	}

	return questions
}

