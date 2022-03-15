# -*- coding: UTF-8 -*-
"""
author: cjhcw
"""
import heapq


def maxSlidingWindow(nums, k):
    n = len(nums)
    # 创建大根堆，python默认是小跟堆
    q = [(-nums[i], i) for i in range(k)]
    heapq.heapify(q)

    # 现在就变成负的最多的数在堆顶
    ans = [-q[0][0]]
    for i in range(k, n):
        heapq.heappush(q, (-nums[i], i))
        # 循环把堆顶超过滑动窗口的元素移除
        while q[0][1] <= i - k:
            heapq.heappop(q)
        ans.append(-q[0][0])
    return ans


if __name__ == '__main__':
    nums = [9, 10, 9, -7, -4, -8, 2, -6]
    k = 5
    print(maxSlidingWindow(nums, k))
