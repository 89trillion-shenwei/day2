package model

import (
	"fmt"
)

var size int //栈的长度

type Stack interface {
	Push(e string) //向栈顶添加元素
	Pop() string   //移除栈顶元素
	Top() string   //获取栈顶元素（不删除）
	Size() int     //获取栈的元素个数
}

type sliceEntry struct {
	element []string
}

func NewSliceEntry() *sliceEntry {
	return &sliceEntry{}
}
func (entry *sliceEntry) Size() int {
	return len(entry.element)
}

func (entry *sliceEntry) Push(e string) {
	entry.element = append(entry.element, e)
}

func (entry *sliceEntry) Pop() string {
	size := entry.Size()
	if size == 0 {
		fmt.Println("stack is empty!")
		return ""
	}
	lastElement := entry.element[size-1]
	entry.element[size-1] = ""
	entry.element = entry.element[:size-1]
	return lastElement
}

func (entry *sliceEntry) Top() string {
	size := entry.Size()
	if size == 0 {
		fmt.Println("stack is empty!")
		return ""
	}
	return entry.element[size-1]
}
