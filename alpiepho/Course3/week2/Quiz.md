1.
Question 1
What unique resources does a process have?

I. Program counter value.

II. Virtual address space.

III. Shared libraries

1 point

I only.


I and II, NOT III


II and III, NOT I


II only

2.
Question 2
What does an operating system do to enable concurrency on a single processor machine?

1 point

Provides a graphic interface to allow the user to control system functions.


Partitions processor hardware to allow parallel execution of multiple processes.


Combines different processes into a single process.


Interleaves the execution of different processes .

3.
Question 3
What is the “context” that is referred to in the term “context switch”?

1 point

Shared libraries used by a process.


Memory and register values unique to a process.


The parameters specific to the operating system.


The set of executing processes.

4.
Question 4
What is the difference between a thread and a process?

1 point

A thread has less unique context than a process.


Threads do not have unique program counter values.


Processes can execute programs but threads cannot.


A thread is another name for a process.

5.
Question 5
What is the main function of the Go runtime scheduler?

1 point

Schedules operating systems processes.


Schedules goroutines inside an operating system thread .


Schedules operating system threads inside a process.


Assigns operating system threads to hardware processors.

6.
Question 6
Suppose that there are two goroutines executing, g1 and g2. Assume that g1 requires 1 second to complete when it is executed alone, and g2 requires 2 seconds to complete when it is executed alone. Assume that there is no synchronization between g1 and g2. Assume that g1 and g2 are executed concurrently and that g1 begins executing first. What do we know about the relative completion times of g1 and g2?

1 point

g2 will complete before g1.


Nothing!


g1 will complete before g2.


g2 and g1 will complete at virtually the same time.

7.
Question 7
Assume that two goroutines are executed concurrently. Which of the following statements is true?

1 point

The relative order of the execution of their instructions is deterministic.


The relative order of the execution of their instructions is unknown, but it is the same each time they are executed together.


The relative order of the execution of their instructions can be different every time that they are executed together.


The relative order of the execution of their instructions can be determined from startup conditions.
