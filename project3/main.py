#!/usr/bin/python

import random
import sys

import numpy as np

class Gibbs(object):
	def __init__(self, fn):
		with open(fn) as f:
			self.iterations = int(f.readline().strip())
			self.alphabet = f.readline().strip()
			self.background_probs = {k: v for (k, v) in zip(self.alphabet, [float(x) for x in f.readline().strip().split()])} # p
			self.length = int(f.readline().strip()) # W
			self.sequences = [] # S_i
			for line in f:
				self.sequences.append(line.strip())
		self.best_prob = float('-inf')
		self.best_q = None

	def rand_ai_start(self):
		aiList = []
		for seq in self.sequences:
			aiList.append(random.randint(0, len(seq) - self.length - 1))
		self.aiList = aiList
		return aiList

	def rand_si(self):
		exc = random.randint(0, len(self.sequences) - 1)
		return (self.sequences[exc], [i for i in range(len(self.sequences)) if i != exc])

	def calc_q(self, seqIndexes): # seqIndexes: list of indexes of remaining seqs
		q = np.empty((len(self.alphabet), self.length))
		q.fill(1.0 / len(self.alphabet))
		for i, b in enumerate(self.alphabet):
			for k in range(self.length):
				for seqIndex in seqIndexes:
					seq = self.sequences[seqIndex]
					pos = k + self.aiList[seqIndex]
					if pos < len(seq) and seq[pos] == b:
						q[i,k] += 1
		# normalize rows of q
		row_sums = q.sum(axis=1)
		qNormed = q / row_sums[:, np.newaxis]
		return qNormed

	def calc_ll_at_k(self, seq, q, k):
		factors = []
		assert k + self.length < len(seq)
		for j in range(k, k + self.length):
			numerator = q[self.alphabet.index(seq[j]), j - k]
			denom = self.background_probs[seq[j]]
			factors.append(numerator / denom)
		return np.prod(factors)

	def calc_ll_all(self, seq, q):
		return [self.calc_ll_at_k(seq, q, k) for k in range(len(seq) - self.length)]

	def new_ai_prob(self, seq, lls):
		lls_sum = sum(lls)
		seqIndex = self.sequences.index(seq)
		for k, ll in enumerate(lls):
			p = ll / lls_sum
			if random.random() <= p:
				self.aiList[seqIndex] = k

	def update_best(self, q):
		probs = []
		for i, seq in enumerate(self.sequences):
			probs.append(self.calc_ll_at_k(seq, q, self.aiList[i]))
		p = np.prod(probs)
		if p > self.best_prob:
			self.best_prob = p
			self.best_q = q

	def get_best_motif(self): # requires self.best_q to be set
		motif = []
		for col in self.best_q.T:
			i = col.argmax()
			motif.append(self.alphabet[i])
		return "".join(motif)

	def run(self):
		self.rand_ai_start()
		for _ in range(self.iterations):
			Si, seqIndexes = self.rand_si()
			q = self.calc_q(seqIndexes)
			lls = self.calc_ll_all(Si, q)
			self.new_ai_prob(Si, lls)
			self.update_best(q)
		print self.best_q
		print self.get_best_motif()

G = Gibbs(sys.argv[1])
G.run()
