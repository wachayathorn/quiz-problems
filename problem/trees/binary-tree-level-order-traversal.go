package trees

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func (t *TreeNode) insert(v int) {
	// Insert in level-order (complete tree insert):
	// find first node that has a missing child (left first, then right)
	queue := []*TreeNode{t}
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		if cur.Left == nil {
			cur.Left = &TreeNode{Val: v}
			return
		}
		queue = append(queue, cur.Left)

		if cur.Right == nil {
			cur.Right = &TreeNode{Val: v}
			return
		}
		queue = append(queue, cur.Right)
	}
}

func levelOrder(root *TreeNode) [][]int {
	// ถ้าไม่มี Root ก็คืนค่าว่างไปเลย (Edge case)
	if root == nil {
		return [][]int{}
	}

	result := [][]int{}
	// 2. สร้าง Queue (แถวคอย) และเอา Root ใส่ลงไปเป็นคนแรก
	queue := []*TreeNode{root}

	// 3. เริ่ม "วนลูปหลัก" (ตราบใดที่ยังมีคนรออยู่ในคิว)
	for len(queue) > 0 {
		// --- จุดสำคัญ: ถ่ายรูป (Snapshot) จำนวนคนในคิวตอนนี้ไว้ก่อน ---
		queueSize := len(queue)
		currentLevel := []int{}

		// 4. "วนลูปลูก" เพื่อจัดการคนที่มีอยู่ในชั้นนี้เท่านั้น
		for i := 0; i < queueSize; i++ {
			// ดึง Node หน้าสุดออกจากคิว (Dequeue)
			currentNode := queue[0]
			queue = queue[1:]

			// เก็บค่าของ Node นี้ใส่ List ของชั้นปัจจุบัน
			currentLevel = append(currentLevel, currentNode.Val)

			// 5. "ส่งไม้ต่อ" ให้ลูกๆ (ถ้ามี) ไปต่อแถวรอที่ท้ายคิว
			if currentNode.Left != nil {
				queue = append(queue, currentNode.Left)
			}
			if currentNode.Right != nil {
				queue = append(queue, currentNode.Right)
			}
		}

		// จบการทำงานของชั้นนี้ เอาผลลัพธ์ใส่ตระกร้าใหญ่
		result = append(result, currentLevel)
	}

	return result
}

func LevelOrder(treeValues []int) [][]int {
	if len(treeValues) == 0 {
		return [][]int{}
	}
	root := &TreeNode{Val: treeValues[0]}
	for i := 1; i < len(treeValues); i++ {
		root.insert(treeValues[i])
	}
	return levelOrder(root)
}
