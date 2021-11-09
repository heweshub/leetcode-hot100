# -*- coding: UTF-8 -*-
"""
author: cjhcw
"""


# 利用哈希集合，将每个字符依次放入set中，如果发现有重复字符，
# 将第一次出现的重复字符移出set，当前字符串长度-1，起始的字符左边也需要往右移动
# （即left+=1），再判断当前字符串长度与最大字符串长度的大小，
# set中加入当前第二次出现重复的字符，最终返回最长字符串长度。


class Solution:
    def lengthOfLongestSubstring(self, s: str) -> int:
        if not s:
            return 0
        left = 0
        lookup = set()
        n = len(s)
        max_len = 0
        cur_len = 0
        for i in range(n):
            cur_len += 1
            while s[i] in lookup:
                lookup.remove(s[left])
                left += 1
                cur_len -= 1
            if cur_len > max_len:
                max_len = cur_len
            lookup.add(s[i])
        return max_len
