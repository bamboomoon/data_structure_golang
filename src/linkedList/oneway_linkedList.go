package linkedList

import (
	"errors"
	"fmt"
	"math/rand"
)

type linkedListNode struct {
	data int             //data domain
	next *linkedListNode //pointer domain
}

type LinkedList struct {
	//头节点
	headerNode       *linkedListNode
	linkedListLength int //length
}

//operation result
const (
	Success int = iota //成功
	Fail               //失败
)

//判空
func (linkedList *LinkedList) LinkedListIsEmpty() bool {
	if linkedList.linkedListLength == 0 {
		fmt.Println("链表为空")
		return true
	}
	return false
}

//链表的长度
func (linkedList *LinkedList) GetLinkedListLength() int {
	return linkedList.linkedListLength
}

//创建链表
func (linkedList *LinkedList) CreateLinkedList(length int) error {
	length = int(length)
	if length <= 0 {
		return errors.New("intput length incorrect")
	}
	//头节点
	linkedListheaderNode := &linkedListNode{data: rand.Int(), next: nil}
	linkedList.linkedListLength = length
	p := linkedListheaderNode
	for i := 1; i <= length; i++ {
		newNode := &linkedListNode{rand.Int(), nil}
		p.next = newNode
		p = newNode
	}
	linkedList.headerNode = linkedListheaderNode
	return nil
}

//打印链表
func (linkedList *LinkedList) PrintLinkList() error {
	if linkedList.LinkedListIsEmpty() {
		return errors.New("linkedList is empty")
	}
	p := linkedList.headerNode.next
	for i := 1; i <= linkedList.linkedListLength; i++ {
		if p == nil {
			return errors.New("need to node is nil!")
		}
		fmt.Printf("nodeAddress:%p-node.data:%d-node.next_address:%p\n", p, p.data, p.next)
		p = p.next
	}
	return nil
}

//查找某个节点
func (linkedList *LinkedList) GetNode(nodeIndex int) (*linkedListNode, error) {
	if linkedList.LinkedListIsEmpty() || nodeIndex > linkedList.GetLinkedListLength() || nodeIndex <= 0 {
		return nil, errors.New("nodeIndex 不合法")
	}
	p := linkedList.headerNode.next
	currentNodeIndex := 1 //当前p所指的节点的下标
	for p != nil && currentNodeIndex < nodeIndex {
		p = p.next
		if currentNodeIndex != nodeIndex {
			currentNodeIndex++
		}
	}
	if p == nil || currentNodeIndex > nodeIndex {
		return nil, errors.New("没有找到这个节点")
	}
	return p, nil
}

//查找某个节点的数据域
func (linkedList *LinkedList) GetNodeData(nodeIndex int) (int, error) {
	node, err := linkedList.GetNode(nodeIndex)
	if err != nil {
		return 0, fmt.Errorf("getNodeData error: %s", err.Error())
	}
	return node.data, nil
}

//插入一个节点
func (linkedList *LinkedList) InsertNode(nodeData, insertIndex int) int {
	if linkedList.LinkedListIsEmpty() {
		return Fail
	}
	//1.插入位置前面一个节点
	var nodeBeforeInsertIndex *linkedListNode
	if insertIndex == 1 { //使用头结点来帮助插入
		nodeBeforeInsertIndex = linkedList.headerNode
	} else {

		var err error
		nodeBeforeInsertIndex, err = linkedList.GetNode(insertIndex - 1)
		if err != nil {
			fmt.Println("插入失败", err)
			return Fail
		}
	}
	//2.生成新的节点并插入到他之前 取代他
	newNode := &linkedListNode{nodeData, nodeBeforeInsertIndex.next}
	nodeBeforeInsertIndex.next = newNode
	linkedList.linkedListLength++
	return Success
}

//在链表末尾增加一个节点
func (l *LinkedList) AddNodeInLikedListEnd(nodeData int) {
	lastNode, err := l.GetNode(l.linkedListLength)
	fmt.Println(lastNode)
	if err != nil {
		fmt.Println("增加失败:", err)
		return
	}
	newNode := &linkedListNode{nodeData, nil}
	lastNode.next = newNode
	l.linkedListLength++
}

//删除某个节点
func (l *LinkedList) RemoveNode(nodeIndex int) {
	//找到这个节点 => 在找的是否判断是否nodeInddex合法
	needRemoveNode, err := l.GetNode(nodeIndex)
	if err != nil {
		fmt.Println("removeNode error:", err)
	}
	//找到他的前一个节点
	var pervNodeofNeedRemoveNode *linkedListNode
	if nodeIndex == 1 { //if nodeIndex then it's pervious node is LinkedList headerNode
		pervNodeofNeedRemoveNode = l.headerNode
	} else {
		pervNodeofNeedRemoveNode, err = l.GetNode(nodeIndex - 1)

	}
	if err != nil {
		fmt.Println("removeNode查找前一个节点error:", err)
	}
	//改变他们之间的关系
	pervNodeofNeedRemoveNode.next = needRemoveNode.next
	l.linkedListLength--
}

//change node data
func (l *LinkedList) ChangeNodeData(newData, nodeIndex int) {
	//find node by index
	node, err := l.GetNode(nodeIndex)
	if err != nil {
		fmt.Println("changeNodeData find a node err:", err)
	}
	node.data = newData
}
