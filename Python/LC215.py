import random


def partition(a, l, r):
    pivot = a[r]
    rightmost = r
    while l <= r:
        while l <= r and a[l] > pivot:
            l += 1
        while l <= r and a[r] <= pivot:
            r -= 1
        if l <= r:
            a[l], a[r] = a[r], a[l]

    a[l], a[rightmost] = a[rightmost], a[l]
    return l


def randomPartition(a, l, r):
    i = random.randint(l, r)
    a[i], a[r] = a[r], a[i]
    return partition(a, l, r)


def quickSort(a, l, r, k):
    index = randomPartition(a, l, r)
    if index == k - 1:
        return a[k - 1]
    else:
        if index > k - 1:
            return quickSort(a, l, index - 1, k)
        else:
            return quickSort(a, index + 1, r, k)


def findKthLargest(nums, k):
    return quickSort(nums, 0, len(nums) - 1, k)
