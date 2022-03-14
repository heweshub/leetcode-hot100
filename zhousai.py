# -*- coding: UTF-8 -*-
"""
author: cjhcw
"""


def findKDistantIndices(nums, key: int, k: int):
    ans = []
    n = len(nums)
    index_key = []
    for i in range(n):
        if nums[i] == key:
            index_key.append(i)

    for j in index_key:
        ans.extend([m for m in range(max(0, j - k), min(n, j + k + 1))])
    ans = list(set(ans))
    ans.sort()
    return ans


if __name__ == '__main__':
    nums = [3, 4, 9, 1, 3, 9, 5]
    key = 9
    k = 1
    print(findKDistantIndices(nums, key, k))
