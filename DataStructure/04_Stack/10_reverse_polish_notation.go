package main

import (
	"fmt"
	"strconv"
	"strings"
)

// 逆波兰表达式 RPN
func main() {

	// 0. 从左到右扫描中缀表达式
	// 1. 遇到操作数直接入栈s2
	// 2. 遇到运算符时，比较运算符栈s1的栈顶运算符的优先级
	// 2.1 若s1为空，或栈顶运算符为左括号 ( 则直接入s1
	// 2.2 若遇到的运算符比s1栈顶的运算符优先级高，则直接入s1
	// 2.3 若遇到的运算符比s1栈顶的运算符优先级小或等于，则先弹出s1栈顶的运算符压入到s2中，再跳转至2.1与s1新的栈顶运算符进行比较
	// 3. 遇到括号时
	// 3.1 若遇到的是左括号 ( 则直接入s2栈
	// 3.2 若遇到的是右括号 ) 则依次弹出s1中的运算符并压入s2，直到遇到右括号为止，并将一对括号丢弃
	// 4. 将s1剩余栈压入s2
	// 5. s2中pop出的元素的逆序即中缀表达式转后缀表达式的结果

	infixNotationExpression := "1+((2+3)*4)-5"
	arr1 := strings.Split(infixNotationExpression, "")
	finalStr := ""
	stack := &Stack{-1, make([]interface{}, 10)}
	for _, str := range arr1 {
		if IsOperator(str) {
			if str == "(" {
				// 入运算符栈
				stack.Put(str)
			} else if str == ")" {
				// pop 到 ( 为止
				for stack.Peek().(string) != "(" {
					finalStr += stack.Pop().(string) + " "
				}
				stack.Pop()
			} else {
				for {
					empty := stack.IsEmpty()
					if !empty && (CalcPriority(str) <= CalcPriority(stack.Peek().(string))) {
						finalStr += stack.Pop().(string) + " "
					} else {
						// 入运算符栈
						stack.Put(str)
						break
					}
				}
			}
		} else {
			finalStr += str + " "
		}
	}

	// 将运算符中的元素都弹出压入s2
	for !stack.IsEmpty() {
		finalStr += stack.Pop().(string) + " "
	}

	finalStr = strings.TrimRight(finalStr," ")
	fmt.Println("finalStr#", finalStr)

	//expression := "3 4 + 5 * 6 -"
	arr := strings.Split(finalStr, " ")
	numStack := &Stack{-1, make([]interface{}, 10)}
	for _, value := range arr {
		if IsOperator(value) {
			after := numStack.Pop()  // 栈顶
			before := numStack.Pop() // 次顶
			num := CalcValue(before.(float32), after.(float32), value)
			numStack.Put(num)
		} else {
			num, _ := strconv.ParseFloat(value, 32)
			numStack.Put(float32(num))
		}
	}

	fmt.Println(numStack.Pop())
}

func IsOperator(char string) bool {
	return char == "+" || char == "-" || char == "*" || char == "/" || char == "(" || char == ")"
}

func CalcValue(num1, num2 float32, operator string) float32 {
	if operator == "*" {
		return num1 * num2
	} else if operator == "/" {
		return num1 / num2
	} else if operator == "+" {
		return num1 + num2
	} else if operator == "-" {
		return num1 - num2
	} else {
		return -1
	}
}

func CalcPriority(operator string) int {
	if operator == "*" || operator == "/" {
		return 1000
	} else if operator == "+" || operator == "-" {
		return 100
	} else {
		return 0
	}
}

type Stack struct {
	Top   int
	Array []interface{}
}

func (stack *Stack) Put(value interface{}) {
	if stack.IsFull() {
		panic("stack is full")
	} else {
		stack.Top++
		stack.Array[stack.Top] = value
	}
}

func (stack *Stack) Pop() interface{} {
	if stack.IsEmpty() {
		panic("stack is empty")
	} else {
		i := stack.Array[stack.Top]
		stack.Top--
		return i
	}
}

func (stack *Stack) IsEmpty() bool {
	return stack.Top < 0
}

func (stack *Stack) IsFull() bool {
	return stack.Top >= len(stack.Array)-1
}

func (stack *Stack) Peek() interface{} {
	if stack.Top < 0 {
		panic("stack is empty")
	} else {
		return stack.Array[stack.Top]
	}
}
