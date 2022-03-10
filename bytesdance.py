# -*- coding: UTF-8 -*-
"""
author: cjhcw
"""


# 1 实现字符串转换成整数的函数:

# def str2number(s):
#     if not s:
#         return s
#     flag = 0
#     if s[0] == '-':
#         flag = s[0]
#         s = s[1:]
#     ans = 0
#     for i in range(len(s)):
#         ans = ans * 10 + int(s[i])
#     if flag:
#         ans = -ans
#     return ans
#
#
# if __name__ == '__main__':
#     s = "169"
#     print(str2number(s))


# 2 合并两个有序链表

# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
#
#
# def mergeTwoLists(l1, l2):
#     prehead = ListNode(0, None)
#     cur = prehead
#     while l1 and l2:
#         if l1.val <= l2.val:
#             cur.next = l1
#             l1 = l1.next
#         else:
#             cur.next = l2
#             l2 = l2.next
#         cur = cur.next
#     cur.next = l1 if l1 is not None else l2
#     return prehead.next
#
#
# if __name__ == '__main__':
#     s1 = input()
#     s2 = input()
#     # print(s1, s2)
#     ls1 = s1.split(" ")
#     ls2 = s2.split(" ")
#     # print(ls1, ls2)
#     head1 = ListNode(0, None)
#     head2 = ListNode(0, None)
#     cur1 = head1
#     cur2 = head2
#     for i in ls1:
#         cur1.next = ListNode(int(i), None)
#         cur1 = cur1.next
#     for j in ls2:
#         node = ListNode(int(j), None)
#         cur2.next = node
#         cur2 = cur2.next
#     newHead = mergeTwoLists(head1.next, head2.next)
#     while newHead:
#         print(newHead.val, end=' ')
#         newHead = newHead.next


# 3 输入n个数，找到其中最小的k个数目
# import heapq
#
#
# def findSmallest(nums, k):
#     n = len(nums)
#     # 建堆
#     q = [(nums[i], i) for i in range(n)]
#     ans = []
#     # 调整小顶堆的结构
#     heapq.heapify(q)
#     # 输出k个最小的数
#     for i in range(k):
#         ans.append(heapq.heappop(q))
#     return ans
#
#
# if __name__ == '__main__':
#     nums = [4, 5, 1, 6, 2, 7, 3, 8]
#     print(findSmallest(nums, 4))


# 4 找到两个链表的第一个公共结点

class ListNode:
    def __init__(self, m_nKey, m_pNext):
        self.m_nKey = m_nKey
        self.m_pNext = m_pNext


def findFirstNode(h1, h2):
    nodeList = {}
    cur1 = h1
    cur2 = h2
    while cur1:
        nodeList[cur1] = True
    while cur2:
        if nodeList.get(cur2):
            return cur2
    return None


# 5. 滑动窗口机制
import heapq


def maxSlidingWindow(nums, k):
    n = len(nums)
    q = [(-nums[i], i) for i in range(k)]

    heapq.heapify(q)
    ans = [-q[0][0]]
    for i in range(k, n):
        heapq.heappush(q, (-nums[i], i))
        while q[0][1] <= i - k:
            heapq.heappop(q)
        ans.append(-q[0][0])
    return ans


# 6.象棋中的马

# 7.缓冲区类

class Cache:
    def __init__(self):
        self.cache = []


# 8.LRU

import collections


# class LRUCache(collections.OrderedDict):
#
#     def __init__(self, capacity):
#         super().__init__()
#         self.capacity = capacity
#
#     def get(self, k: int):
#         if k not in self:
#             return -1
#         self.move_to_end(k)
#
#     def put(self, k, v):
#         if k in self:
#             self.move_to_end(k)
#         self[k] = v
#         if len(self) > self.capacity:
#             self.popitem(last=False)

class DLinkedNode:
    def __init__(self, key=0, value=0):
        self.key = key
        self.value = value
        self.prev = None
        self.next = None


class LRUCache:
    def __init__(self, cap):
        self.cache = dict()

        self.head = DLinkedNode()
        self.tail = DLinkedNode()
        self.head.next = self.tail
        self.tail.prev = self.head
        self.capacity = cap
        self.size = 0

    def get(self, key):
        if key not in self.cache:
            return -1
        node = self.cache[key]
        # 如果 key 存在，先通过哈希表定位，再移到头部
        self.moveToHead(node)
        return node.value

    def put(self, key, value):
        if key not in self.cache:
            node = DLinkedNode(key, value)
            self.cache[key] = node
            self.addToHead(node)

            self.size += 1
            if self.size > self.capacity:
                removed = self.removeTail()

                self.cache.pop(removed.key)
                self.size -= 1
        else:
            node = self.cache[key]
            node.value = value
            self.moveToHead(node)

    def addToHead(self, node):
        node.prev = self.head
        node.next = self.head.next
        self.head.next.prev = node
        self.head.next = node

    def removeNode(self, node):
        node.prev.next = node.next
        node.next.prev = node.prev

    def moveToHead(self, node):
        self.removeNode(node)
        self.addToHead(node)

    def removeTail(self):
        node = self.tail.prev
        self.removeNode(node)
        return node


# 9.LC143
def recorderList(head):
    if not head:
        return

    cur = head
    vec = []

    while cur:
        vec.append(cur)
        cur = cur.next

    i, j = 0, len(vec) - 1
    while i < j:
        vec[i].next = vec[j]
        i += 1
        if i == j:
            break
        vec[j].next = vec[i]
        j -= 1
    vec[i].next = None
