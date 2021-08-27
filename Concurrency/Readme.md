### Why concurrency is hard

It is very common for bugs to exist in concurrent code. These bugs often gets discovered in course of time when heavier disk utilization occurs or more users logs into the system. Hopefully these bugs are common and and computer scientists have been able to label them, and found out how to solve them as well. Some of these common issues are Race conditions , Atomicity, Memory access synchronization, deadlock , livelock, Starvation. 



### Race Conditions

A race condition occurs when two or more operations must execute in the correct order, but the program has not been written so that this order is guaranteed to be maintained. There is an instance called data race when one concurrent operation attempts to read a variable while other concurrent operation is trying to write the same variable. 

Race conditions may occur in a code where a concurrent function is being used. There's no guarranty that the operations in a code will execute sequentially. So, there could be multiple scenarios which are possible. 
```go
func main() {

	go func() {
		data++
	}()
	time.Sleep(1*time.Millisecond)
	if data == 0 {
		fmt.Printf("the value is %v.\n", data)
	}
}

```

we can use sleep to so that we can give concurrent functions to execute. Introducing sleeps into your code can be a handy way to debug concurrent programs, but they are not a solution.


### Atomicity

Regarding concurrency, atomicity rather means that when a thread modifies the state of some object (or set of objects), another thread can't see any intermediary state. **Either it sees the state as it was before the operation, or it sees the state as it is after the operation.**

For example, changing the value of a long variable is not an atomic operation. It involves setting the value of the 32 first bits, and then setting the state of the 32 last bits. If the access to the long variable is not properly synchronized, a thread might see the intermediary state: the 32 first bits have been changed, but the 32 last bits haven't been changed yet.

Something may be
atomic in one context, but not another. Operations that are atomic within the
context of your process may not be atomic in the context of the operating
system; operations that are atomic within the context of the operating system
may not be atomic within the context of your machine; and operations that
are atomic within the context of your machine may not be atomic within the
context of your application. In other words, the atomicity of an operation can
change depending on the currently defined scope. Making the operation atomic is dependent on which context you’d like it to be atomic within. If your context is a program with no
concurrent processes, then this code is atomic within that context. If your
context is a goroutine that doesn’t expose i to other goroutines, then this code
is atomic. Most statements are not atomic, let alone functions, methods, and programs.

**Atomicity is important because if something is atomic,
implicitly it is safe within concurrent contexts.**

### Memory Access Synchronization

**The critical section** is a code segment where the shared variables can be accessed exclusively. One way to solve this is Memory Access Synchronization. For this, There is a convention for developers to follow. Anytime developers want to access the variable’s memory, they must first call Lock , and when they’re finished they must call Unlock . Code between those two statements can then assume it has exclusive access to data.
```go
func main() {
	var memoryAccess sync.Mutex
	var value int
	go func() {
		memoryAccess.Lock()
		value++
		memoryAccess.Unlock()
	}()
	memoryAccess.Lock()
	if value == 0 {
		fmt.Printf("the value is %v.\n", value)
	} else {
		fmt.Printf("the value is %v.\n", value)
	}
	memoryAccess.Unlock()
}

```
we have solved our data race, we haven’t actually solved our race condition.The order of operations in this program is
still nondeterministic; we’ve just narrowed the scope of the nondeterminism a
bit. In this example, either the goroutine will execute first, or both our if and
else blocks will. We still don’t know which will occur first in any given
execution of this program. Synchronizing access to the memory in this manner also has performance
ramifactions. the calls to Lock you see can
make our program slow. Every time we perform one of these operations, our
program pauses for a period of time. 


**The previous sections have all been about discussing program correctness in
that if these issues are managed correctly, your program will never give an
incorrect answer. Unfortunately, even if you successfully handle these classes
of issues, there is another class of issues to contend with: deadlocks,
livelocks, and starvation. If not handled properly, your program
could enter a state in which it will stop functioning altogether.**

### Deadlock

A deadlocked program is one in which all concurrent processes are waiting
on one another. In this state, the program will never recover without outside
intervention.

```go
type value struct {
	mu    sync.Mutex
	value int
}

func main() {
	var wg sync.WaitGroup
	printSum := func(v1, v2 *value) {
		defer wg.Done()
		v1.mu.Lock()
		defer v1.mu.Unlock()
		time.Sleep(2 * time.Second)
		v2.mu.Lock()
		defer v2.mu.Unlock()
		fmt.Printf("sum=%v\n", v1.value+v2.value)
	}
	var a, b value

	wg.Add(2)
	go printSum(&a, &b) // our first call to printSum locks a and then attempts to lock b
	go printSum(&b, &a) // second call to printSum has locked b and has attempted to lock a
	//Both goroutines wait infinitely on each other.
	wg.Wait()
}


```
sleep for a period of time to simulate work (and trigger a deadlock).

### Concurrency VS Parallelism

Concurrency is a property of the code; parallelism is a property of the
running program.


Mutual exclusion (mutex) implies that only one process can be inside the critical section at any time. If any other processes require the critical section, they must wait until it is free.

### Goroutines

A **thread** is a basic unit of CPU utilization, consisting of a program counter, a stack, and a set of registers, ( and a thread ID. ) Traditional ( heavyweight ) processes have a single thread of control - There is one program counter, and one sequence of instructions that can be carried out at any given time. **Green threads** are created and scheduled by Virtual machine without using OS libraries.

Coroutines are simply concurrent subroutines (functions, closures, or
methods in Go) that are nonpreemptive — that is, they cannot be interrupted.coroutines have multiple points throughout which allow for
suspension or reentry.Goroutines don’t define their own suspension or reentry points;.Go’s
runtime observes the runtime behavior of goroutines and automatically
suspends them when they block and then resumes them when they become
unblocked. Go’s mechanism for hosting goroutines is an implementation of what’s called
an M:N scheduler, which means it maps M green threads to N OS threads.Goroutines are then scheduled onto the green threads. When we have more
goroutines than green threads available, the scheduler handles the distribution
of the goroutines across the available threads and ensures that when these
goroutines become blocked, other goroutines can be run. Go follows a model of concurrency called the fork-join model. 1 The word
fork refers to the fact that at any point in the program, it can split off a child
branch of execution to be run concurrently with its parent. The word join
refers to the fact that at some point in the future, these concurrent branches of
execution will join back together. Where the child rejoins the parent is called
a join point.

If the main goroutine executes then other goroutines will not execute and the program will be terminated. We can put a time.sleep after invoking a go routine. But it will create a race condition only. we need to create join points to remove race condition.

we can use `sync.WaitGroup` for that purpose.

### Waitgroup

`waitgroup is a great way to wait for a set of concurrent operations to
complete when you either don’t care about the result of the concurrent
operation, or you have other means of collecting their results. If neither of
those conditions are true, I suggest you use channels and a select statement`

