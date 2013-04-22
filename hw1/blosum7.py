#!/usr/bin/python

import numpy as np
import itertools
import math
from collections import defaultdict

msaStr = """ACTCATGGCT
CATCATGGCA
CCTCCTGGCC
AATCAGGGAG
AATAATTGCT"""

AA = "ACGT"

msa = []
for row in msaStr.split('\n'):
    msa.append(list(row))

msa = np.array(msa)

n = msa.shape[1]
k = msa.shape[0]

q = defaultdict(int)
for row in msa:
    for col in row:
        q[col] += 1

for a in q:
    q[a] = float(q[a]) / (n * k)

f = defaultdict(int)
for pair in itertools.product(AA, repeat=2):
    for col in msa.T:
        for msaPair in itertools.product(col, repeat=2):
            if pair == msaPair:
                f[pair] += 1

score = {}
for pair in f:
    score[pair] = math.log((float(f[pair]) / (n * k * (k - 1))) / (q[pair[0]] * q[pair[1]]), 2)

print "MSA", msa

print "n", n
print "k", k

print "q", q
print "f", f

print "Score", score

def blast():
    for word in ('AAAA', 'AAAG', 'AAGC', 'AGCT', 'GCTT'):
        for rword in itertools.product(AA, repeat=4):
            s = 0
            for i in range(4):
                s += score[(word[i], rword[i])]
            if s > 4.5:
                print "".join(rword), s


blast()
