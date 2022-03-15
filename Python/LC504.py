# -*- coding: UTF-8 -*-
"""
author: cjhcw
"""


def convertToBase(num):
    if num == 0:
        return "0"
    ans = []
    flag = num < 0
    num = abs(num)
    while num > 0:
        ans.append(str(num % 7))
        num //= 7

    if flag:
        ans.append('-')
    return ''.join(reversed(ans))


if __name__ == '__main__':
    num = int(input("num="))
    print(convertToBase(num))
