# # # -*- coding: UTF-8 -*-
# # """
# # author: cjhcw
# # """
# #
# #
# # def evalRPN(tokens):
# #     operators = {
# #         '+': lambda x, y: x + y,
# #         '-': lambda x, y: x - y,
# #         '*': lambda x, y: x * y,
# #         '/': lambda x, y: int(x / y)
# #     }
# #     s = Stack()
# #     for i in tokens:
# #         try:
# #             num = int(i)
# #         except ValueError:
# #             num2 = s.pop()
# #             num1 = s.pop()
# #             num = operators[i](num1, num2)
# #         finally:
# #             s.push(num)
# #     return s.pop()
# #
# #
# # class Stack:
# #     def __init__(self):
# #         self.items = []
# #
# #     def push(self, x):
# #         self.items.append(x)
# #
# #     def pop(self):
# #         return self.items.pop()
# #
# #     def isEmpty(self):
# #         return self.items == []
# #
# #
# # def find_num(nums, key):
# #     l, r = 0, len(nums)
# #     while l < r:
# #         mid = (r + l) // 2
# #         if nums[mid] == key:
# #             return True
# #         # 判断key的位置
# #         if nums[0] <= nums[mid]:
# #             # 确定0-mid是否有序
# #             if nums[0] <= key < nums[mid]:
# #                 r = mid - 1
# #             else:
# #                 l = mid + 1
# #         else:
# #             # 判断mid-len(nums)-1是否有序
# #             if nums[mid] < key <= nums[-1]:
# #                 l = mid + 1
# #             else:
# #                 r = mid - 1
# #
# #     return False
# #
# #
# # if __name__ == '__main__':
# #     print(find_num([6, 7, 8, 1, 2, 3, 4, 5], 7))
#
#
# def deleteK(a, k):
#     n = len(a)
#     b = sorted(a)[:k]
#     s = ""
#     for i in a:
#         if i in b:
#             continue
#         s += str(i)
#
#     return int(s)
#
#
# if __name__ == '__main__':
#     print(deleteK([3, 1, 4, 5, 2], 3))

from collections import Counter


def addTwoNumber(arr, m):
    counter = Counter(arr)
    arr = list(counter.keys())
    print(arr)
    n = len(arr)
    l, r = 0, n - 1
    cnt = 0
    # print(counter)
    # print(counter.keys())
    while l < r:
        s = arr[l] + arr[r]
        if s == m:
            cnt += counter[arr[l]] * counter[arr[r]]
            l += 1
        elif s > m:
            r -= 1
        else:
            l += 1
    return cnt


if __name__ == '__main__':
    print(addTwoNumber([1, 1, 2, 3, 4, 6, 7, 7, 8, 9], 8))
