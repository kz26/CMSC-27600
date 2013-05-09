#!/usr/bin/python

import math
import random
import sys

def parseInputFile(fn):
	"""
	Parses an input file in the format specified for this assignment.
	"""
	with open(fn) as f:
		poly = [float(x) for x in f.readline().strip().split()]
		iters = int(f.readline().strip())
		return (poly, iters)

def boltzmann(deltaE, T):
	"""
	The Boltzmann acceptance probability function.
	"""
	assert T != 0
	k_b = 1.3806488e-23
	return math.exp(-abs(deltaE) / k_b * T)

def newRandomState(x):
	"""
	The state transition function.
	"""
	return x + random.uniform(-1, 1)

def simpleCool(T0, i):
	"""
	A basic cooling schedule function.
	"""
	alpha = 0.99995
	return (alpha ** i) * T0 

class Polynomial(object):
	"""
	Polynomial class that includes an evaluation function.
	"""
	def __init__(self, coeffs):
		self.coeffs = coeffs

	def evalAt(self, x):
		val = self.coeffs[0]
		for i, c in enumerate(self.coeffs[1:], 1):
			val += c * math.pow(x, i)
		return val

class SimAnneal(object):
	"""
	A modular class for simulated annealing. Can plug in different functions
	for state transition, acceptance probability, and cooling schedule.
	"""
	def __init__(self, func, iters, x0, T0, newStateFunc, acceptanceFunc, coolingFunc):
		self.func = func
		self.iters = iters
		self.x0 = x0
		self.T0 = T0
		self.newStateFunc = newStateFunc
		self.acceptanceFunc = acceptanceFunc
		self.coolingFunc = coolingFunc

	def run(self):
		"""
		Perform the simulated annealing.
		"""
		curX = self.x0
		curY = self.func.evalAt(curX)
		curT = self.T0
		for i in range(1, self.iters + 1):
			if curT == 0:
				break
			nX = self.newStateFunc(curX)
			nY = self.func.evalAt(nX)
			if nY <= curY or self.acceptanceFunc(nY - curY, curT) > random.random(): # accept if lower energy or probability check passes
				curX = nX
				curY = nY 
			curT = self.coolingFunc(self.T0, i)	
		return (curX, curY)

if len(sys.argv) != 2:
	sys.stderr.write("Usage: %s <input file>\n" % (sys.argv[0]))
	sys.exit(1)

poly_coeffs, numIters = parseInputFile(sys.argv[1])
polyFunc = Polynomial(poly_coeffs)
sa = SimAnneal(polyFunc, numIters, 0, 1000, newRandomState, boltzmann, simpleCool)
print sa.run()[0]
