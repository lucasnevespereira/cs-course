package main

import (
	"fmt"
)

type LinkedList struct {
	headNode *Node
}

func (l *LinkedList) getHeadNode() *Node {
	return l.headNode
}

func (l *LinkedList) insertBeginning(newValue string) {
	newNode := &Node{value: newValue}
	newNode.setNextNode(l.headNode)
	l.headNode = newNode
}

func (l *LinkedList) listValues() []string {
	var valueList []string
	const empty = ""
	currentNode := l.getHeadNode()

	for currentNode != nil {
		if currentNode.getValue() != empty {
			valueList = append(valueList, currentNode.getValue())
			currentNode = currentNode.getNextNode()
		}
	}

	return valueList
}

func (l *LinkedList) removeNode(valueToRemove string) {
	currentNode := l.getHeadNode()

	if currentNode.getValue() == valueToRemove {
		l.headNode = currentNode.getNextNode()
	} else {
		for currentNode != nil {
			nextNode := currentNode.getNextNode()
			if nextNode.getValue() == valueToRemove {
				currentNode.setNextNode(nextNode.getNextNode())
				currentNode = nil
			} else {
				currentNode = nextNode
			}
		}
	}
}

func main() {
	myNode := &Node{value: "5"}

	ll := &LinkedList{headNode: myNode}
	ll.insertBeginning("70")
	ll.insertBeginning("5675")
	ll.insertBeginning("90")

	fmt.Println(ll.listValues())
}
