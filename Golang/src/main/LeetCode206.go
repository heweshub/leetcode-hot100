package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	// 定义前驱节点
	var pre *ListNode
	cur := head
	for cur != nil {
		// 找到当前节点的下一个指针
		next := cur.Next
		// 按照题目要求，当前指针需要指向前一个节点
		cur.Next = pre
		// 当前节点的下一个指针的前驱节点就是它本身
		pre = cur
		// 指针继续指向下一个节点，重复指向循环体内的操作
		cur = next
	}
	return pre
}

//func reverseList(head *ListNode) *ListNode {
//	if head == nil || head.Next == nil {
//	return head
//	}
//	newHead := reverseList(head.Next)
//	// 把head->Next看作一个整体
//	head.Next.Next = head
//	// 将原先正向的指针置为空
//	head.Next = nil
//	return newHead
//}
