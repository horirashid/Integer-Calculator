package main

import (
	"errors"
	"fmt"
)

type Stack struct {
	maxNum int          //规定栈最多放几个元素
	top    int          //目前栈顶的下标
	arr    [10]fraction //模拟栈
}

func (this *Stack) Push(val fraction) (err error) {
	if this.isFull() {
		fmt.Println("stack full")
		err = errors.New("stack full")
		return
	}
	//开始入栈操作
	//先向上走一步
	this.top++
	//再赋值
	this.arr[this.top] = val
	return
}
func (this *Stack) Pop() (val fraction, err error) {
	if this.isEmpty() {
		fmt.Println("stack empty")
		err = errors.New("stack empty")
		return
	}
	val = this.arr[this.top]
	this.top--
	return
}
func (this *Stack) List() (err error) {
	if this.isEmpty() {
		fmt.Println("stack empty")
		err = errors.New("stack empty")
		return
	}
	for i := this.top; i >= 0; i-- {
		fmt.Printf("arr[%d]=%d\n", i, this.arr[i])
	}
	return
}

func (this *Stack) isFull() bool {
	return this.top+1 >= this.maxNum
}
func (this *Stack) isEmpty() bool {
	return this.top == -1
}
