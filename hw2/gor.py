#!/usr/bin/python

import sys

import numpy as np

alpha_data = np.loadtxt("gor-alpha.txt")	
beta_data = np.loadtxt("gor-beta.txt")	
turn_data = np.loadtxt("gor-turns.txt")	
coil_data = np.loadtxt("gor-coils.txt")	

aa_order = ["G", "A", "V", "L", "I", "S", "T", "D", "E", "N", "Q", "K", "H", "R", "F", "Y", "W", "C", "M", "P"]
aa_lookup = {}
for i, v in enumerate(aa_order):
	aa_lookup[v] = i

def get_value(aa, t, pos):
	pos += 8
	if t == "alpha":
		data = alpha_data
	elif t == "beta":
		data = beta_data
	elif t == "turn":
		data = turn_data
	elif t == "coil":
		data = coil_data
	return data[aa_lookup[aa], pos]

seq = sys.argv[1]
start = int(sys.argv[2])
end = int(sys.argv[3])

for pos1 in range(start, end + 1):
	aa1 = seq[pos1]
	maxScores = []
	for t in ("alpha", "beta", "turn", "coil"):
		score = 0
		for pos2 in range(pos1 - 8, pos1 + 9):
			aa2 = seq[pos2]
			offset = pos2 - pos1
			score += get_value(aa2, t, offset)
		maxScores.append((t, score))
	maxScores.sort(key=lambda x: x[1], reverse=True)
	print aa1, maxScores[0][0], maxScores[0][1]

