package main

import (
	"math/rand"
	"time"
)


type Expr interface {
	String() string
}


type Not struct {
	Expr Expr
}

type And struct {
	Left, Right Expr
}

type Or struct {
	Left, Right Expr
}

type Var struct {
	Name string
}


func (n Not) String() string {
	return "¬" + n.Expr.String()
}

func (a And) String() string {
	return "(" + a.Left.String() + "∧" + a.Right.String() + ")"
}

func (o Or) String() string {
	return "(" + o.Left.String() + "∨" + o.Right.String() + ")"
}

func (v Var) String() string {
	return v.Name
}


func deMorganShift(expr Expr) Expr {
	switch e := expr.(type) {
	case Not:
		switch inner := e.Expr.(type) {
		case Or:
			return And{deMorganShift(Not{inner.Left}), deMorganShift(Not{inner.Right})}
		case And:
			return Or{deMorganShift(Not{inner.Left}), deMorganShift(Not{inner.Right})}
		case Not:
			return deMorganShift(inner.Expr)
		}
	}
	return expr
}

func randomVar(ran *rand.Rand) Var {
	variables := []string{"p", "q", "r", "s"}
	return Var{variables[ran.Intn(len(variables))]}
}

func generateExpressionAndAnswersFordeMorgan(ran *rand.Rand) (string, string, []string) {

	expressions := []Expr{
		Not{Or{randomVar(ran), randomVar(ran)}},
		Not{And{randomVar(ran), randomVar(ran)}},
		Not{Or{randomVar(ran), And{randomVar(ran), randomVar(ran)}}},
		Not{And{randomVar(ran), Or{randomVar(ran), randomVar(ran)}}},
		Not{And{Or{randomVar(ran), randomVar(ran)}, randomVar(ran)}},
		Not{Or{And{randomVar(ran), randomVar(ran)}, randomVar(ran)}},
		Not{Or{Or{randomVar(ran), randomVar(ran)}, randomVar(ran)}},
		Not{And{And{randomVar(ran), randomVar(ran)}, randomVar(ran)}},
		Not{Or{Not{randomVar(ran)}, randomVar(ran)}},
		Not{And{Not{randomVar(ran)}, randomVar(ran)}},
	}


	randomIndex := ran.Intn(len(expressions))
	selectedExpression := expressions[randomIndex]
	simplifiedExpression := deMorganShift(selectedExpression)


	question := selectedExpression.String()
	answer := simplifiedExpression.String()


	wrongAnswers := make(map[string]struct{})
	wrongAnswers[answer] = struct{}{} 

	for len(wrongAnswers) < 4 {
		wrongExpr := deMorganShift(expressions[ran.Intn(len(expressions))])
		wrongAnswer := wrongExpr.String()
		wrongAnswers[wrongAnswer] = struct{}{}
	}


	delete(wrongAnswers, answer)
	wrongAnswersList := make([]string, 0, 3)
	for ans := range wrongAnswers {
		wrongAnswersList = append(wrongAnswersList, ans)
	}

	return question, answer, wrongAnswersList
}

func deMorganGenerate(n int) []map[string]interface{} {
	ran := rand.New(rand.NewSource(time.Now().UnixNano()))
	uniqueQuestions := make(map[string]struct{})
	questions := make([]map[string]interface{}, 0, n)

	for len(uniqueQuestions) < n {
		question, answer, wrongAnswers := generateExpressionAndAnswersFordeMorgan(ran)
		if _, exists := uniqueQuestions[question]; !exists {
			uniqueQuestions[question] = struct{}{}
			questions = append(questions, map[string]interface{}{
				"question":       question,
				"correct_answer": answer,
				"wrong_answers":  wrongAnswers,
			})
		}
	}

	return questions
}



