package linkedList

import (
	"fmt"
	"math/rand"
)

type linkedListNode struct {
	data int             //数据域
	next *linkedListNode //指针域
}

type LinkedList struct {
	//头节点
	headerNode       *linkedListNode
	linkedListLength int //长度
}

//判空
func (linkedList *LinkedList) LinkedListIsEmpty() bool {
	if linkedList.linkedListLength == 0 {
		fmt.Println("链表为空")
		return true
	}
	return false
}

//创建链表
func (linkedList *LinkedList) CreateLinkedList(length int) {
	//头节点
	linkedListheaderNode := &linkedListNode{data: rand.Int(), next: nil}
	linkedList.linkedListLength = length
	p := linkedListheaderNode
	for i := 1; i < length; i++ {
		newNode := &linkedListNode{rand.Int(), nil}
		p.next = newNode
		p = newNode
	}
	linkedList.headerNode = linkedListheaderNode
}

//打印链表
func (linkedList *LinkedList) PrintLinkList() {
	if linkedList.LinkedListIsEmpty() {
		return
	}
	p := linkedList.headerNode.next
	for i := 0; i < linkedList.linkedListLength; i++ {
		fmt.Printf("nodeAddress:%p-node.data:%d-node.next_address:%p\n", p, p.data, p.next)
		p = p.next
	}
}
