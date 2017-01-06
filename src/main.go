package main

import (
	"./linkedList"
	"fmt"
)

func main() {
	linkedList := new(linkedList.LinkedList)
	linkedList.CreateLinkedList(10)
	fmt.Println(linkedList.HeaderNode)
}
