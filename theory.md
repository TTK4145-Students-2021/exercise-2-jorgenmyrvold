Exercise 2 - Theory questions
-----------------------------

### What is an atomic operation?
> *Atomic operation means that the operation is done as one instruction. This means that the operation can not be interupted by operations running in other threads or processes and therefor cause concurrency problems*

### What is a critical section?
> *It is a part of the program where data manipulation is done. If a process is interupted in the critical section this can cause problems that corrupt data*

### What is the difference between race conditions and data races?
> *A data race occures when two threads or processes tries to access the same memory location and at least one is a write. A data race can be prevented by locking resources. A race condition is a situation where the order in which threads executes code affects the output. Two threads can write different values to memory but the output is not nessecarily deterministic and can not be solved simply by using locks.*

### What are the differences between semaphores, binary semaphores, and mutexes?
> *Mutexes are locks on resources that can prohibit data races. A mutex lock is owned by the task (thread or process) that takes it, and can oly be unlocked by the same task. Semaphores on the other hand have no ownership. A binary semaphore can be treated somewhat similar to a mutex, but the notion of no ownership means that a task can not know for sure that it is the only task accessing the data as it can with a mutex. A non-binary semaphore can be used to restrict a limited number of tasks to access a resource or used in the implementation of a buffer as in task 3 in the C buffer problem.*

### What are the differences between channels (in Communicating Sequential Processes, or as used by Go, Rust), mailboxes (in the Actor model, or as used by Erlang, D, Akka), and queues (as used by Python)? 
> *Channels, such as in Go, are in many ways similar to mailboxes or queues, but they only have space for one element by default. They can be buffered so they can contain more elements if needed. If there is no space in the channel it blocks. A mailbox by default is unbounded and does not block if any sender tries to send to it it. It can be bounded in the case that the task receives more data than it can process and then the new messages will be bassed to deadletters when it is full. Queues in python are by default not a communication method between processes unless it shares the memory. It can still be used in much the same way if treated correctly.*

### List some advantages of using message passing over lock-based synchronization primitives.
> *In many cases it can be easier to maintain controll over a message passing program rather than a shared variable program. This is patialy due to not needing to keep track of all locks. The message passing system also sends messages when events happen and the receiver can act upon the messages received rather than checking variables all the time to see if something has changed and if it needs to respond in som way*

### List some advantages of using lock-based synchronization primitives over message passing.
> *Although message passing systems has it's advantages, lock-based synchronization primitives also has some advateges. The problem of race conditions can still happen in message passing systems, but this can often be solved simpler with the use of locks. The use of memory can also be reduced by using lock-based primitives instead of having large buffered channels that demands more memory.*