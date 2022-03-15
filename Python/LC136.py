# -*- coding: UTF-8 -*-
"""
author: cjhcw
"""


def singleNumber(nums):
    seen = set()
    for i in range(len(nums)):
        if nums[i] not in seen:
            seen.add(nums[i])
        elif nums[i] in seen:
            seen.remove(nums[i])
    return seen.pop()


if __name__ == '__main__':
    nums = [4, 1, 2, 1, 2]
    print(singleNumber(nums))
