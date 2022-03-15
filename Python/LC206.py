# -*- coding: UTF-8 -*-
"""
author: cjhcw
"""


class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next


def reverseList(head):
    # 定义前驱节点
    pre = ListNode()
    cur = head
    while cur:
        # 保存原始链表的联式结构
        next = cur.next
        # 根据题目要求反转
        cur.next = pre
        pre = cur
        cur = next
    return pre


if __name__ == '__main__':
    nums = [1, 2, 3, 4, 5]
    cur = ListNode(nums[0])
    head = cur
    for i in range(1, len(nums)):
        node = ListNode(val=nums[i])
        cur.next = node
        cur = cur.next
    newHead = reverseList(head)
    cur = newHead
    while cur.next:
        print(cur.val, end=' ')
        cur = cur.next
