# -*- coding: UTF-8 -*-
"""
author: cjhcw
"""


def maopaoSort(nums):
    """
    冒泡排序
    :param nums: 无序数组
    :return: 每次将前后元素进行对比，一次循环将最大或最小的元素放置在最终额位置
    """
    n = len(nums)
    for i in range(n - 1, -1, -1):
        for j in range(i):
            if nums[j] > nums[j + 1]:
                temp = nums[j]
                nums[j] = nums[j + 1]
                nums[j + 1] = temp
    return nums


def kuaisuSort(nums, left, right):
    """
    快速排序
    :param nums: 无序数组
    :param left: 需要排序的左边界
    :param right: 需要排序的右边界
    :return: 递归方法实现，直接在原始数组上进行操作
    """

    if left >= right:
        return
    l, r = left, right
    pivot = nums[l]
    while l < r:
        while l < r and nums[r] >= pivot:
            r -= 1
        if l < r:
            nums[l] = nums[r]
        while l < r and nums[l] <= pivot:
            l += 1
        if l < r:
            nums[r] = nums[l]
        if l >= r:
            nums[l] = pivot
    kuaisuSort(nums, left, l - 1)
    kuaisuSort(nums, l + 1, right)


def kuaisuSort2(nums):
    n = len(nums)
    if n <= 1:
        return nums

    pivot = nums[0]

    left = [i for i in nums[1:] if i <= pivot]
    right = [i for i in nums[1:] if i > pivot]

    return kuaisuSort2(left) + [pivot] + kuaisuSort2(right)


def xuanzeSort(nums):
    """
    选择排序
    顾名思义，选择最小或者最大的数放在数组对应的位置
    :param nums: 无序数组
    :return: 有序数组
    """
    n = len(nums)

    for i in range(n):
        min_index = i
        for j in range(i + 1, n):
            if nums[j] < nums[min_index]:
                min_index = j
        temp = nums[i]
        nums[i] = nums[min_index]
        nums[min_index] = temp
    return nums


import heapq


def heapSort(nums):
    """
    用到了python自带的heapq小根堆
    :param nums: 无序数组
    :return: 有序数组
    """
    q = [(nums[i], i) for i in range(len(nums))]
    heapq.heapify(q)
    ans = []

    while len(q) > 0:
        ans.append(q[0][0])
        heapq.heappop(q)
    return ans


def mergeSort(nums):
    """
    归并排序
    :param nums: 无序数组
    :return:
    """
    n = len(nums)
    if n <= 1:
        return nums

    mid = n // 2
    # 有点二分的意思
    left = mergeSort(nums[:mid])
    right = mergeSort(nums[mid:])
    # 然后再合并
    return merge(left, right)


def merge(left, right):
    l_index, r_index = 0, 0
    merged = []

    while l_index < len(left) and r_index < len(right):
        # 双指针，选择较小的值放在输出数组中
        if left[l_index] <= right[r_index]:
            merged.append(left[l_index])
            l_index += 1
        else:
            merged.append(right[r_index])
            r_index += 1
    # 加上原先两个数组中未比较的数
    merged = merged + left[l_index:]
    merged = merged + right[r_index:]
    return merged


if __name__ == '__main__':
    # nums = [1, 2, 4, 5, 3, 6]
    input_str = input()
    nums = input_str.split(" ")
    print(nums)
    # print(maopaoSort(nums))
    # kuaisuSort(nums, 0, len(nums) - 1)
    # print(nums)
    # print(kuaisuSort2(nums))
    # print(xuanzeSort(nums))
    # print(heapSort(nums))
    # print(mergeSort(nums))
