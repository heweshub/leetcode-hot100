# -*- coding: UTF-8 -*-
"""
author: cjhcw
"""


# 两个非负整数按倒序存储在两个链表中，题目要求两数之和也按倒序存储在链表中。
#
# 以上就是题目，下面是求解思路：
#
# 输入是倒序，输出要求的也是倒序，那就直接在初始的两个链表上操作，
# 即对应位置上的元素相加，这里需要主要的是进位，可以用一个变量存储，
# 很重要的一点是两个的位数可能不一致，下面是求解的代码：


class Solution:
    def addTwoNumbers(self, l1: ListNode, L2: ListNode) -> ListNode:
        result = current = ListNode()
        # 进位标志
        reminder = 0
        while l1 or l2:
            x = l1.val if l1 else 0
            y = l2.val if l2 else 0

            sum = x + y + reminder
            current.next = ListNode(sum % 10)
            reminder = sum // 10

            if l1:
                l1 = l1.next
            if l2:
                l2 = l2.next

            current = current.next
        if reminder:
            current.next = ListNode(reminder)
        return result.next
