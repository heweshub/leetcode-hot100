# -*- coding: UTF-8 -*-
"""
author: cjhcw
"""

# 删除有序数组中的重复项


def removeDuplicates(nums):
    j = 1
    n = len(nums)
    for i in range(1, n):
        if nums[i] != nums[i - 1]:
            nums[j] = nums[i]
            j += 1
    return j


if __name__ == '__main__':
    nums = [0, 0, 1, 1, 1, 2, 2, 3, 3, 4]

    print(removeDuplicates(nums))
