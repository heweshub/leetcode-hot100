# -*- coding: UTF-8 -*-
"""
author: cjhcw
"""


# 5. 最长回文子串
# 给你一个字符串 s，找到 s 中最长的回文子串

class Solution:
    # 寻找最长回文串，返回下标
    def expandAroundCenter(self, s, left, right):
        while left >= 0 and right < len(s) and s[left] == s[right]:
            left -= 1
            right += 1
        return left + 1, right - 1  # 结束循环前left和right又执行了一遍循环体

    def longestPalindrome(self, s: str) -> str:
        start, end = 0, 0
        for i in range(len(s)):
            # 回文字符串长度为奇数的情况
            left1, right1 = self.expandAroundCenter(s, i, i)
            # 回文字符串长度为偶数的情况
            left2, right2 = self.expandAroundCenter(s, i, i + 1)
            # 判断是否找到更长的回文字符串
            if right1 - left1 > end - start:
                start, end = left1, right1
            if right2 - left2 > end - start:
                start, end = left2, right2
        return s[start:end + 1]  # 注意字符串截取时的+1操作
