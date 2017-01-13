package stack

import "fmt"

type sqDoubleStack struct {
	data []int //两个栈一起所占用的空间
	top1 int   //stack1的栈顶
	top2 int   //指向stack2的栈顶
}

func InitDoubleStack(length int) *sqDoubleStack {
	return &sqDoubleStack{make([]int, length), -1, length}
}

//压栈
func (doubleStack *sqDoubleStack) Push(data, stackNum int) {
	if doubleStack.top1+1 == doubleStack.top2 { //他们连个相邻的时候就表明=》栈满
		return
	}
	if stackNum == 1 {
		doubleStack.top1++
		doubleStack.data[doubleStack.top1] = data
	}
	if stackNum == 2 {
		doubleStack.top2--
		doubleStack.data[doubleStack.top2] = data
	}
}

//出栈
func (doubleStack *sqDoubleStack) Pop(stackNum int) int {
	if stackNum == 1 {
		if doubleStack.top1 == -1 { //空栈
			return 404
		} else {
			index := doubleStack.top1
			doubleStack.top1--
			return doubleStack.data[index]

		}

	}
	if stackNum == 2 {
		if doubleStack.top2 == len(doubleStack.data) { //空
			return 404
		} else {
			index := doubleStack.top2
			doubleStack.top2++
			return doubleStack.data[index]

		}
	}
	return 404
}

func (doubleStack *sqDoubleStack) PrintDoubleStack() {
	for v := range doubleStack.data {
		fmt.Println(v)
	}
}
