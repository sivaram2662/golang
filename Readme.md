what is Golang 
       Golang is Free open source programming language created by Google.Also referred to as go it is designed to be simple, scable,ad effective.since its inceptions 2009, Golang has grown in popularity in the software development communiny,a popular general-purpose computing programming language.

 ---->Go was created to combine.
          1) The ease of programming of an interpreted, dynamically typed language (such as python)
          2) With the Efficiency and safety of a statically typed, compiled language.(such as c++)
          3) It also amied to be modern ,with support for networked and multicore computing.


Lets at understanding Basic programing in Golang:
=================================================
Package declaration:
                  In Go, every program is part of a package. We define this using the package keyword. In this example, the program belongs to the main package

Import packages: 
                import ("fmt") lets us import files included in the fmt package.A blank line. Go ignores white space. Having white spaces in code makes it more readable.

Functions: 
           func main() {} is a function. Any code inside its curly brackets {} will be executed.

  fmt.Println() is a function made available from the fmt package. It is used to output/print text. In our example it will output "Hello World!".
  
Basic code in Golang
====================
package main

import "fmt"

func main() {
	fmt.Println("Hello World")
}


Go comments:
============
comments are two types :
1) single-line-comments:
     Single-line comments start with two forward slashes (//).
     Any text between // and the end of the line is ignored by the compiler (will not be executed).

  Example:
  =======
     package main

     import "fmt"
  
     func main(){
      fmt.println("This for single-line-comments) // This is comments
     }


2) Multi-line Comments:
   Multi-line comments start with /* and ends with */.
   Any text between /* and */ will be ignored by the compiler:
 
 Example:
 ========
  /*
   package main

     import "fmt"
  
     func main(){
      fmt.println("This for single-line-comments) 
     } */
