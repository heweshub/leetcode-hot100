# -*- coding: UTF-8 -*-
"""
author: cjhcw
"""


def moveZeroes(nums):
    """
    Do not return anything, modify nums in-place instead.
    """
    n = len(nums)
    l, r = 0, 0
    # nums[r]不等于0的时候l才会向右移动
    while r < n:
        if nums[r] != 0:
            nums[l], nums[r] = nums[r], nums[l]
            l += 1
        r += 1
    print(nums)


if __name__ == '__main__':
    nums = [0, 0, 1]
    moveZeroes(nums)
