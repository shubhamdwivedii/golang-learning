# OS Scheduler

OS schedulers take into account the existence of multiple processors, chores, CPU caches and NUMA. Without this knowledge, the scheduler can't be  as efficient as possibnle. '

Your program is just a series of machine instructions that need to be executed one after the other sequentially. 

To make this happen OS uses the concept of "Thread". Its the job of the Thread to account for and sequentially execute the set of instructions it's assigned. 

Execution continues until there are no more instructions for the Thread to execute. This is why Threads are called "a path of execution"

### Thread is a path of execution

Every program you run creates a "Process" and each Process is given an initial "Thread". Threads have ability to create more Threads. 

All these Threads run independently of each other and scheduling decisions are made at the Thread level, not at the Process Level. 

#### Threads can run "concurrently" (each taking a turn on an individual core) or in "parallel" (each running at the same time on different cores)

Threads maintain their own "state"  to allow for the safe, local, and independent execution of their instructions. 

The OS scheduler is responsible for making sure cores are not idle if there are Threads that can be executing. It must also create the illusion that all the Threads that can execute are executing at the same time. 

In the process of creating this illusion, the scheduler needs to run Threads with a higher priority over lower priority Threads. However, Threads with a lower priority can't be starved of execution time. 

The schedular also needs to minimize scheduling latencies as much as possible by making quick and smart decisions. 

## Executing Instructions

The "Program Counter (PC)", which is sometimes called the instruction pointer (IP), is what allows the Thread to keep track of the next instruction to execute. 

In most processors, the PC points to the "next instruction" and NOT the current instruction. 

- The computer keeps track of the next line to be executed by keeping its ADDRESS in the special register called the "Instruction Pointer" (IP) or "Program Counter" (PC)

- The contents of IP/PC is updated with every instruction executed. 

- Thus a program is executed sequentially line by line. 

In Go stack trace, the small hexadecimal numbers at the end of each line (like: +0x39 or +0x72) represent the PC/IP value offset from the top of the respective function. 


## Thread States 

Thread can be in three states: "Waiting", "Runnable" or "Executing" 

1. Waiting: This means the Thread is stopped and waiting for something in order to continue. 
    - This could be for reasons like, waiting for the hardware(disk, network)
    - OS calls or synchronization calls (atomic, mutexes)
    
These type of "latencies" are a root cause for bad performance. 

2. Runnable: This means the Thread wants time on a core so it can execute its assigned machine instructions. 
    - If you have a lot of threads that want time, then Threads have to wait longer to get time. 
    - Also, the indivitual amount of time any given Thread gets is shortened, as more Threads compete for time.  

These type of scheduling latency can also cause bad performance. 

3. Executing: This means the Thread has been placed on a core and is executing its machine instructions. 
    - The work related to the application is getting done. 
    - This is what everyone wants. 


## Types Of Work 

There are two types of work a Thread can do: "CPU-Bound" and "IO-Bound" 

1. CPU-Bound: This is work that never creates a situation where the Thread may by placed in Waiting states. 
    - This is work that is constantly making calculations. 
    - A Thread calculating Pi to Nth digit would be CPU-Bound. 

2. IO-Bound: This is work that causes Threads to entro into "Waiting" states. 
    - This is work that consists in requesting access to a resource over the network or making system calls into the operating system.
    - A Thread that needs to access a database would be IO-Bound. It would include synchronization events (mutexes, atomic), that cause the Thread to wait as part of this category. 


## Context Switching 

OS like MacOS, Linux, Windows have a "preemptive scheduler". It means two things: 
1. It means the scheduler is unpredicable when it comes to what Threds will be chosen to run at any given time. 
    - Thread priorities together with event, (like receiveing data on the network) make it impossible to determine what the scheduler will choose to do and when. 

2. It means you must never write code based on some perceived behavior that you have been lucky to experience but is not guranteed to take place every time. 
    - You must control the sychronization and orchestration of Threads if you need determinism in your application. 

#### The physical act of swapping Threads on a core is called a "context switch".

A context switch happens when the scheduler pulls an Executing thread off a core and replaces it with a Runnable Thread. 

    - The Thread that was selected from the run queue moves into an "Executing" state. 

    - The Thread that was pulled can move back into "Runnable" state (if it still has the ability to run), or into a "Waiting" state (if it was replaced because of IO-Bound type of request).


### Context Switches are considered to be expensive because it takes time to swap Threads on and off a core. 

The amount of latency incurred during a context switch depends on different factors but it's not unreasonable for it to take between ~1000 and ~1500 nanoseconds. 

Considering the hardware should be able to reasonably execute (on average) 12 instructions per nanosecond per core, a context switch can cost you ~12k to ~18k instructions of latency. 

In essence, your program is losing the ability to execute a large number of instructions during a context switch. 

If you have a program that is focused on IO-Bound work, then context switches are going to be an advantage.  
    - Once a Thread moves into a Waiting state, another Thread in a Runnable state is there to take its place. 
    - This allows the core to always be doing work. 

#### ^ This is one of the most important aspect of scheduling.  

#### Don't allow cores to go idle if there is work (Threads in a Runnable state) to be done. 

If you program is focused on CPU-Bound work, then context switching are going to be a performance nightmare.  
    - Since Thread always has work to do, the context switch is stopping that work from progressing.  
    - This situation is in stark contrast with what happens with an IO-Bound workload. 


# Less Is More 

Scheduling wasn't overly complicated in early days when processors had only one core, only one Thread could execute at any given time. 

The idea is to define a "Scheduler Period" and attempt to execute all the Runnable Threads withing that period of time. 

If Scheduler Period is 1000ms (1 second) and you have 10 Threads, then each Thread gets 100ms each. 

If you have 100 Thread, each Thread gets 10ms each. 

What if you have 1000 Threads, giving each Thread 1ms doesn't work because the percentage of time you're spending in "context switches" will be significant related to the amount of time you're spending on application work. 

#### We should limit on how small a given time slice can be

- If minimum time slice was 10ms and you have 1000 Threads, the scheduler period needs to be 10000ms (10 seconds)

You should control the number of Threads you use in your application. 

When there are more Threads to consider, and IO-Bound work happening, there is more chaos and nondeterministic behavior.
    - Things take longer to schedule and execute 


#### Thus the rule "Less Is More". 
- Less Threads in a Runnable state means less scheduling overhead and more time each Thread gets over time. 

- More Threads in a Runnable state means less time each Thread gets over time. This means less of your work is getting done over time as well. 


# Find The Balance 

There is a balance you need to find between the number of cores you have and the number of Threads you need to get the best throughput for your application. 

When it comes to managing this balance, "Thread Pools" were a great answer (no longer necessary with Go).

In languages like C++ and C# (or Java?), the user of IOCP (IO Completion Ports) thread pools were critical to writing multithreaded software. 

Ideal number was 3 Threads per Core. 


# Cache Lines

Accessing data from main memory (RAM) has such a high latency cost (~100 to ~300 clock cycles) that processors and cores have local caches to keep data close to the hardware threads that need it.  

Accessing data from caches have a much lower cost (~3 to ~40 clock cycles) depending on the cache being accessed.

Today, one aspect of performance is about how efficiently  you can get data into a processor to reduce these data-access latencies. 

Writing multithreaded applications that mutate state need to consider the mechanics of the caching system. 

##### Data is exchanged between process and main memory (RAM) using "Cache Lines"

- A cache line is a 64-byte chunk of memory that is exchanged between main memory and the caching system. 

- Each core is given its own copy of any cache line it need, which means the harware uses "value semantics" 

- This is why mutations to memory in multithreaded applications can create performance nightmares. 


#### When multiple Threads running in parallel are accessing the same data value or even data values near one another, they will be accessing data on the same cache line. 

#### Any Thread running on any core will get its own copy of that same cache line. 


# False Sharing 

Suppose Core 0 accesss A and Core 1 accesses A+1 

- A and A+1 are independent pieces of memroy, concurrent access is safe. 

- But A and A+1 probably map to the same cache line. 

If Core 0 writes to A invalidates A+1's cache line in Core 1. And vice versa. This is "False Sharing" 

If one Thread on a given core makes a change to its copy of the cache line, then through the magic of hardware, all other copies of the same cache line have to be marked dirty. 

When a Thread attempts read or write access to a dirty cache line, main memory access (~100 to ~300 clock cycles) is required to get a new copy of the cache line.

Maybe on a 2-core processor this isn’t a big deal, but what about a 32-core processor running 32 threads in parallel all accessing and mutating data on the same cache line? 

What about a system with two physical processors with 16 cores each? This is going to be worse because of the added latency for processor-to-processor communication.

The application is going to be thrashing through memory and the performance is going to be horrible and, most likely, you will have no understanding why. 

#### This is called the "Cache-Coherency Problem" and also introduces problems like false sharing. 

- When writing multithreaded applications that will be mutating shared state, the caching systems have to be taken into account.


# Scheduling Decision Scenario

Start application and the main Thread is created and is executing on core 1. 

As the Thread starts executing its instructions, cache lines are retrieved because data is required. 

The Thread now decides to create a new Thread for some concurrent processing. Here is the question: 

Once the Thread is created and ready to go, should the scheduler: 

1. Context-switch the main Thread off of core 1 ? \
    - Doing this could help performance, as the chances that this new Thread needs the same data that is already cached is pretty good. 
    - But the main Thread does not get its full time slice. 

2. Have the Thread wait for core 1 to become available pending the completion of the main Thread's time slice ?
    - The Thread is not running but latency on fetching data will be eliminated once it starts. 

3. Have the Thread wait for the next available core ? 
    - This would mean cache lines of the selected core would be flushed, retrieved, and duplicated, causing latency. 
    - However, the Thread would start more quickly and the main Thread could finish its time slice. 


#### These are the interesting questions that OS Scheduler needs to take into account when making decisions. 

### If there's an idle core, it's going to be used. You want Threads running when they can be running. 