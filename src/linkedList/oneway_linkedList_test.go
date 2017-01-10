package linkedList

import (
	"testing"
)

var list LinkedList = LinkedList{nil, 0}

//TEST CreateLinkedList
func TestLinkedList_CreateLinkedList(t *testing.T) {
	for i := -10; i <= 10; i++ {
		err := list.CreateLinkedList(i)
		if err != nil && err.Error() != "intput length incorrect" {
			t.Errorf("TestLinkedList_CreateLinkedList test %d fail:%s", i, err)
		}
	}
}

//TEST PrintLinkedList
func TestLinkedList_PrintLinkList(t *testing.T) {
	err := list.PrintLinkList()
	if err != nil && (err.Error() != "need to node is nil!" || err.Error() != "linkedList is empty") {
		t.Error("test PrintLikedList()  fail:", err)
	}
}

// find a node
func TestLinkedList_GetNode(t *testing.T) {
	for i := -10; i < 20; i++ { //condition
		_, err := list.GetNode(i)
		if err != nil {
			t.Error("test GetNode()  fail:", i, err.Error())
		}
	}
}

//TEST find a node data domain
func TestLinkedList_GetNodeData(t *testing.T) {
	for i := -10; i <= 10; i++ {
		_, err := list.GetNodeData(i)
		if err != nil {
			t.Error("test GetNodedata fail:", err)
		}
	}
}

//test insert
