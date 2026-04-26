package trees

func InvertTree(treeValues []int) []*int {
	if len(treeValues) == 0 {
		return []*int{}
	}
	root := &TreeNode{Val: treeValues[0]}
	for i := 1; i < len(treeValues); i++ {
		root.insertCompletedTree(treeValues[i])
	}
	invertTree(root)
	return root.print()
}

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		currentNode := queue[0]
		queue = queue[1:]
		if currentNode.Left != nil {
			queue = append(queue, currentNode.Left)
		}
		if currentNode.Right != nil {
			queue = append(queue, currentNode.Right)
		}
		currentNode.Left, currentNode.Right = currentNode.Right, currentNode.Left
	}
	return root
}
