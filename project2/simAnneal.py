#!/usr/bin/python

from decimal import Decimal
import math
import random
import sys

def parseInputFile(fn):
	with open(fn) as f:
		poly = [float(x) for x in f.readline().strip().split()]
		iters = int(f.readline().strip())
		return (poly, iters)

def boltzmann(deltaE, T):
	assert T != 0
	k_b = 1.3806488e-23
	return math.exp(-abs(deltaE) / k_b * T)

def newRandomState(x):
	return x + random.uniform(-1, 1)

def simpleCool(current_T, totalIters, i):
	alpha = 0.99995
	return alpha * current_T

class Polynomial(object):
	def __init__(self, coeffs):
		self.coeffs = coeffs

	def evalAt(self, x):
		val = self.coeffs[0]
		for i, c in enumerate(self.coeffs[1:], 1):
			val += c * math.pow(x, i)
		return val

class SimAnneal(object):
	def __init__(self, func, iters, x0, T0, newStateFunc, acceptanceFunc, coolingFunc):
		self.func = func
		self.iters = iters
		self.x0 = x0
		self.T0 = T0
		self.newStateFunc = newStateFunc
		self.acceptanceFunc = acceptanceFunc
		self.coolingFunc = coolingFunc

	def run(self):
		curX = self.x0
		curY = self.func.evalAt(curX)
		curT = self.T0
		for i in range(self.iters):
			curT = self.coolingFunc(curT, self.iters, i)
			if curT == 0:
				break
			nX = self.newStateFunc(curX)
			nY = self.func.evalAt(nX)
			if nY <= curY or self.acceptanceFunc(nY - curY, curT) > random.random():
				curX = nX
				curY = nY 
			curT = self.coolingFunc(curT, self.iters, i)	
		return (curX, curY)

poly_coeffs, numIters = parseInputFile(sys.argv[1])
polyFunc = Polynomial(poly_coeffs)
sa = SimAnneal(polyFunc, numIters, 0, 3000, newRandomState, boltzmann, simpleCool)
print sa.run()[0]
