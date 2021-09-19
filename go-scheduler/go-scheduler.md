# Go Scheduler 


When you Go program starts up, it's given a Logical Processonr (P) for every virtual core that is identified on the host machine. 

If you processor has multiple hardware threads per physical core (Hyper-Threading (intel) or Simultaneous Multi-Threading (AMD)), each hardware thread will be presented to your Go program as a virtual core. 

To see number of virtual cores 

import (
    "fmt"
    "runtime
)

func main() {
    // NumCPU returns the number of logical CPUs usable by the current process 
    fmt.Println(runtime.NumCPU(), "Ps") // 16 if 8 core CPU with Multi-Threading. 
}

Every P is assigned an OS Thread("M"). The "M" stands for machine. 

This Thread is still managed by the OS and the OS (OS Scheduler) is still responsible for placing the Thread on a Core for execution. 


Evey Go program is also given an initail Goroutine ("G"), which is the path of execution for a Go program. 

A Goroutine is essentially a "Coroutine" but this is Go, so we replace letter "C" with a "G". 

Goroutines are application-level threads and they are similar to OS Threads in many ways. 

Just as OS Threads are context-switched on and off a core, Goroutines are context-switched on and off an M ("M" is the OS Thread)


# Run Queues 

There are two different run queues in the Go scheduler: 
1. The Global Run Queue (GRQ)
2. The Local Run Queue (LRQ)

Each P is given a LRQ that manages the Goroutines assigned to be executed withing the context of a P (P is Logical Core).

These Goroutines take turns being context-switched on and off the M (OS Thread) assigned to that P (Logical Core) 

The GRQ is for Goroutines that have not been assigned to a P yet. 
There is a process to move Goroutines from GRQ to a LRQ that will discuss later. 

# Cooperating Scheduler 

The OS Scheduler is a preemptive scheduler. You can't predict what the scheduler is going to do at any given time.  

The kernel is making decisions and everything is non-deterministic. 

Applications that run on top of lthe OS have no control over what is happening inside the kernel with scheduling unless they leverage synchronization primitives like "atomic" instructions and "mutex" calls. 

### The Go scheduler is part of the Go runtime, and the Go runtime is build into your application

This means the Go scheduler runs in "user space", above the kernel.  

The current implementation of a Go Scheduler is NOT a "Preemptive" scheduler BUT a "Cooperating" scheduler. 

Being a cooperating scheduler means the scheduler needs well-defined user space events that happen at safe points in the code to make scheduling decisions. 

- What's brilliant about Go cooperating scheduler is that is looks and feels preemptive. 

- You can't predict what the Go scheduler is going to do. 

- This is because decision making for this cooperating scheduler doesn't rest in the hands of developers, but in the Go runtime. 

- It’s important to think of the Go scheduler as a preemptive scheduler and since the scheduler is non-deterministic, this is not much of a stretch.


# Goroutine States 

Just like Threads, Goroutines have the same three high-level states: Waiting, Runnable and Executing. 

- These states dictates the role the Go scheduler takes with any given Goroutine

1. Waiting - This means the Goroutine is stopped and waiting for something in order to continue. 
    - This could be OS calls, synchronization calls (atomic & mutex operations).
    - These types of "latencies" are a root cause for bad performance. 

2. Runnable - This means the Goroutine wants time on an M (OS Thread) so it can execute its assigned instructions. 
    - If you have a lot of Goroutines that want time, then Goroutines have to wait longer to get time. 
    - Also, the individual amount of time any given Goroutine gets is shortened as more Goroutines compete for time. 
    - This type of sheduling latency can also be a cause of bad performance. 


3. Executing - This means the Goroutine has been placed on an M (OS Thread) and is executing its instructions. 
    - The work related to the application is getting done. 
    - This is what everyone wants.
