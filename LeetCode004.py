# -*- coding: UTF-8 -*-
"""
author: cjhcw
"""


def findMedianSortedArrays(nums1, nums2):
    # 设置nums1的长度默认比nums2的要短
    if len(nums1) > len(nums2):
        return findMedianSortedArrays(nums2, nums1)

    m, n = len(nums1), len(nums2)
    totalleft = (m + n + 1) // 2  # 找到最中间位置
    left, right = 0, m
    # 利用二分查找，找到两个数组的分界线
    # 分界线需要满足的条件是nums1[i-1] <= nums2[j] && nums2[j-1] <= nums1[i]
    while left < right:
        i = left + (right - left + 1) // 2
        j = totalleft - i

        if nums1[i - 1] > nums2[j]:
            right = i - 1
        else:
            left = i
    # 找到分界线之后
    i = left
    j = totalleft - i
    maxInteger = 10 ** 10
    minInteger = -10 ** 10
    nums1leftMax = minInteger if i == 0 else nums1[i - 1]
    nums1rightMin = maxInteger if i == m else nums1[i]
    nums2leftMax = minInteger if j == 0 else nums2[j - 1]
    nums2rightMin = maxInteger if j == n else nums2[j]
    if (m + n) % 2 == 1:
        return max(nums1leftMax, nums2leftMax)
    else:
        return (max(nums1leftMax, nums2leftMax) + min(nums1rightMin, nums2rightMin)) / 2


if __name__ == '__main__':
    print(findMedianSortedArrays([1,3], [2]))
