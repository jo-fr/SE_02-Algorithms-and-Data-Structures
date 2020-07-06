package datastructures

import (
	"fmt"
)

//Node data type
type Node struct {
	data int
	next *Node
}

// NewNode initializes new Node with given data
func NewNode(data int) *Node {
	n := &Node{}

	n.data = data

	fmt.Println("New Node with value:", n.data)
	return n
}

//Next sets the next Node
func (n *Node) Next(toNode *Node) {
	n.next = toNode
}

//GetNextNode returns the following Node
func (n *Node) GetNextNode() *Node {
	return n.next
}

//ShowLink shows the value of the next Node
func (n *Node) ShowLink() {
	fmt.Println(n.data, "->", n.next.data)
}

//CountNodes returns the number of Nodes starting with a given head
func CountNodes(head *Node) int {
	var countedNodes int
	currentNode := head
	for currentNode != nil {
		countedNodes++
		currentNode = currentNode.GetNextNode()
	}
	return countedNodes

}

//ShowCompleteList shows all links starting from a given head
func ShowCompleteList(head *Node) {
	currentNode := head

	for {

		fmt.Print(currentNode.data)
		currentNode = currentNode.GetNextNode()

		if currentNode == nil {
			fmt.Print("\n")
			break
		}

		fmt.Print("->")
	}
}
