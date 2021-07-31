package handler

import (
	model2 "day2/internal/model"
	"strconv"
)

//逆波兰解法

//Str2Strs 将数字整合到一起
func Str2Strs(exp string) []string {
	strs := make([]string, 0)
	for i := 0; i < len(exp); i++ {
		k := string(exp[i])
		if sortString(string(exp[i])) {
			strs = append(strs, string(exp[i]))
		} else {
			stack := model2.NewSliceEntry()
			stack.Push(k)
			for j := i + 1; j < len(exp); j++ {
				if sortNum(exp[j]) {
					//if s[j]>=48&&s[j]<=57 {
					stack.Push(string(exp[j]))
					if j == len(exp)-1 {
						i = len(exp)
						break
					}
					continue
				} else {
					i = j - 1
					break
				}

			}
			var s string
			len := stack.Size()
			if len != 0 {
				for k := 0; k < len; k++ {
					s = stack.Pop() + s
				}
				strs = append(strs, s)
			}

		}
	}
	return strs
}

// 中缀表达式转后缀表达式
func Infix2ToPostfix(strs []string) []string {
	stack := model2.NewSliceEntry()
	postfix := make([]string, 0)
	expLen := len(strs)
	// 遍历整个表达式
	for i := 0; i < expLen; i++ {

		char := strs[i]
		switch char {
		case " ":
			continue
			// 数字则直接输出
		case "+", "-", "*", "/":
			// 操作符：遇到高优先级的运算符，不断弹出，直到遇见更低优先级运算符
			for stack.Size() != 0 {
				top := stack.Top()
				if isLower(top, char) {
					break
				}
				postfix = append(postfix, top)
				stack.Pop()
			}
			// 低优先级的运算符入栈
			stack.Push(char)

		default:
			postfix = append(postfix, strs[i])
		}
	}

	// 栈不空则全部输出
	for stack.Size() != 0 {
		postfix = append(postfix, stack.Pop())
	}

	return postfix
}

//计算器方法
func Calculate(postfix []string) int {
	stack := model2.NewSliceEntry()
	fixLen := len(postfix)
	for i := 0; i < fixLen; i++ {
		nextChar := string(postfix[i])
		// 数字：直接压栈
		if !sortString(postfix[i]) {
			stack.Push(nextChar)
		} else {
			// 操作符：取出两个数字计算值，再将结果压栈
			num1, _ := strconv.Atoi(stack.Pop())
			num2, _ := strconv.Atoi(stack.Pop())
			switch nextChar {
			case "+":
				stack.Push(strconv.Itoa(num2 + num1))
			case "-":
				stack.Push(strconv.Itoa(num2 - num1))
			case "*":
				stack.Push(strconv.Itoa(num2 * num1))
			case "/":
				stack.Push(strconv.Itoa(num2 / num1))
			}
		}
	}
	result, _ := strconv.Atoi(stack.Top())
	return result
}

// 比较运算符栈栈顶 top 和新运算符 newTop 的优先级高低
func isLower(top string, newTop string) bool {
	switch top {
	case "+", "-":
		if newTop == "*" || newTop == "/" {
			return true
		}
	}
	return false
}

func StringIslegal(s string) bool {
	//运算符不能在首位和尾位
	if sortString(string(s[0])) || sortString(string(s[len(s)-1])) {
		return false
	}
	//临时变量判断是否连续出现运算符
	temp := 0
	//判断字符串是否合法
	for i := 0; i < len(s); i++ {
		if sortNum(s[i]) {
			temp = 0
		} else if sortString(string(s[i])) {
			if temp == 1 {
				return false
			}
			temp = 1
		} else {
			return false
		}
	}
	return true
}
