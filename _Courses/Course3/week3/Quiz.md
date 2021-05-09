1.
Question 1
Suppose you want to start a goroutine which executes a function called test1(). What code would create this goroutine?

1 point

test1()  go


start test1()


goroutine test1()


go test1()

2.
Question 2
When does a goroutine complete?

I. When its code completes.

II. When all goroutines complete.

III. When the main goroutine completes.

1 point

I and II, NOT III.


I and III, NOT II.


I, II, and III.


I only.

3.
Question 3
Synchronization is useful for what purpose?

I. Restrict illegal interleavings.

II. Force events in different goroutines to occur in sequence.

III. Allow a goroutine to continue to execute after the main goroutine has completed.

1 point

I, II, and III.


I only.


I and III, NOT II.


I and II, NOT III.

4.
Question 4
If a goroutine g1 is using a WaitGroup wg to wait until another goroutine g2 completes a task, what method of the the WaitGroup should be called when g2 has finished the task?

1 point

wg.Done()


wg.End()


wg.Finished()


wg.Alarm()

5.
Question 5
If a goroutine g1 is using a WaitGroup wg to wait until another goroutine g2 completes a task, what method of the the WaitGroup should be called before g2 starts its task?

1 point

wg.Fork()


wg.Start()


wg.Add()


wg.Begin()

6.
Question 6
How might you write code to allow a goroutine to receive data from a channel c?

1 point

x <- c


x = <- c


x = c


x <-- c

7.
Question 7
What is the difference between a buffered channel and an unbuffered channel?

1 point

A buffered channel can hold multiple objects until they are read. An unbuffered channel cannot.


A buffered channel delays the transmission of data. An unbuffered channel does not.


A buffered channel delays the reception of data. An unbuffered channel does not.


A buffered channel can communicate between more than 2 goroutines. An unbuffered channel cannot.