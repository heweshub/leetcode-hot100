import cv2
import os
import numpy as np
from matplotlib import pyplot as plt

if __name__ == '__main__':
    xueyuan_dir = 'faxueyuanxianchang/'
    gaokao_dir = 'faxue/'
    xueyuan_list = os.listdir(xueyuan_dir)
    gaokao_list = os.listdir(gaokao_dir)
    count_zai = 0
    count_buzai = 0
    cnt = 1
    for path in gaokao_list:  # path = JPG     path = jpg
        path1 = path.replace('.jpg', '.JPG')
        id = path.split('.')[0]
        # print("the %d photo" % cnt)
        # cnt += 1
        if path1 in xueyuan_list:
            # gaokaoname = path1
            gaokaoname = gaokao_dir + path1
            # xueyuanname = path
            xueyuanname = xueyuan_dir + path
            # print(gaokaoname)
            # print(xueyuanname)
            gaokao_ = cv2.imread(gaokaoname)
            xueyuan_ = cv2.imread(xueyuanname)
            gaokao_ = cv2.cvtColor(gaokao_, cv2.COLOR_BGR2RGB)
            xueyuan_ = cv2.cvtColor(xueyuan_, cv2.COLOR_BGR2RGB)

            plt.subplot(2, 2, 1)
            plt.imshow(gaokao_, cmap='gray')
            title = 'gaokao' + id
            plt.title(title)
            plt.xticks([])
            plt.yticks([])

            plt.subplot(2, 2, 2)
            plt.imshow(xueyuan_, cmap='gray')
            title = 'xueyuan' + id
            plt.title(title)
            plt.xticks([])
            plt.yticks([])

            # plt.show()
            count_zai += 1
        else:
            count_buzai += 1
            print(id)
    print('zai:', count_zai)
    print('buzai:', count_buzai)
    # print(xueyuan_list)
