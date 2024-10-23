/*
Given the root of a binary tree, replace the value of each node in the tree with the sum of all its cousins' values.

Two nodes of a binary tree are cousins if they have the same depth with different parents.

Return the root of the modified tree.

Note that the depth of a node is the number of edges in the path from the root node to it.



Example 1:

Input: root = [5,4,9,1,10,null,7]
Output: [0,0,0,7,7,null,11]
Explanation: The diagram above shows the initial binary tree and the binary tree after changing the value of each node.
- Node with value 5 does not have any cousins so its sum is 0.
- Node with value 4 does not have any cousins so its sum is 0.
- Node with value 9 does not have any cousins so its sum is 0.
- Node with value 1 has a cousin with value 7 so its sum is 7.
- Node with value 10 has a cousin with value 7 so its sum is 7.
- Node with value 7 has cousins with values 1 and 10 so its sum is 11.

Example 2:

Input: root = [3,1,2]
Output: [0,0,0]
Explanation: The diagram above shows the initial binary tree and the binary tree after changing the value of each node.
- Node with value 3 does not have any cousins so its sum is 0.
- Node with value 1 does not have any cousins so its sum is 0.
- Node with value 2 does not have any cousins so its sum is 0.

*/

package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func replaceValueInTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	// 1. get the depth of the tree
	depth := getDepth(root)

	// 2. get the parent of each node
	parents := make(map[int]*TreeNode)
	getParents(root, parents, nil)

	// 3. get the cousins of each node
	cousins := make(map[int][]int)
	getCousins(root, parents, cousins, depth)

	// 4. replace the value of each node with the sum of all its cousins' values
	replaceValue(root, cousins)

	return root
}

func getDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	left := getDepth(root.Left)
	right := getDepth(root.Right)

	if left > right {
		return left + 1
	}

	return right + 1
}

func getParents(root *TreeNode, parents map[int]*TreeNode, parent *TreeNode) {
	if root == nil {
		return
	}

	parents[root.Val] = parent

	getParents(root.Left, parents, root)
	getParents(root.Right, parents, root)
}

func getCousins(root *TreeNode, parents map[int]*TreeNode, cousins map[int][]int, depth int) {
	if root == nil {
		return
	}

	if root.Left != nil {
		if root.Left.Left != nil {
			cousins[root.Left.Val] = append(cousins[root.Left.Val], root.Left.Left.Val)
		}

		if root.Left.Right != nil {
			cousins[root.Left.Val] = append(cousins[root.Left.Val], root.Left.Right.Val)
		}
	}

	if root.Right != nil {
		if root.Right.Left != nil {
			cousins[root.Right.Val] = append(cousins[root.Right.Val], root.Right.Left.Val)
		}

		if root.Right.Right != nil {
			cousins[root.Right.Val] = append(cousins[root.Right.Val], root.Right.Right.Val)
		}
	}

	getCousins(root.Left, parents, cousins, depth)
	getCousins(root.Right, parents, cousins, depth)
}

func replaceValue(root *TreeNode, cousins map[int][]int) {
	if root == nil {
		return
	}

	sum := 0
	for _, cousin := range cousins[root.Val] {
		sum += cousin
	}

	root.Val = sum

	replaceValue(root.Left, cousins)
	replaceValue(root.Right, cousins)
}
