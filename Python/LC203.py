# -*- coding: UTF-8 -*-
"""
author: cjhcw
"""


def removeElements(head, val):
    """
    删除链表中元素值为val的节点
    :param head: 链表的第一个节点
    :param val: 整数val
    :return:
    """
    # 方法1 迭代
    # 设置头节点
    # prehead = ListNode(0,head)
    # cur = prehead
    # while cur.next:
    #     if cur.next.val == val:
    #         cur.next = cur.next.next
    #     else:
    #         cur = cur.next
    # return prehead.next

    # 方法2 递归
    # 终止条件
    if head is None:
        return head
    # 递归返回指向head.next
    head.next = removeElements(head.next, val)
    # 题目要求删除值为val的结点
    if head.val == val:
        return head.next
    return head
