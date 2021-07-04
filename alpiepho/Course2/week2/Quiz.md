1.
Question 1
Is the highlighted assignment to f
in the following code a legal variable assignment?

var f func(string) int

func test(x string) int {

return len(x)

}

func main() {

f = test

}

1 point

X - Yes


No

2.
Question 2
Which of the following statements correctly
declares a function whose argument is another function which takes an integer
as an argument and returns a string?

1 point

func fA(fB (int) string)


X - func fA(fB func (int) string) 


func fA(fB func (int))
string


func fA(fB func (string) int) 

3.
Question 3
What is an anonymous function?

1 point

A function with no return value


A function with multiple names


X - A function with no name


A function with no arguments

4.
Question 4
Which of the following statements correctly
declares a function whose return value is another function which takes a string
as an argument and returns an integer?

1 point

func fA(fB (int) string) func (string) int


X - func fA(fB func (int) string) {}


func fA(int) string {} 


func fA() fB func (string) int{}

5.
Question 5
func fA() func() int {
    i := 0
    return func() int {
        i++
        return i
    }
}
func main() {
   fB := fA()
   fmt.Print(fB())
   fmt.Print(fB())
}

What does the above code print on the screen?

1 point

X - 12


11


01


1

6.
Question 6
What symbols are used in a function
declaration to indicate that it is a variadic function?

1 point

“->”


X - ”...”


"---"


"[]"

7.
Question 7
What does this routine produce?

package main

import "fmt"

func main() {

i := 1

fmt.Print(i)


1 point

132


134


234


X - 123
