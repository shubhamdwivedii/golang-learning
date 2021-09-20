# Go Scheduler 


When you Go program starts up, it's given a Logical Processor (P) for every virtual core that is identified on the host machine. 

If you processor has multiple hardware threads per physical core (Hyper-Threading (intel) or Simultaneous Multi-Threading (AMD)), each hardware thread will be presented to your Go program as a virtual core. 

To see number of virtual cores 
```Go
import (
    "fmt"
    "runtime
)

func main() {
    // NumCPU returns the number of logical CPUs usable by the current process 
    fmt.Println(runtime.NumCPU(), "Ps") // 16 if 8 core CPU with Multi-Threading. 
}
```

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

![GRQ-LRQ](https://www.ardanlabs.com/images/goinggo/94_figure2.png)


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


# Context Switching 

The Go scheduler requries well-defined user-space events that occur at safe points in the code to context-switch from. 

These events and safe points manifest themselves within function calls. Function calls are critical to the health of the Go Scheduler. 

Today (with Go 1.11 or less), if you run any tight loops that are not making function calls, you will cause latencies within the scheduler and garbage collection.

It’s critically important that function calls happen within reasonable timeframes.


There are four classes of events that occur in your Go programs that allow the scheduler to make scheduling decisions. This doesn't mean it will always happen on one of these events. It means the scheduler gets the opportunity: 

- The use of the keyword "go" 
- Garbage collection 
- System calls 
- Synchronization and Orchestration 

1. The Keyword "go": 
    - The keyword "go" is how you create Goroutines.
    - Once a new Goroutine is created, it gives the scheduler an opportunity to make a scheduling decision. 

2. Garbage collection (GC): 
    - Since the GC runs using its own set of Goroutines, those Goroutines need time on an M (OS Thread) to run. 
    - This causes the GC to create a lot of scheduling chaos. 
    - However, the scheduler is very smart about what a Goroutine is doing and it will leverage that intelligence to make smart decisions. 
    - Once smart decision is context-switching a Goroutine that wants to touch the heap with those that don't touch the heap during GC. 
    - When GC is running, a lot of scheduling decisions are being made. 

3. System Calls: 
    - If a Goroutine makes a system call that will cause the Goroutine to block the M (OS Thread), sometimes the scheduler is capable of context-switching the Goroutine off the M and context-switch a new Goroutine onto that same M. 
    - However, sometimes a new M (OS Thread) is required to keep executing Goroutines that are queued up in the P (Logical Core)

4. Synchronization and Orchestration: 
    If an atomic, mutex, or channel operation call will cause the Goroutine to block, the scheduler can context-switch a new Goroutine to run. Once The Goroutine can run again, it can be re-queued and eventually context-switched back on an M (OS Thread)



# Asynchronous System Calls

When the OS you are running on has the ability to handle a "system call" asynchronously, somthing called the "Network Poller" can be used to process the system call more efficiently. 

This is accomplished by using a "kqueue" (MacOS), "epoll" (Linux) or "iocp" (Windows) within these respective OS's. 


- Networking-based system calls can be process asynchronously by many of the OSs we use today. 

- This is where the "Network Poller" gets its name, since its primary use is handling networking operations.  

- By using the network poller for networking system calls, the scheduler can prevent Goroutines from blocking the M (OS Thread) when those system calls are made. 

- This helps the M available to execute other Goroutines in the P's LRQ without the need to create new Ms (OS thread). 

- This helps to reduce scheduling load on the OS. 


![network-poller1](https://www.ardanlabs.com/images/goinggo/94_figure3.png)

Say Goroutine-1 is executing on the M and there are 3 more Goroutines waiting in the LRQ to get their time on the M. The Network Poller is idle with nothing to do. 

![network-poller2](https://www.ardanlabs.com/images/goinggo/94_figure4.png)

Goroutine-1 wants to make a network system call, so Goroutine-1 is moved to the network poller and the asynchronous network system call is processed. 

Once Goroutine-1 is moved to the network poller, the M is now avaialble to execute a different Goroutine from the LRQ (Local Run Queue). 

In this case, Goroutine-2 is context-switched on the M. 

![network-poller3](https://www.ardanlabs.com/images/goinggo/94_figure5.png)

When the async network system call is completed by the network poller, Goroutine-1 is moved back into the LRQ for the P.  

Once Goroutine-1 can be context-switched back on the M, the Go related code it's responsible for can execute again. 


#### The big win here is that, to execute network system calls, no extra Ms are needed. 
#### The network poller has an OS Thread and it is handling an efficient event loop.


# Synchronous System Calls

What happens when the Goroutine wants to make a system call that can't be done asynchronously ? 

In this case, the network poller can't be used and the Goroutine makeing the system call is going to block the M (OS Thread).

This is unfortunate but there's no way to prevent this from happening. 

One example of such system call is "file-based" system calls. 

If you are using CGO, there may be other situations where calling C functions will block the M (OS Thread) as well. 

**Note: The Windows OS does have the capability of making file-based system calls asynchronously. Technically when running on Windows, the network poller can be used.**

Here Goroutine-1 is going to make a synchronous system call that will block M1

![sync-sys-call1](https://www.ardanlabs.com/images/goinggo/94_figure6.png)

The Scheduler is able to identify that Goroutine-1 has caused the M to block. At this point the scheduler detaches M1 from P (Logical Core) with the blocking Goroutine-1 still attached. 

Then the Scheduler brings in a new M2 (OS Thread) to service the P. 

![sync-sys-call2](https://www.ardanlabs.com/images/goinggo/94_figure7.png)

At that point, Goroutine-2 can be selected from the LRQ and context-switched on M2. If an M (OS Thread) already exists because of a previous swap, this transition is quicker than having to create a new M. 

![sync-sys-call3](https://www.ardanlabs.com/images/goinggo/94_figure8.png)

The blocking system call that was made by Goroutine-1 finishes. At this point, Goroutine-1 can move back into the LRQ and be serviced by the P again. 

M1 is then placed on the side for future use if this scenario needs to happen again. 


# Work Stealing 

Another aspect of the scheduler is that it's a work-stealing scheduler. 

This helps in a few areas to keep scheduling efficient. 

For one, the last thing you want is an M to move into a waiting state because, once that happens, the OS will context-switch the M off the Core (P).

This means the P can't get any work done, even if there is a Goroutine in a runnable state, until an M (OS Thread) is context-switched back on a Core. 

The work stealing also helps to balance the Goroutines across all the P's so the work is better distributed. 

![work-steal1](https://www.ardanlabs.com/images/goinggo/94_figure9.png)

We have a multi-threaded Go program with two P's servicing four Goroutines each and a single Goroutine in the GRQ. 

![work-steal2](https://www.ardanlabs.com/images/goinggo/94_figure10.png)

When P1 ha no more Goroutines to execute. But there are Goroutines in a runnable state, both in the LRQ and P2 and in the GRQ. 

This is the moment where P1 needs to steal work. 

The rules for "Stealing Work" are as follow: 

```Go
runtime.schedule() {
    // only 1/61 of the time, check the global runnable queue for a G (Goroutine)
    // if not found, check the local queue 
    // if not found, 
        // try to steal from other Ps. 
        // if not, check the global runnable queue. 
        // if not found, poll network. 
}
```

Based on these rules, P1 needs to check P2 for Goroutines in its LRQ and TAKE HALF of what it finds. 

![work-steal3](https://www.ardanlabs.com/images/goinggo/94_figure11.png)

P1 has taken half the Goroutines from P2 and now P1 can execute them. 

What happens if P2 finishes servicing all of its Goroutines and P1 has nothing left in its LRQ ? 

![work-steal4](https://www.ardanlabs.com/images/goinggo/94_figure12.png)

P2 (finished all its work) now needs to steal some work. 

First it will look at the LRQ of P1 but it won't find any Goroutines. 

Next, it will look at the GRQ. There it will find Goroutine-9

![work-steal5](https://www.ardanlabs.com/images/goinggo/94_figure13.png)

P2 steals Goroutine-9 from GRQ and begins to execute the work. 

What's great about "work stealing" is that it allows the Ms to stay busy and not go idle. 

**This work stealing is considered internally as "Spinning" the M (OS Thread)**


# Practical Example