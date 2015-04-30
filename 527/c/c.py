#!/usr/bin/env python
# -*- coding: utf-8 -*-

""" c.py


"""

import sys

def insert(arr, p):
    l = 0
    r = len(arr) - 1
    while l < r - 1:
        n = (l + r) // 2
        if arr[n] < p:
            l = n
        elif arr[n] > p:
            r = n
    return arr[:l+1] + [p] + arr[l+1:], arr[l+1], arr[l+2]
    
def main():
    sarr = sys.stdin.readline().strip().split(" ")
    w = int(sarr[0])
    h = int(sarr[1])
    n = int(sarr[2])

    xa = [0, w]
    ya = [0, h]
    
    for i in range(n):
        sarr = sys.stdin.readline().strip().split(" ")
        pos = int(sarr[1])
        if sarr[0] == "H":
            ya, y0, y1 = insert(ya, pos)
        else:
            xa, x0, x1 = insert(xa, pos)
            
if __name__ == "__main__":
    main()
