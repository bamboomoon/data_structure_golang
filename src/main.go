package main

import (
	"./linkedList"
	"fmt"
)

func main() {
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
