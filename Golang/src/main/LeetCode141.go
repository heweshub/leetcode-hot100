package main

// 141 Linked List Cycle

func hasCycle(head *ListNode) bool {
	// 快慢指针方法
	//fast, slow := head, head
	//for slow != nil && fast != nil {
	//	slow = slow.Next
	//	fast = fast.Next.Next
	//	if slow == fast{
	//		return true
	//	}
	//}
	//return false
	// 哈希集合
	isExistsMap := map[*ListNode]struct{}{}
	for head != nil {
		if _, ok := isExistsMap[head]; ok {
			return true
		}
		isExistsMap[head] = struct{}{}
		head = head.Next
	}
	return false
}
