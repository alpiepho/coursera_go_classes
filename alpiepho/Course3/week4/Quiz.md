1.
Question 1
What line of code could be used to define a loop which iteratively reads from a channel ch1?

1 point

for i <- ch1


X - for i := range ch1


for i, err <- range ch1


X! - for i := ch1

2.
Question 2
What does the select keyword do?

1 point

Executes a set of case statements.


X - Allows a choice of channels to wait on.


Chooses the greatest of a set of numbers.


Chooses an element from a list based on a user-defined criterion.

3.
Question 3
What is the meaning of the default clause inside a select?

1 point

X - The default clause is executed if all case clauses are blocked.


The default clause is executed before any case clause is executed.


The default clause is executed after any case clause is executed.


The default clause is executed only if a case clause is executed.

4.
Question 4
Suppose that there are two goroutines, g1 and g2, which share a variable x. X is initialized to 0. The only instruction executed by g1 is  x = 4. The only instruction executed by g2 is x = x + 1. What is a possible value for x after both goroutines are complete?

I.   0

II.   1

III.   4

IV.   5

1 point

I and II only.


II and III only.


I, II, and III but not IV.


X - II, III, IV, but not I.

5.
Question 5
What is mutual exclusion?

1 point

When a single goroutine can execute only one of two blocks of code.


When a single goroutine cannot execute a block of code.


X - When multiple goroutines cannot execute blocks of code concurrently.


When a single goroutine is the only goroutine which ever accesses a variable.

6.
Question 6
What is true about deadlock?

I.   It can always be detected by the Golang runtime

II.   Its caused by a circular dependency chain between goroutines

III.   It can be caused by waiting on channels

1 point

I and II only.


X - II and III only.


I and III only.


I, II, and III.

7.
Question 7
What is the method of the sync.mutex type which must be called at the beginning of a critical region?

1 point

X - Lock()


Unlock()


Take()


Block()
