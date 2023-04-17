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






=======================================================================================================================================================================================================================================

----Handler function the process an http request and http response 

----The Function takes two parameters : Http.responses write ,object which is used to write the response data,and a http.request object which represent the incomming Http request .The Fmt.Fprintln() function writes the response to the writer object.


----This handler function can be registered with the default HTTP server using the http.Handlefunc() function

----%S is used specially to format a string value .prints the string as it without any modification.

---%V is a general purpose formating verb that can represent any value human-readable format. %V is prints the field names and their corresponding values.

LOG ----go language the log package is used for logging messages during the execution of a program. Logging is the process of recording the events and actions that occur during the execution of a program.

fatal----In the Go standard library, the "fatal" package provides functions for reporting fatal errors, such as "Fatalf" and "Panicf". These functions are typically used to print an error message and terminate the program with a non-zero exit code.

Defre------In Golang, defer is a keyword used to schedule a function call to be executed when the surrounding function completes The defer statement takes an expression or a function call as an argument, and the expression or function call is evaluated immediately, but the execution is deferred until the surrounding function returns.



By using  golang ec2 instance stop and start link :https://dev.to/aws-builders/starting-and-stopping-ec2-instances-using-aws-sdk-for-go-5ig



====================================================================================================================

API SECURITY
============
what is mena by api security?

A tool that allows applications to communicate with other application. it is used to also authentication and authorization.api works on backend framework for and web application it critical to procete the sensitive data they trasfer.

JWT---JSON WEB TOKEN
=====================
what is mean by jwt ?

-----jwt stands for json web token it is used to authentication and authorization based on token .
  
The golang-jwt package is the most popular package for implementing JWTs in Go, owing to its features and ease of use. The golang-jwt package provides functionality for generating and validating JWTs.

you have import the golang-jwt package.

--------go get github.com/golang-jwt/jwt to excute the command in server.

create login api  user login authentication and authorization and every request passwd to the client 
homeapi and auhentication and authorization token then token is valide and user is allowed to the homepage
request api 


json web token structure is three parts 
         1) Header : this header consists of two parts type of token and which algotitham 
                 for example : {
                   "alg":"HS256"
                   "typ":"JWT"
                 }

       2) payload : The second part of the token is payload .claims are statemts about an entity and additional data.there are three types of claims 

                     1) registered claims ---these are the predefined claims,it is set of usefull claims and sub,aud,exp ,iss
                     2)public claims ---defined as a URI that contains a collision resistant namespace both are the claims and methods used it 
                     3)private claims ---These are the custom claims created to share information between parties that agree on using them and are neither registered or public claims.

                     {
                        "sub": "1234567890",
                        "name": "John Doe",
                        "admin": true
                                      }
         3) signature :
    To create the signature part you have to take the encoded header, the encoded payload, a secret, the algorithm specified in the header

       this signature also one algorithm is there for example 
       HMACSHA256(
  base64UrlEncode(header) + "." +
  base64UrlEncode(payload),
  secret)







   





