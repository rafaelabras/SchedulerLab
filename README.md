# CPU Virtualization Simulator (OSTEP-inspired)

This project is a **CPU virtualization simulator** inspired by the book
*Operating Systems: Three Easy Pieces (OSTEP)*.

It models how an operating system:
- creates processes
- executes instructions
- handles traps and system calls
- performs context switches
- schedules the CPU using different algorithms

The simulator is **fully sequential** (no threads or concurrency) and focuses
on **mechanisms and policies**, not performance or real execution.

---

##  Goals

- Understand **Limited Direct Execution**
- Model **process states and lifecycle**
- Simulate **system calls** (`fork`, `exit`, `yield`, `wait`)
- Implement classic **CPU scheduling algorithms**
- Build intuition about how an OS controls the CPU

---

##  OSTEP Chapters Covered

- Chapter 4 – Processes
- Chapter 5 – Process API
- Chapter 6 – Limited Direct Execution
- Chapter 7 – CPU Scheduling
- Chapter 8 – Multi-Level Feedback Queue (MLFQ)
- Chapter 9 – Lottery Scheduling
- Chapter 10 – Multi-CPU Scheduling (simulated, single-threaded)

---

##  Core Concepts

### Process
Each process is a data structure that represents a running program.
It contains:
- PID
- State (`READY`, `RUNNING`, `BLOCKED`, `EXIT`)
- CPU time used
- Priority
- Lottery tickets
- Remaining instructions

Processes do **not** run real code — each instruction is simulated.

---

### Limited Direct Execution
Processes execute instructions directly until:
- a **system call** occurs (trap)
- the **time quantum** expires
- the process exits

When a trap happens, control is transferred back to the scheduler,
which decides what to run next.

---

### System Calls (Simulated)

The simulator models classic system calls:
- `fork()` – creates a child process
- `exit()` – terminates the process
- `yield()` – voluntarily gives up the CPU
- `wait()` – blocks until a child exits

No real OS syscalls are used; everything is conceptual.

---

### Scheduling Algorithms

Implemented or planned schedulers:
- FIFO (First-Come, First-Served)
- Round Robin
- Multi-Level Feedback Queue (MLFQ)
- Lottery Scheduling

Each scheduler selects the next process based on its policy.

---

##  Execution Model

The system runs in discrete **ticks**.
At each tick:
1. The scheduler selects a process
2. One instruction is executed
3. The process state may change
4. Context switches are simulated

Example output:

```
TICK 10] PID 3 RUNNING
[TICK 11] PID 3 SYSCALL yield
[TICK 12] PID 5 RUNNING
```


---

## ️ Design Philosophy

- No concurrency or threads
- Deterministic and easy to debug
- Focused on **learning OS internals**
- Code clarity over optimization

---

##  Future Extensions

- Preemption via timer interrupts
- Simulated multi-core CPU
- Per-process instruction streams
- Integration with a virtual memory simulator

---

##  References

- *Operating Systems: Three Easy Pieces* — Arpaci-Dusseau

