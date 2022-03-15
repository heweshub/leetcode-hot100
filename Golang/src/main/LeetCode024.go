package main

//func swapPairs(head *ListNode) *ListNode {
//	// 创建一个超前指针指向
//	preHead := &ListNode{0, head}
//	temp := preHead
//	for temp.Next != nil && temp.Next.Next != nil {
//		node1 := temp.Next
//		node2 := temp.Next.Next
//		//到目前为止就是temp->node1->node2,下面就是要调整为temp->node2->node1这样的顺序
//		temp.Next = node2
//		node1.Next = node2.Next
//		node2.Next = node1
//		// 最后temp重新归位
//		temp = node1
//	}
//	return preHead.Next
//}

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	// 新的头节点是原先链表的第二个节点
	newHead := head.Next
	// 这是把原先第二个节点后的链表看作一个整体
	head.Next = swapPairs(newHead.Next)
	// 调换两个节点的顺序
	newHead.Next = head
	// 返回新的头节点
	return newHead
}
