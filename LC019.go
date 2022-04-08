package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	// 首先先遍历链表确定长度
	prev := &ListNode{Next: head}
	cur := head
	cnt := 1
	for cur != nil {
		cnt++
		cur = cur.Next
	}
	// 找到要删除的元素的前一位元素
	findIndex := cnt - n - 1
	cur = prev
	i := 1
	for i < findIndex {
		cur = cur.Next
		i++
	}
	// 执行删除操作
	cur.Next = cur.Next.Next
	return prev.Next
}
