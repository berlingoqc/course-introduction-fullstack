---
title: "Main programming concept"
date: 2022-05-27T17:09:46-03:00
weight: 40
---

In the chapiter get your head ready we gonna start really learning the basic of programmation.

## Function

If you remember in the hello world exemple that we did prior we had our `main` function that 
was the entry point of our application , this function is different from other because it's call
once and not by us manually.

But normally a function is a piece of reusable code that you can call from other function in your
code. Because all statement reside inside of function.

### What is a statement

We said them earlier but a statement is when you do a logical action (NEED REAL DEFINITION), like assign
a variable , do a calculation and etc.. All that make the code moving. The pretty mutch all reside inside
function except some static variable definition.

### Back to function

A function is define by it' signature.

A signature is compose of:
* The availability (private vs public vs other things depending of the language)
* The name of the function
* It's argument
* It's return type

In some language and other case there is other information in the signature but dont worry for now.

Here is an exemple of a function.

```golang
// The signature , this function take one argument x of type Int and return another Int.
// Function can have multiple argument of different type and the always return One value
// But we will learn latter that One value can be multiple value
func power(x: Int): Int {
    return x * x
}

func main() {

  x := 5
  power_of_x := power(x)

  // This will print 25
  fmt.PrintLn(power_of_x)

  // We dont have to assign the result of the function call the a variable we can pass it
  // directly to a another function.
  // this return the same result.
  fmt.PrintLn(power(5))
}
```

There is lot more to learn and that we can do with function but this is the basic for now.

## Scope

The scope is like your current context, it defined what is available to you.

We will start with the previous exemple.

```go

// The first scope is the file your in

// Here we defined a constant variable in the base of the file
// this will be available to every one in the file.
const y := 5

// A scope is defined between {} so a function has it's own scope
func mystery_calculation(x: Int): Int {
   // In my scope i have the variale x defined with my arguments
   // And i call access y the global scope 
   return x * y
}

func main() {
   // In the scope of my main function i declare g
   g := 5
   // I create a second scope
   {
     // In the new scope a create the variable h
     h := g * y
   }
   // Here i quit the scope , so the variable h is no longer there because it was living in the
   // inside scope, we will lean more about this with the Stack in the next section.
   // this will throw an error because h is not defined
   k := h * g 
}

```

## Stack and Heap

Stack and the Heap are two fondamental concept on how your software work and here are data
store in the memory and how your programm is being run. I guess you can be a web developer
without really know mutch about this but it does not hurt to know about this.

The stack and the Heap are the two place where youre program store it's variable.

The stack is like a pile of plate , when you put stuff on the stack it's pile up and when you
remove stuff it's always the last element added.

The stack and the scope are link together.
When you enter a scope and you create variable it's added to the stack end when the scope end.
All variable created in the scope are remove from the stack. So when you enter a function , enter
a new scope block it's get added to the stack and when you leave this scope , all things created are
deleted.

This is great but what do we do when we need to create variable of dynamic size that persist there scope.
In those case we must use the heap. The heap is what's getting garbage collected.

(WIP more info on the heap i'm a noob without internet)

## Different type

We only talk for now about two type : `Int` and `String`. But there is many more type and most
of the programming job you will have to do is create new type to represent some more complex value.

First we have the primitive type. Those type all have a fixed size in memory and are store in the stack (????)

* Int
* Float
* Boolean
* Char
* ...

After we have the "???" type.

* String : a list of character, text
* Struct : A structure in a ensemble of multiple property that all have a type and name.
* Tuple : It's a collection of type without name
* Array : It's a list of value of a particuliar type
* Map : Map is a collection of Key and Value. All key have the same type and all value have the same type.
* Typealias : When we create a alias of a particular type for syntastic reason.


```go
struct Time {
    hour: Int,
    minute: Int,
    second: Int
}

func main() {
    a_int := 5
    a_float := 0.5
    a_boolean := true
    a_char := 'd'
    a_string := "I'm a string motherlover"

    a_struct_time := Time {
       hour: 5,
       minute: 40,
       second: 34
    }

    # You can access your property
    g := a_struct_time.hour

    a_array_float = [0.4, 0.5, 0.7, 0.7]

    # You can acces an array by index
    x := a_array_float[1] # the first index is zero so the variable x gonna have the value of 0.5 and a float type

    a_map = {
        "claude": 50,
        "mario": 40
    }

    y := a_map["mario"] # you access with the key , so the value of y gonna be 40

}
```
From this point in the tutorial . I'll do a guide video on doing the go tour that you can go check
on your own. It will teach more in detail what  i started to teach in this and give your on good exemple
of all the condition and loop things.


## Standard Library

* https://pkg.go.dev/std

> A standard library in computer programming is the library made available across implementations of a programming language.

This is normally part of the language and maintan by the same team that build the language.

All of the function from the standard library can be use anywhere and most other library that we will cover are based
on those functionallity.

Small comeback to the compilation. The standard library is ship inside your exuctable with your code. It is statically link
like all library that are written in go.

We will explore some on them in the live exemple and in the homework.

We aready use a package of the std (standard library) the `fmt` package to write to the terminal.

## Live Exercice

* https://go.dev/tour/welcome/1
* https://gobyexample.com/ (Complete to the Method section, read the code and try it localy if you want)

## Homework

Now with all that we learn today we can try to make a small homework application to put all of this together.
I will gave some exemple of starting point to create logic with interaction with the user to generate some kind
on modify ouput data.

Like this exemple:

Write a program with a variable called age assigned to an integer that prints different strings depending on what integer age is.

The program will have the following flow:

* Ask the user the age of the person requiring advive. (May ask for more data to foreast it's prediciton)
* Print some kind of predicition based on the given data.
* Most only use functionality from the standard library.

Optional extra

* Give a prediciton based on the day. (The result is always the same for a name and a date)

For this appli

### Code snippet to get your started

#### Ask for a answer from the 



## Things to read

* https://gobyexample.com/
* 


