package main

import (
	"errors"
	"fmt"
	"strconv"
)

type Node struct {
	data string
	next *Node
}

type NodeList struct {
	firstNode *Node
}

func (nodeList *NodeList) insertAfter(node, newNode *Node) error {
	newNode.next = node.next
	node.next = newNode
	return nil
}

func (nodeList *NodeList) insertBeginning(newNode *Node) error {
	newNode.next = nodeList.firstNode
	nodeList.firstNode = newNode
	return nil
}

func (nodeList *NodeList) removeAfter(node *Node) error {
	if node.next == nil {
		return errors.New("End of list")
	}
	node.next = node.next.next
	return nil
}

func (nodeList *NodeList) removeBeginning() error {
	nodeList.firstNode = nodeList.firstNode.next
	return nil
}

func main() {
	nodeList := &NodeList{&Node{"first", nil}}
	nodeList.insertAfter(
		nodeList.firstNode,
		&Node{"second", nil},
	)

	fmt.Printf("nodeList: %+v\n", nodeList)
	fmt.Printf("nodeList.firstNode: %+v\n", nodeList.firstNode)
	fmt.Printf("nodeList.firstNode.next: %+v\n", nodeList.firstNode.next)

	nodeList.insertBeginning(&Node{"third", nil})

	fmt.Printf("nodeList: %+v\n", nodeList)
	fmt.Printf("nodeList.firstNode: %+v\n", nodeList.firstNode)
	fmt.Printf("nodeList.firstNode.next: %+v\n", nodeList.firstNode.next)
	fmt.Printf("nodeList.firstNode.next.next: %+v\n", nodeList.firstNode.next.next)
}
