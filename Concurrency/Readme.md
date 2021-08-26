### Why concurrency is hard

It is very common for bugs to exist in concurrent code. These bugs often gets discovered in course of time when heavier disk utilization occurs or more users logs into the system. Hopefully these bugs are common and and computer scientists have been able to label them, and found out how to solve them as well. Some of these common issues are Race conditions , Atomicity, Memory access synchronization, deadlock , livelock, Starvation. 



### Race Conditions

A race condition occurs when two or more operations must execute in the correct order, but the program has not been written so that this order is guaranteed to be maintained. 








### Concurrency VS Parallelism

Concurrency is a property of the code; parallelism is a property of the
running program.
