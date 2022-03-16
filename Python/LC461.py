# -*- coding: UTF-8 -*-
"""
author: cjhcw
"""
import collections


def hammingDistance(x: int, y: int) -> int:
    return collections.Counter(str(bin(x ^ y))[2:])['1']


if __name__ == '__main__':
    print(hammingDistance(1, 4))
