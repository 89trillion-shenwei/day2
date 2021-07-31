package handler

import (
	"day2/internal"
	model2 "day2/internal/model"
	"strconv"
)

//区分是不是运算符
func sortString(s string) bool {
	if s == "+" || s == "-" || s == "*" || s == "/" {
		return true
	}
	return false
}

func sortNum(b byte) bool {
	if b >= 48 && b <= 57 {
		return true
	}
	return false
}

// Analy 解析字符串，运算符入一个栈，数字合到一起在进一个栈
func Analy(s string) (e1, e2 model2.Stack, err error) {
	//e1 := model.NewSliceEntry()
	//e2 := model.NewSliceEntry()
	//判断字符串中会不会连续出现两个以上的运算符，tmp=1说明上次入栈的是运算符
	var tmp int = 1
	for i := len(s) - 1; i >= 0; i-- {
		k := string(s[i])
		//进栈，运算符直接进栈，数字拼接完在进栈
		if sortString(k) {
			//连续两个运算符，或者运算符在首位或者尾位会报错
			if tmp == 1 || i == 0 || i == len(s)-1 {
				return nil, nil, internal.IllegalExpressionError("非法的表达式")
			}
			e1.Push(k)
		} else if sortNum(s[i]) {
			ee := model2.NewSliceEntry()
			ee.Push(k)
			tmp = 0
			for j := i - 1; j >= 0; j-- {
				if sortNum(s[j]) {
					//if s[j]>=48&&s[j]<=57 {
					ee.Push(string(s[j]))
					if j == 0 {
						i = -1
						break
					}
					continue
				} else {
					i = j + 1
					break
				}

			}
			var s string
			len := ee.Size()
			if len != 0 {
				for k := 0; k < len; k++ {
					s += ee.Pop()
				}
				e2.Push(s)
			}
		} else {
			return nil, nil, internal.IllegalExpressionError("有非法符号")
		}
	}
	return e1, e2, nil
}

// Couter 计算方法
func Couter(e1, e2 model2.Stack) int {
	if e1.Size() == 0 {
		in, _ := strconv.Atoi(e2.Top())
		return in
	}
	len1 := e1.Size()
	e3 := model2.NewSliceEntry()
	e4 := model2.NewSliceEntry()
	//先算乘除
	for j := 0; j < len1; j++ {
		var sum int
		if e1.Top() == "*" || e1.Top() == "/" {
			if e1.Top() == "*" {
				in1, _ := strconv.Atoi(e2.Pop())
				in2, _ := strconv.Atoi(e2.Pop())
				sum = in1 * in2
				e1.Pop()
				str := strconv.Itoa(sum)
				e2.Push(str)
			} else if e1.Top() == "/" {
				in1, _ := strconv.Atoi(e2.Pop())
				in2, _ := strconv.Atoi(e2.Pop())
				sum = in1 / in2
				e1.Pop()
				str := strconv.Itoa(sum)
				e2.Push(str)
			}
			if j == len1-1 {
				e4.Push(e2.Top())
			}
		} else {
			e4.Push(e2.Pop())
			e3.Push(e1.Pop())
		}
	}
	e4.Push(e2.Pop())
	//将栈内数据更改为正确的顺序
	e5 := A2B(e3)
	e6 := A2B(e4)
	len2 := e5.Size()
	//再算加减
	for i := 0; i < len2; i++ {
		if e5.Top() == "+" {
			in1, _ := strconv.Atoi(e6.Pop())
			in2, _ := strconv.Atoi(e6.Pop())
			str := strconv.Itoa(in1 + in2)
			e6.Push(str)
		} else {
			in1, _ := strconv.Atoi(e6.Pop())
			in2, _ := strconv.Atoi(e6.Pop())
			str := strconv.Itoa(in1 - in2)
			e6.Push(str)
		}
		e5.Pop()
	}
	in3, _ := strconv.Atoi(e6.Top())
	return in3
}

// A2B 将栈里的数据逆序
func A2B(e model2.Stack) model2.Stack {
	len := e.Size()
	e1 := model2.NewSliceEntry()
	for i := 0; i < len; i++ {
		e1.Push(e.Pop())
	}
	return e1
}
