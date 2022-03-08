# -*- coding: UTF-8 -*-
"""
author: cjhcw
"""


def romanToInt(s):
    dic = {"I": 1, "V": 5, "X": 10, "L": 50, "C": 100, "D": 500, "M": 1000}
    ans = 0
    # 确定遇到的字符的级别
    highlevel = 1
    for i in s[::-1]:
        level = dic[i]
        # 大于当前最高级别的就加上
        if level >= highlevel:
            ans += level
            highlevel = level
        # 小于当前级别的就减去
        else:
            ans -= level
    return ans


if __name__ == '__main__':
    print(romanToInt("IV"))
