#Makefile is a set of instructions for the make program,
#which helps to assemble a software project literally with one touch.
.PHONY: build 
#A phony target is one that is not really the name of a file; 
#rather it is just a name for a recipe to be executed when you make 
#an explicit request. There are two reasons to use a phony target: 
#to avoid a conflict with a file of the same name, and to improve performance. 

build:
	go build -v ./cmd/apiserver

.PHONY: test
test:
	go test -v -race -timeout 30s ./...

.DEFAULT_GOAL := build
#goal automatically works when the make command is executed

