package trees

type KthTreeNode struct {
	Val   int
	Left  *KthTreeNode
	Right *KthTreeNode
}

func (k *KthTreeNode) insert(v int) {
	if v == 0 {
		return
	}
	if v < k.Val {
		if k.Left == nil {
			k.Left = &KthTreeNode{Val: v}
		} else {
			k.Left.insert(v)
		}
	} else if v > k.Val {
		if k.Right == nil {
			k.Right = &KthTreeNode{Val: v}
		} else {
			k.Right.insert(v)
		}
	} else {
		k.Left.insert(v)
	}
}

// root = [2,1,3], k = 1
// root = [4,3,5,2,null], k = 4
func kthSmallest(root *KthTreeNode, k int) int {
	// เราใช้ Stack (LIFO) เพื่อเก็บทางผ่านที่เราเดินลงไปทางซ้าย
	stack := []*KthTreeNode{}
	curr := root
	count := 0

	for curr != nil || len(stack) > 0 {
		// 1. "ดิ่งซ้าย" ให้สุดทาง: เพราะตัวที่น้อยที่สุดอยู่ซ้ายสุดเสมอ
		for curr != nil {
			stack = append(stack, curr) // เก็บทางผ่านไว้ใน Stack
			curr = curr.Left
		}

		// 2. เมื่อซ้ายสุดแล้ว "ถอยกลับมา" (Pop)
		curr = stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		// 3. "นับแต้ม" ทุกครั้งที่ถอยกลับมาที่ตัวกลาง
		count++
		if count == k {
			return curr.Val // เจอตัวที่ k แล้ว! 🎯
		}

		// 4. "เลี้ยวขวา" เพื่อไปสำรวจค่าที่มากกว่าต่อไป
		curr = curr.Right
	}

	return -1 // เคสหาไม่เจอ
}

func KthSmallest(nums []int, k int) int {
	root := &KthTreeNode{
		Val: nums[0],
	}
	for _, n := range nums[1:] {
		root.insert(n)
	}
	return kthSmallest(root, k)
}
