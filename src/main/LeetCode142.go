package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// 哈希表实现
//func detectCycle(head *ListNode) *ListNode {
//	isExistsMap := map[*ListNode]int{}
//	for head!=nil{
//		if _,ok := isExistsMap[head];ok{
//			return head
//		}
//		isExistsMap[head]++
//		head = head.Next
//	}
//	return nil
//}
// 快慢指针
func detectCycle(head *ListNode) *ListNode {
	fast, slow := head, head
	for fast != nil {
		slow = slow.Next
		if fast.Next == nil {
			return nil
		}
		fast = fast.Next.Next
		// 慢指针还需要走从首节点到入环点之间的距离才能到入环的第一个节点
		//  具体入环前的节点a，入环到环中相遇点b，相遇点再到入环点c
		// fast f=a+b+c+b
		// slow s=a+b
		// f=2s => a=c
		// 然后p从head出发，slow从相遇点出发，p和slow正好相遇在入环点
		if fast == slow {
			p := head
			for p != slow {
				p = p.Next
				slow = slow.Next
			}
			return p
		}
	}
	return nil
}
