---
title: "What is code"
date: 2022-05-27T17:07:50-03:00
weight: 20
---


Topics of the page:

* Computer Architecture (basic)
* OS Architecture (basic)
* Source code and Machine Code
* Compilation

## The computer

Explanation of the computer architecture and the different layer that build the modern operating system that we use.

### Computer Architecture

Modern computer have the following component that our program work with

* CPU
* RAM
* GPU
* Storage

### OS Architecture


!(abstraction-layer)[https://developer.ibm.com/developer/default/articles/l-linux-kernel/images/figure2.jpg]

#### Hardware Abstraction Layer

On top of those component we have the HAL (Hardware Abstraction Layer)

#### Kernel

And an layer is build on top of it called a kernel that implement all the logic
of interaction with this ever changing hardware.

#### User space

The user space in where our program are run. This space have less right and use the kernel
and system service that expose the core functionnality of the os.

## Back to the code

### Analogy with cooking receipe

Programming is like the good old probleme solving that i did in elementary school.

You start with a basic problem or need and you write steps to accomplish it , like a cooking receipe.

When you write a cooking receipe you start with the element that you need , you defined what you need
and sometime it can vary depending of the availability , personal requirement and soo one.

After that you have the instruction on how accomplish the end goal the dish.

Sometime in receipe you don't want to repeat yourself so you link another receipe or cooking technique.

All those concept translate directly to concept of programming, it's pretty mutch the same but way way
more complexe.

For receipe we write them in our speaking language and the user of the receipt read those step and
interpret them. In this case we are the computer and we understand the language directly.

But when it come to computer they don't understand at all our language ou even the language that
we wrote the intruction in. We need to transform our instruction to instruction that the computer
will understand.

### From source code to machine code

What the understand is [Machine Code](https://en.wikipedia.org/wiki/Machine_code) and [Instruction Set](https://en.wikipedia.org/wiki/Instruction_set_architecture)

You may be familliar with [x86](https://en.wikipedia.org/wiki/X86) , it's the instruction set that is use by most
home computer since the 80's (but it's changing now with ARM like in the new MacBook M1).

To go from our programming language to this machine code we must use a compiler (not entirely true will see more in the next chapiter)
[this page](https://www.webopedia.com/definitions/compilation/) give a fairly straight forward explanation.


### Exemple of this process

Here is an exemple of this process with the infamous hello world writting in go the language that we will use in the course.

This is the content of our `hello_world.go` source code file.

```go
package main

// I'm a commentary , i'm not gonna get turn into machine code
// I'm there so the programmer can shared some additional information

// Don't bother yet with this put in order to print information to the terminal
// we need to import this functionality
import "fmt" 
 
// All code instruction reside inside function , the main function in the starting
// point of our application , this is what gonna get execute by the machine.
// We will learn more about function in the fourth chapiter
func main() {
    // Call the Println function from the fmt package to write this string
    // to the terminal.
    fmt.Println("Hello, world!")
}
```

To convert this source code to machine code we need to use a compiler. Each language has it's own compiler (sometime multiple language shared the same)

The go compiler is call `go` it's an executable that we are gonna setup in the third chapiter when you gonna start to get your hand dirty.


{{% notice note %}}
The next snippet is not code per said , but it could be in a script (chapiter 11). Each line is a command that you type in a terminal
{{% /notice%}}

```bash
# This command output is a executable file with the name hello_world
# that is compile to run on the os and the architecture of our computer
➜  golang_hello_world git:(master) ✗ go build -o hello_world ./hello_world.go
# When we execute the file we get the output
➜  golang_hello_world git:(master) ✗ ./hello_world
Hello, world!
```

This is the really easy step to compile go source code , not all language are that easy. In language like C and C++ we have to manually
do the multiple step like converting to object and linking manually, nobody got time for that. If you want you can go learn more 
about compilation , machine code and assembly but it's totally optional to learn this to become a full stack developer or most kind
of developer.

This is our first coding exemple. We have manage to turn source code to executable how wonderful.

In the next chapiter we will see how this is not always the case for the user but it's always the case for the machine.

## Cross-plateform code

## Lecture from this chapiter

* https://en.wikipedia.org/wiki/Central_processing_unit

* https://www.webopedia.com/definitions/compilation/ (required)
* https://en.wikipedia.org/wiki/Machine_code (optional)
* https://en.wikipedia.org/wiki/Instruction_set_architecture (optional)
* https://en.wikipedia.org/wiki/X86 (optional)
* https://getstream.io/blog/how-a-go-program-compiles-down-to-machine-code/ (very optional , advanced)

