#!/usr/bin/env python3
import random
import sys

with open("test.csv", "w") as f:
    f.write(','.join(['product_id', 'price', 'stock']))
    f.write('\n')
    n = int(sys.argv[1])
    for i in range(n):
        f.write(','.join(list(map(str, [i+1, random.randint(1, 100), random.randint(1, 100)]))))
        f.write('\n')
