# -*- coding: UTF-8 -*-
"""
author: cjhcw
"""


def isPalindrome(x):
    if x < 0:
        return False
    # nums = [i for i in str(x)]
    # n = len(nums)
    # for i in range(n // 2 + 1):
    #     if nums[i] != nums[n - i - 1]:
    #         return False
    # return True

    reverted_number = 0
    while x > reverted_number:
        reverted_number = reverted_number * 10 + x % 10
        x //= 10
    return x == reverted_number or x == reverted_number // 10


if __name__ == '__main__':
    print(isPalindrome(121))
