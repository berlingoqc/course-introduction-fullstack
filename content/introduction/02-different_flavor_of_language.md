---
title: "Different language type"
date: 2022-05-27T17:08:48-03:00
weight: 20
---


# Interpretted vs Compiled language


In the previous chapiter i talk about how we need to convert our source code to machine code.
But it's not always the case , something a middle man is doing it for us.

The exemple that i give is in go witch is a compile language , it need to be transform into machine
code and link to the required library to work. But watch this magic.

```bash
➜  golang_hello_world git:(master) ✗ go run hello_world.go
Hello, world!
```

Boom i don't need to compile the source code to execute it does that mean that it's not compile.
No it's just trickery. Behind the scene the same step are executed that in the previous chapiter
it is only for pratical reason to save some step when running a small program.

But what is an interpretted language, let's go straight out of wikipedia


In computer science, an interpreter is a computer program that directly executes instructions written in a programming or scripting language, without requiring them previously to have been compiled into a machine language program. An interpreter generally uses one of the following strategies for program execution:

* Parse the source code and perform its behavior directly;
* Translate source code into some efficient intermediate representation or object code and immediately execute that;
* Explicitly execute stored precompiled bytecode[1] made by a compiler and matched with the interpreter Virtual Machine.

I will not write more here about this , the wikipedia page does a really good explanation of everything that there is to know.

The big conn of interpretted language is that you need to have the interpreter install on the machine to execute the code, it's fine
in a lot of case but can be quite annoying for non technical user who need to install another application to run there application.

They other conn is performance but on the modern world it's rarely an issue to have overhead with the power on modern computer.

And for the pros it's the fast iteration speed. By iteration it means that we can make modification faster and run it more easily
( but i can argue that some interpretted language are more difficul to run that compile language, looking at you python )

I think to most reliable exemple of this is Java (hello to all Minecraft player). Java fall in the last of the case enumerate earlier. It's a language that is
compile into bytecode that is latter interpreter with the Java Virtual Machine (JVM) , what people are really talking about when they said: `do you have java ?`

Speaking trully it's really pratical like you only need to download the `jar` file like for minecraft and you can execute it on linux , macos , windows name it.
But for noobies it's a annoying step to have to intall java, it's abstract for them , i just want to play minecraft why can i just install it directly.

And for me , i'm not a fan of interpreter language when i need to deploy it to user but in control environment like for server code it's fine.
That where language like go shine, in most case they don't need anything at all other than the application to run making it really easy to
deploy to thousand of computer. (you will see that a lot of tooling are writting in code for this exemple reason but not mutch in python) 


## Lecture from this chapiter 

* https://en.wikipedia.org/wiki/Interpreter_(computing)
