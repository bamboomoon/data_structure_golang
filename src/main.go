package main

import (
	"fmt"
	"linkedList"
	"commentTeacher"

	"stack"
)

func main() {
	doubleStackExample()
}
func oneWayList() {
	//带头结点的单向链表

	linkedList := new(linkedList.LinkedList)
	linkedList.CreateLinkedList(10) //init length == 10 linkedList
	linkedList.PrintLinkList()
	fmt.Println("插入前长度：", linkedList.GetLinkedListLength())
	linkedList.InsertNode(2, 10) //insert one node
	fmt.Println("插入后长度：", linkedList.GetLinkedListLength())
	linkedList.PrintLinkList()
	linkedList.AddNodeInLikedListEnd(200) // add one node in the linkedList end
	linkedList.PrintLinkList()
	linkedList.RemoveNode(1) //remove one node
	linkedList.PrintLinkList()
	linkedList.ChangeNodeData(222, linkedList.GetLinkedListLength())
	linkedList.PrintLinkList()

}
func twoWayList() {

	//带头节点的双向链表
	twoWayLinkedList := new(linkedList.TwoWay)
	nodeData := []int{1, 2, 3}
	twoWayLinkedList.InitTwoWay_linkedList(nodeData)
	twoWayLinkedList.PrintLinkedList()
	//插入
	fmt.Println("插入:")
	err := twoWayLinkedList.InsertNode(10, 5)
	if err != nil {
		fmt.Println(err)
	}
	twoWayLinkedList.PrintLinkedList()
	//移除
	fmt.Println("移除:")
	err = twoWayLinkedList.RemoveNode(4)
	if err != nil {
		fmt.Println(err)
	}
	twoWayLinkedList.PrintLinkedList()

	//改
	fmt.Println("改：")
	err = twoWayLinkedList.ChangeNodeData(100, 3)
	if err != nil {
		fmt.Println(err)
	}
	twoWayLinkedList.PrintLinkedList()
}
func commentTeacherFunc(){
	//锦城学院教务网 评教
	commentTeacher.BeginComment()
}
func doubleStackExample(){
	doubleStack := stack.InitDoubleStack(20)
	for i := 1;i <=21;i++{
		if i < 8{
			doubleStack.Push(i,1)
		}else {
			doubleStack.Push(i,2)
		}

	}

	for i := 1;i <=21;i++{
		if i < 8 {
			fmt.Println(doubleStack.Pop(1))
		}else {
			fmt.Println(doubleStack.Pop(2))
		}
	}
}