package linkedList

import (
	"fmt"
	"github.com/pkg/errors"
)

//node type
type twoWay_node struct {
	data           int
	pervious, next *twoWay_node
}

//twoWay linkedList
type TwoWay struct {
	twoWay_linkedList_lenth int
	headerNode              *twoWay_node
}

//init
func (twoway_linkedList *TwoWay) InitTwoWay_linkedList(nodeData []int) error {
	length := len(nodeData)
	headerNode := &twoWay_node{0, nil, nil}
	twoway_linkedList.headerNode = headerNode
	twoway_linkedList.twoWay_linkedList_lenth = 0

	perviousNode := twoway_linkedList.headerNode
	var newNode *twoWay_node
	for i := 0; i < length; i++ {
		newNode = &twoWay_node{data: nodeData[i], pervious: perviousNode, next: nil}
		perviousNode.next = newNode
		perviousNode = newNode
	}
	twoway_linkedList.headerNode.pervious = newNode
	newNode.next = twoway_linkedList.headerNode
	twoway_linkedList.twoWay_linkedList_lenth = length
	return nil
}

// is Empty
func (twoway_linkedList *TwoWay) IsEmpty() bool {
	if twoway_linkedList.twoWay_linkedList_lenth == 0 {
		return true
	}
	return false
}

//length
func (twoway_linkedList *TwoWay) GetLinkedListLength() int {
	return twoway_linkedList.twoWay_linkedList_lenth
}

//print
func (twoway_linkedList *TwoWay) PrintLinkedList() {
	if twoway_linkedList.IsEmpty() {
		return
	}
	node := twoway_linkedList.headerNode
	for i := 1; i <= twoway_linkedList.GetLinkedListLength()+1; i++ { //把头结点也打印出来
		fmt.Printf("nodeAddress:%p====pervirous:%p====next:%p===data:%d\n", node, node.pervious, node.next, node.data)
		node = node.next
	}
}

//add
//在 insertIndex 位置上插入一个节点
func (twoway_linkedList *TwoWay) InsertNode(data, insertIndex int) error{
	var per,ne *twoWay_node
	if insertIndex == 1 {
		per = twoway_linkedList.headerNode
		ne =  twoway_linkedList.headerNode.next
	}else {

		//查找insertIndex 前一个节点
		var err error
		per, err = twoway_linkedList.FindNode(insertIndex - 1)
		if err != nil {
			return errors.New(fmt.Sprintf("插入失败:",err.Error()))
		}
		ne = per.next
	}
	newNode := &twoWay_node{data:data,pervious:per,next:ne}
	per.next = newNode
	ne.pervious = newNode
	twoway_linkedList.twoWay_linkedList_lenth++
	return nil
}

//remove
func (twoway_linkedList *TwoWay)RemoveNode(nodeIndex int) error{
	node,err := twoway_linkedList.FindNode(nodeIndex)
	if err != nil{
		return errors.New(fmt.Sprintf("remove fail:",err.Error()))
	}
	if nodeIndex == 1 {
		twoway_linkedList.headerNode.next =  node.next
		node.next.pervious = twoway_linkedList.headerNode
	} else {
		node.pervious.next = node.next
		node.next.pervious = node.pervious
	}
	twoway_linkedList.twoWay_linkedList_lenth--
	return nil
}
//change
func (twoWay_linkedList *TwoWay)ChangeNodeData(data,nodeIndex int) error{
	node,err := twoWay_linkedList.FindNode(nodeIndex)
	if err != nil {
		return errors.New(fmt.Sprintf("change fail:",err.Error()))
	}
	node.data = data
	return nil
}

//find
func (twoway_linkedList *TwoWay) FindNode(nodeIndex int) (*twoWay_node, error) {
	//判断位置是否超出当前数组的长度

	if twoway_linkedList.IsEmpty() || nodeIndex > twoway_linkedList.twoWay_linkedList_lenth || nodeIndex <= 0 {
		return nil, errors.New("下标越界")
	}

	p := twoway_linkedList.headerNode.next
	i := 1

	for i < nodeIndex {
		p = p.next
		i++
	}
	if  i > nodeIndex {
		return nil, errors.New("查找失败")
	}
	return p, nil
}

