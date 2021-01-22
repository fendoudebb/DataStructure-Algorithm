package main

import (
	"fmt"
	"strconv"
)

func main() {
	//fmt.Println("+#", '+', "-#", '-', "*#", '*', "/#", '/')
	expression := "2+7*2-2*4+6"
	calc := NewCalc(expression)
	fmt.Printf("%+v", calc)
}

type Calculator struct {
	NumTop        int
	OperatorTop   int
	NumStack      []float32
	OperatorStack []int32
}

func NewCalc(expression string) *Calculator {
	calc := &Calculator{-1, -1, make([]float32, 32), make([]int32, 32)}

	// 处理多位数
	keepNum := ""
	for i, v := range expression {
		if v >= 48 && v <= 57 {
			// 0-9
			keepNum += string(v)
			if i == len(expression)-1 {
				num, _ := strconv.Atoi(keepNum)
				calc.PutNum(float32(num))
			} else {
				if expression[i+1] < 48 || expression[i+1] > 57{
					// 下一位操作符
					num, _ := strconv.Atoi(keepNum)
					calc.PutNum(float32(num))
					keepNum = ""
				}
			}
		} else {
			// 操作符 + - * /
			operator := calc.PeekOperator()
			if operator == -1 {
				calc.PutOperator(v)
			} else if Priority(operator) >= Priority(v) {
				// 运算等级小于或等于，直接pop出NumStack两个数进行计算
				calc.Calc()
				calc.PutOperator(v)
			} else {
				calc.PutOperator(v)
			}
		}
	}

	for calc.PeekOperator() != -1 {
		calc.Calc()
	}

	fmt.Println("最终结果#", calc.PopNum())

	return calc
}

func (calc *Calculator) PutNum(value float32) {
	if calc.NumTop < len(calc.NumStack)-1 {
		calc.NumTop++
		calc.NumStack[calc.NumTop] = value
	} else {
		fmt.Println("num stack is full")
	}
}

func (calc *Calculator) PopNum() float32 {
	if calc.NumTop < 0 {
		return -1
	} else {
		fmt.Println("pop num#", calc.NumStack[calc.NumTop])
		defer func() {
			calc.NumTop--
		}()
		return calc.NumStack[calc.NumTop]
	}
}

func (calc *Calculator) PutOperator(value int32) {
	if calc.OperatorTop < len(calc.OperatorStack)-1 {
		calc.OperatorTop++
		calc.OperatorStack[calc.OperatorTop] = value
	} else {
		fmt.Println("operator stack is full")
	}
}

func (calc *Calculator) PopOperator() int32 {
	if calc.OperatorTop < 0 {
		fmt.Println("operator stack is empty")
		return -1
	} else {
		fmt.Println("pop operator#", calc.OperatorStack[calc.OperatorTop])
		defer func() {
			calc.OperatorTop--
		}()
		return calc.OperatorStack[calc.OperatorTop]
	}
}

func (calc *Calculator) PeekOperator() int32 {
	if calc.OperatorTop < 0 {
		return -1
	} else {
		return calc.OperatorStack[calc.OperatorTop]
	}
}

func (calc *Calculator) Calc() {
	last := calc.PopNum()
	previous := calc.PopNum()
	popOperator := calc.PopOperator()
	// - 号 需变换符号
	// 7-2+6 先运算2和6，则更换+为-
	if calc.PeekOperator() == '-' {
		if popOperator == '+'{
			popOperator = '-'
		} else if popOperator == '-'{
			popOperator = '+'
		}
	}
	num := Calc(previous, last, popOperator)
	calc.PutNum(num)
}

func Calc(num1, num2 float32, operator int32) float32 {
	if operator == '*' {
		return num1 * num2
	} else if operator == '/' {
		return num1 / num2
	} else if operator == '+' {
		return num1 + num2
	} else if operator == '-' {
		return num1 - num2
	} else {
		return -1
	}
}

func Priority(operator int32) int {
	if operator == '*' || operator == '/' {
		return 1000
	} else if operator == '+' || operator == '-' {
		return 100
	} else {
		return 0
	}
}
