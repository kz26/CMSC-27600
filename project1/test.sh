#!/bin/bash
# Test suite

./main data/test1.in | diff - data/test1.out
./main data/test2.in | diff - data/test2.out
./main data/test3.in | diff - data/test3.out
./main data/test4.in | diff - data/test4.out
./main data/test5.in | diff - data/test5.out
./main data/test6.in | diff - data/test6.out
