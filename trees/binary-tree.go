package trees

import "fmt"

type BSTNode struct {
	value      int
	leftChild  *BSTNode
	RightChild *BSTNode
}

type BinarySearchTree struct {
	root *BSTNode
}

func (b *BinarySearchTree) New(value int) {
	newNode := &BSTNode{
	value:      value,
		RightChild: nil,
		leftChild:  nil,
	}
	b.root = newNode
}

func (b *BinarySearchTree) Insert(value int) {
	newNode := &BSTNode{
		value:      value,
		RightChild: nil,
		leftChild:  nil,
	}

	var parentNode *BSTNode

	currentNode := b.root

	for currentNode != nil {

		parentNode = currentNode

		if value > currentNode.value {

			currentNode = currentNode.RightChild

		} else if value < currentNode.value {

			currentNode = currentNode.leftChild

		} else if value == currentNode.value {

			fmt.Println("Node already exists")

			return
		}
	}
	if value > parentNode.value {

		parentNode.RightChild = newNode

	} else if value < parentNode.value {

		parentNode.leftChild = newNode

	}

}

func (b *BinarySearchTree) LookUp(value int) bool {
	currentNode := b.root
	if currentNode.value == b.root.value {
		return true
	}
	for currentNode != nil {
		if value > currentNode.value {

			currentNode = currentNode.RightChild

		} else if value < currentNode.value {

			currentNode = currentNode.leftChild
		} else if value == currentNode.value {
			return true
		}

	}
	return false
}
func (b *BinarySearchTree) Remove(value int) {
	currentNode := b.root
	var parentNode *BSTNode

	for currentNode != nil {
		if value < currentNode.value {
			parentNode = currentNode
			currentNode = currentNode.leftChild
		} else if value > currentNode.value {
			parentNode = currentNode
			currentNode = currentNode.RightChild
		} else {
			// Node with the value found

			 // Case 1: Node has no right child
			if currentNode.RightChild == nil {
				if parentNode == nil {
					// Deleting the root node
					b.root = currentNode.leftChild
				} else {
					// Reassign parent's child pointer
					if parentNode.leftChild == currentNode {
				 		parentNode.leftChild = currentNode.leftChild
					} else {
						parentNode.RightChild = currentNode.leftChild
					}
				}

				// Case 2: Right child has no left child
			} else if currentNode.RightChild.leftChild == nil {
				currentNode.RightChild.leftChild = currentNode.leftChild

				if parentNode == nil {
					// Deleting the root node
					b.root = currentNode.RightChild
				} else {
					// Reassign parent's child pointer
					if parentNode.leftChild == currentNode {
						parentNode.leftChild = currentNode.RightChild
					} else {
						parentNode.RightChild = currentNode.RightChild
					}
				}

				// Case 3: Right child has a left child (find in-order successor)
			} else {
				// Find the leftmost child of the right subtree (in-order successor)
				leftmost := currentNode.RightChild.leftChild
				leftmostParent := currentNode.RightChild

				// Move to the leftmost node
				for leftmost.leftChild != nil {
					leftmostParent = leftmost
					leftmost = leftmost.leftChild
				}

				// Replace leftmostParent's left with leftmost's right child
				leftmostParent.leftChild = leftmost.RightChild

				// Assign children to the leftmost node (successor)
				leftmost.leftChild = currentNode.leftChild
				leftmost.RightChild = currentNode.RightChild

				if parentNode == nil {
					// Deleting the root node
					b.root = leftmost
				} else {
					// Reassign parent's child pointer
					if parentNode.leftChild == currentNode {
						parentNode.leftChild = leftmost
					} else {
						parentNode.RightChild = leftmost
					}
				}
			}

			break // Node removed, exit the loop
		}
	}
}

func (b *BinarySearchTree) PrintTree() {
	printTreeHelper(b.root, "", true)
}

// Helper function to recursively print the tree.
func printTreeHelper(node *BSTNode, indent string, isLeft bool) {

	if node != nil {
		fmt.Print(indent)
		if isLeft {
			fmt.Print("L---")
			indent += "|   "
		} else {
			fmt.Print("R---")
			indent += "    "
		}
		fmt.Println(node.value)

		printTreeHelper(node.leftChild, indent, true)
		printTreeHelper(node.RightChild, indent, false)
	}
}
