# CMSC 27600, Project 1
* Kevin Zhang
* April 25, 2013


## Compile:

You will need the latest stable version of go1: http://code.google.com/p/go/downloads/list

Set the GOPATH environment variable to the main project directory, i.e. the directory 
containing the 'src' subdirectory, main.go, this readme file.
If you are in this directory, simply run:

    export GOPATH=`pwd`

Now run:

    go build main.go

This will produce an executable "main" file.


## Run:
The program takes exactly one argument, the path to an input file.

    ./main data/test1.txt

To run the test suite, which diffs against all six posted test files, do:

    ./test.sh
