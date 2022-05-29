---
title: "Golang"
date: 2022-05-27T17:08:19-03:00
weight: 35
---


## Introduction to golang


### Why did i pick golang

I don't relly code in go anymore but i'm very found of the language and in
my opinion in one of the best language to learn programming.

It as the following good point:
* Very clean and simple syntax that is really easy to read
* Purrated list of feature to not overwhelm the user
* Great way of handling error
* Good exemple to learn about compiled language without beging hard to setup
* Awsome documentation and tooling.
* Not a fragmented ecosystem
* Go very well with vim
* Fun and enjoyable for all of the reason above.

Also for the context of web application for full stack i could have go with a high
level language and more dynamic language like Kotlin, Typescript, Python but i prefer
go to be allow to also teach some more low level programming that interfer with memory
more closely but not to mutch. And we will be able to build MVC style Rest API farelly
easily.

### What kind of language is go

Lets for a another time bring the first sentance from wikipedia. And split in down for you to understanding all of the jargon.

> Go is a statically typed, compiled programming language designed at Google by Robert Griesemer, Rob Pike, and Ken Thompson.It is syntactically similar to C, but with memory safety, garbage collection, structural typing,[5] and CSP-style concurrency.[12] It is often referred to as Golang because of its former domain name, golang.org, but its proper name is Go.

#### Statically Typed.

We didnt talk about type and variable yet. Let's do a small introduction to variable and type.

A variable is a "" that is use to store a value. Like a mathematical exemple `x = 5`. We call this an assignment.

We put the value of `5` inside a variable call `x`. After this we can use this variable for calculation.

Here is a small exemple in go.

```go
// We assign 5 to the variable x
func main() {
    x := 5
    fmt.PrintLn("x = " + x)
    fmt.PrintLn("x * x = " + (x * x))
}
```

```bash
$ go run exemple_variable.go
x = 5
x * x = 25
```

Here there is two new concept. First we use a string (a list of character between the "") and we additinal the result
of the multiplication of `x * x` to the string to get the result that is shown.

Lets go back to the variable `x` it receive the value of `5` witch is a hole number. So the variable `x` is an `Integer`.
In a language like go all variable must have a type and only value of this type can be assign to the variable

```go
// := tell that we create a new variable and go infer the type with the value pass of the right side
x := 5
// We could explicitly tell go compiler that the variable is an Integer
var y: Int = 3 // this is valide

var x: Int = "a string" // this is invalid you must pass a Integer to x , the compiler will not accept that
// You can reassign the value of the variable , since y and x are integer we can pass the value of x to y
// so now the value of y is 5 instead of 3
y = x

// Invalid , we must prove a Integer
y = "string"
```

to go back to what is statically typed , that mean first that every variable must have a type when creating the variable,
you can leave the variable empty (no value) but it must have a type. The reason for that is that the program will allocate
memory based on the size of the type. Each type have specific size in memory (RAM).

Secondly it also mean that the type of a variable cannot change in the scope (will talk later about scope and stack)
because this would possibly overflow the memory. Like a Integer take 64 bit of memory but a string could take 1..N bit
of memory since it's a list of character so if the compiler would allowed use to put the string in the integer we would
overflow the memory allocated and corrupt something else possibly.

Statically Typed is in oposition to dynamically typed witch is more commun with interpreted language like python or javascript
is those language the compiler does care about what you put in your variable you can reassign it with different type and etc..

The advantage of statically type is that the compiler and your code editor always knwo the type so it can optimize and gave
you special recommandation about the option of what you can or cannot do with it. It leave less space for error.

Will see later on in web development but language have been build on top of javascript to add typing to facilitate the
work by big team and to catch possible error at compile time instead that a runtime.

Sorry this part is heavy yet in information , don't worry ask question and you will understand at some point.

### Memory safety

I cover a little bit of this in the previous section. But a memory safety language is made to protect you agains
error like buffer overlow (when you try to access memory bewound what is allocated and some other things)

### Garbage collection

Most language have this feature in one way or another. Except language like C and C++.
Garbage collection is code that run with your code to clean the memory of things that are no longer use by your program.
Like in the exemple prior when i'm done using x , the garbage collection will remove  the variable from memory to free up
space. This is half true , when we will learn about Scope, Stack and Heap this will make more sence.

In language like C when you allocate a block a memory you are responsible to clear it when done using it, if not
your program will start to use more and more of memory and nobody want this , except maybe hacker.

The conn of the garbage collector is that it run after you code from time to time to clear the memory, this introduce
some latency to your application and in some really high reactive application this can be an issue. But take my world
it may never be a issue for most of us.

Some modern language like Rust and Modern C++ have garbage collection without garbage collection with a principale of Ownership
or Smart Pointer. To cut it short they keep reference of how is using it and when it drop to zero they delete it from memory.

### Structural typing

We will come back to this latter on in the pradigm section of the course.


### Concurrent CSP style

A concurrent language mean that the language can run multiple task at the same time. In different core of the computer
or inside the same core. This break the linearity of the code. We will not go to deep in this for now but go have a awsome
and simple to use system to run code concurrently.


## Thing to read

* https://en.wikipedia.org/wiki/Go_(programming_language)
* https://en.wikipedia.org/wiki/Memory_safety
* https://en.wikipedia.org/wiki/Garbage_collection_(computer_science)
* https://en.wikipedia.org/wiki/Structural_type_system
* https://en.wikipedia.org/wiki/Concurrency_(computer_science)



