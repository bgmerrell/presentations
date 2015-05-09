************************
Introduction to Go Notes
************************

===========
Hello World
===========
“func"

UTF-8

=====================
Variable Declarations
=====================
Multiple assignments

:= vs =

Println takes any type, and we will talk about how later

=========
Functions
=========
Calling

Return types

================
Type conversions
================
Run and show that the conversion fails

Do the conversion so that it passes

Show that there’s no implicit conversion even when you don’t lose
precision (i.e., change printInt to printFloat and try to pass in an int)

=======
Looping
=======
First time we’ve seen +=
First time we’ve seen const

============
More looping
============
Removing outside of semicolons leaves just the conditional (loop1())

Go allows you to just remove the semicolons in this case (loop2())

=======
If/else
=======
Create compound if statement with ``x := getX()``

Pitfall: scoping

========================
Switch with no condition
========================
Alternative if/else

=================================
Structs (defining your own types)
=================================
Various forms of (C-like) printing 

Pointer vs. non-pointers

======================
Don’t fear Go pointers
======================
Show pointer syntax: ``*, &``

Can return reference of local variables

Talk about dereferencing syntax

Run and compare values and types

======
Slices
======
make initializes the slice, map, and channel types (while new allocates memory)

slicing expression

array backing, len vs cap

``[…]int{1,2,3}`` or ``[3]int{1,2,3}`` would be an array

=====
Range
=====
Can take off ``v`` to ignore value

Can use ``_`` to ignore first value

Could change ``[]int{1,2,4,8,16}`` to ``[…]int{1,2,4,8,16}`` to make it an array

====
Maps
====
"Comma OK" idiom

=====================
Functions vs. Methods
=====================
People sometimes use the two interchangeably, but they are two distinct things

Go has a neat and novel way of creating methods

=======
Methods
=======
“Receiver”

Notice how the function takes no arguments and the data of the object is being acted on

=================
Methods continued
=================
Ask what will be printed before and after ``DoubleValues()``

Run, notice that the printed values are the same, why?

Change receiver to a pointer receiver ``(&p).DoubleValues()`` could also be called

When do I know when to use a pointer receiver vs a value receiver?  That’s on the next slide.

==========
Interfaces
==========
Overview of Go interfaces and implicit implementation
Run program and see it fail, read the reason
Implement Fox.GetNoise()
Talk about power of interfaces, composition vs. inheritance, etc.

=======================================
Empty interfaces and variadic functions
=======================================
Everything implements ``interface{}``

This is how ``Println()`` worked on all those varying arguments in the first slides

Basically allows for dynamic typing, but beware, you lose your compile-time checks

======
Errors
======
The foo, err idiom
Run and show error case
Add some numbers to the slices
Run and show non-error case

================
Errors continued
================
A lot going on here

struct

method  that implements the error interface

so when ``&MyError{…}`` is returned from ``run()`` it can be returned as an error type

``if _, ok := err.(*MyError); ok { fmt.Println(“yep” }``

===========
Concurrency
===========
Brief description of channels

Describe functions

Run and show that the program just exist, why?

Add ``time.Sleep(2 * time.Second)`` to the end and rerun

=====================
Concurrency continued
=====================
``time.Tick`` (keep channel open and communicate every interval) and ``time.After`` (one-time thing)
Select statement wrapped in a for loop
