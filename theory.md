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
> *Your answer here*

### List some advantages of using message passing over lock-based synchronization primitives.
> *By using message passing you eliminate lots of the challanges with data races and syncronization of variables.*

### List some advantages of using lock-based synchronization primitives over message passing.
> *Your answer here*