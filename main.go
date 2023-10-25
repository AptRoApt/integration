package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/Pramod-Devireddy/go-exprtk"
)

const e = 2.71828182845904523536

func getBounds() (a float64, b float64) {
	for {
		fmt.Println("Введите нижнюю границу интегрирования: ")
		fmt.Scanf("%f", &a)
		fmt.Println("Введите верхнюю границу интегрирования: ")
		fmt.Scanf("%f", &b)
		if a == b {
			fmt.Println("Границы не могут быть равны!")
			continue
		} else if b < a {
			fmt.Println("Верхняя граница не может быть меньше нижней!")
			continue
		}
		return a, b
	}
}

func getExpression() (expr string) {
	fmt.Println("Введите интегрируемое выражение: ")
	reader := bufio.NewReader(os.Stdin)
	expr, _ = reader.ReadString('\n')
	expr = strings.ReplaceAll(expr, " ", "")
	return expr
}

func getIterations() (n int) {
	for {
		fmt.Println("Введите количество итераций: ")
		fmt.Scan(&n)
		if n%2 == 1 {
			fmt.Println("Количество итераций должно быть чётным!")
			continue
		}
		return n
	}
}

func integrate(a float64, b float64, expression string, n int) (float64, error) {
	parser := exprtk.NewExprtk()
	defer parser.Delete()
	parser.SetExpression(expression)
	parser.AddDoubleVariable("x")
	parser.AddDoubleVariable("e")
	parser.SetDoubleVariableValue("e", e)
	delta := (b - a) / float64(n)
	var answer float64
	for x, i := a, 0; i < n; x, i = x+delta, i+1 {
		parser.SetDoubleVariableValue("x", x)
		var arg float64
		if (i == 0) || (i == n-1) {
			arg = 1
		} else if i%2 == 1 {
			arg = 4
		} else if i%2 == 0 {
			arg = 2
		}
		err := parser.CompileExpression()
		if err != nil {
			return 0, errors.New(err.Error())
		}
		value := parser.GetEvaluatedValue()
		answer += arg * value
	}
	answer = answer * (delta / 3)
	return answer, nil
}
func main() {
	a, b := getBounds()
	expression := getExpression()
	n := getIterations()
	answer, err := integrate(a, b, expression, n)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Значение интеграла: %f\n", answer)
	fmt.Println(expression)
	if expression == "e^(-((x^2)/2))" || expression == "e^-((x^2)/2)" {
		lpl := 0.398942280401 * answer
		fmt.Printf("Значение нормированной функции Лапласа: %f", lpl)
	}
	// parser := exprtk.NewExprtk()
	// defer parser.Delete()
	// parser.SetExpression("2.7^(-((x^2)/2))")
	// parser.AddDoubleVariable("x")
	// parser.SetDoubleVariableValue("x", 5)
	// err := parser.CompileExpression()
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }
	// value := parser.GetEvaluatedValue()
	// fmt.Printf("%f", value)
}
