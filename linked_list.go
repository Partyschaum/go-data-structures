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
	currentNode := &Node{"first", nil}
	nodeList := &NodeList{currentNode}

	// add 100 nodes
	for i := 0; i < 100; i++ {
		newNode := &Node{"node" + strconv.Itoa(i), nil}
		nodeList.insertAfter(currentNode, newNode)
		fmt.Printf("Inserted %+v after %+v\n", newNode, currentNode)
		currentNode = newNode
	}

	currentNode = nodeList.firstNode

	// remove evey second node
	//for i := 0; i < 100; i++ {
	i := 0
	for currentNode != nil {
		if i%2 != 0 {
			continue
		}
		err := nodeList.removeAfter(currentNode)
		if err != nil {
			fmt.Println(err)
		}
		currentNode = currentNode.next
	}

	currentNode = nodeList.firstNode

	for currentNode != nil {
		fmt.Printf("%+v\n", currentNode)
		currentNode = currentNode.next
	}
}
