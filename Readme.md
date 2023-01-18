Golang language
====================
1) golang is open-source platform.It was developed in 2007 Rebert griesemer and ken thompson.
2) Go is fast , statically typed, compiled language that feels like a dynamically typed, interpreted language.
3) Go syntax is simillare to c++ and python 

Go install process:
==================
You can find the relevant installation files at : https://golang.org/dl/.
To check if Go was installed successfully, you can run the following command in a terminal window: go version


go syntax
===========
Package declaration: In Go, every program is part of a package. We define this using the package keyword. In this example, the program belongs to the main package

Import packages: import ("fmt") lets us import files included in the fmt package.A blank line. Go ignores white space. Having white spaces in code makes it more readable.

Functions: func main() {} is a function. Any code inside its curly brackets {} will be executed.

fmt.Println() is a function made available from the fmt package. It is used to output/print text. In our example it will output "Hello World!".

comments
==========
comments are two types :
1) single-line-comments:
     Single-line comments start with two forward slashes (//).
     Any text between // and the end of the line is ignored by the compiler (will not be executed).
2) Multi-line Comments:
   Multi-line comments start with /* and ends with */.
   Any text between /* and */ will be ignored by the compiler:

Basic.go
=========
==our 1 st program wiil print the  hello world message.

To run the program, put the code in basic.go and use go run.

values.go
===========
Go has various value types including strings, integers, floats, booleans, etc. Here are a few basic examples in 

varibles.go
===========
in go different types of variables.
1) Int 
2) float 
3) string
4) bool 
 
declaring and creating variables are two types:
===============================================
1) var keywords 

constant.go
============
Go supports constants of character, string, boolean, and numeric values.
constant declares a constant value.

