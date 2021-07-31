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

//栈内存string类型的数据
type sliceEntry struct {
	element []string
}

// NewSliceEntry 新建栈
func NewSliceEntry() *sliceEntry {
	return &sliceEntry{}
}

// Size 获取栈的元素个数
func (entry *sliceEntry) Size() int {
	return len(entry.element)
}

// Push 向栈顶添加元素
func (entry *sliceEntry) Push(e string) {
	entry.element = append(entry.element, e)
}

// Pop 移除栈顶元素
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

// Top 获取栈顶元素（不删除）
func (entry *sliceEntry) Top() string {
	size := entry.Size()
	if size == 0 {
		fmt.Println("stack is empty!")
		return ""
	}
	return entry.element[size-1]
}
