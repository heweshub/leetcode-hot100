# -*- coding: UTF-8 -*-
"""
author: cjhcw
"""


def maxProfit(ps):
    """
    求利润最大
    :param ps: 每日股票价格数组
    :return: 最大利润
    """
    max_pro = 0
    min_pri = int(1e9)
    for pric in ps:
        max_pro = max(pric - min_pri, max_pro)
        min_pri = min(pric, min_pri)
    return max_pro


if __name__ == '__main__':
    print(maxProfit([7, 1, 5, 3, 6, 4]))
