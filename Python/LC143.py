# -*- coding: UTF-8 -*-
"""
author: cjhcw
"""
# 143 重排链表

# 1 - 2 - 3 - 4 - 5 - 6 - 7 - 8
# 1 - 8 - 2 - 7 - 3 - 6 - 4 - 5


def reorderList(head):
    """
    Do not return anything, modify head in-place instead.
    """
    if not head:
        return
    cur = head
    vec = []
    # 用数组来按序保存链表中的节点
    while cur:
        vec.append(cur)
        cur = cur.next
    # 双指针强行构建新的链表
    i, j = 0, len(vec) - 1
    while i < j:
        vec[i].next = vec[j]
        i += 1
        if i == j:
            break
        vec[j].next = vec[i]
        j -= 1
    # 需要注意的是链表的结尾指向空
    vec[i].next = None
