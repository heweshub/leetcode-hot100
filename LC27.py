# -*- coding: UTF-8 -*-
"""
author: cjhcw
"""


def removeElement(nums, val):
    j = 0
    n = len(nums)
    for i in range(n):
        if nums[i] != val:
            nums[j] = nums[i]
            j += 1
    return j


if __name__ == '__main__':
    nums = [0, 1, 2, 2, 3, 0, 4, 2]
    val = 2
    print(removeElement(nums, val))
