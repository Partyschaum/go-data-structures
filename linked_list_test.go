package main

import "testing"

// setup

var firstNode = &Node{"first node", nil}
var secondNode = &Node{"second node", nil}
var thirdNode = &Node{"third node", nil}

var nodeList = &NodeList{firstNode}

func resetNodes() {
	firstNode.next = nil
	secondNode.next = nil
	thirdNode.next = nil
}

// insertAfter

func TestNodeLinksToNewNode(t *testing.T) {
	resetNodes()

	nodeList.insertAfter(firstNode, secondNode)

	if firstNode.next != secondNode {
		t.Errorf(
			"Expected firstNode.next to be %+v, got %+v",
			secondNode,
			firstNode.next,
		)
	}
}

func TestNewNodeLinksToFormerNode(t *testing.T) {
	resetNodes()

	nodeList.insertAfter(firstNode, secondNode)
	nodeList.insertAfter(firstNode, thirdNode)

	if thirdNode.next != secondNode {
		t.Errorf(
			"Expected thirdNode.next to be %+v, got %+v",
			secondNode,
			thirdNode.next,
		)
	}
}

// insertBeginning

func TestListsFirstNodeLinksToNewNode(t *testing.T) {
	resetNodes()

	nodeList.insertBeginning(secondNode)

	if nodeList.firstNode != secondNode {
		t.Errorf(
			"Expected nodeList.firstNode to be %+v, got %+v",
			secondNode,
			nodeList.firstNode,
		)
	}
}

func TestNewNodeLinksToListsFormerFirstNode(t *testing.T) {
	resetNodes()

	nodeList = &NodeList{firstNode}
	nodeList.insertBeginning(secondNode)

	if nodeList.firstNode.next != firstNode {
		t.Errorf(
			"Expected nodeList.firstNode.next to be %+v, got %+v",
			firstNode,
			nodeList.firstNode.next,
		)
	}
}

// removeAfter

func TestNodeLinksToNodesNextNext(t *testing.T) {
	resetNodes()

	nodeList = &NodeList{firstNode}
	nodeList.insertAfter(firstNode, secondNode)
	nodeList.insertAfter(secondNode, thirdNode)

	nodeList.removeAfter(firstNode)

	if firstNode.next != thirdNode {
		t.Errorf(
			"Expected firstNode.next to be %+v, got %+v",
			thirdNode,
			firstNode.next,
		)
	}
}

func TestReturnsErrorIfNextNodeIsNil(t *testing.T) {
	resetNodes()

	nodeList = &NodeList{firstNode}

	err := nodeList.removeAfter(firstNode)

	if err == nil {
		t.Error("Expected error")
	}
}

// removeBeginning

func TestNodeListFirstNodeLinksToFirstNodeNext(t *testing.T) {
	resetNodes()

	nodeList = &NodeList{firstNode}
	nodeList.insertAfter(firstNode, secondNode)

	nodeList.removeBeginning()

	if nodeList.firstNode != secondNode {
		t.Errorf(
			"Expected nodeList.firstNode to be %+v, got %+v",
			secondNode,
			nodeList.firstNode,
		)
	}
}
