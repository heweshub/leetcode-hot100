# -*- coding: UTF-8 -*-
"""
author: cjhcw
"""


def func(nums):
    i, j = 0, len(nums) - 1
    while i < j:
        if nums[i] < 0:
            i += 1
        if nums[j] > 0:
            j -= 1
        if nums[i] == 0 and nums[j] == 0:
            break

    return nums[i - 1], nums[j + 1]


if __name__ == '__main__':
    nums = [-9, -5, -2, 0, 0, 1, 3, 6, 8]
    print(func(nums))
