package process

type Status int

const (
	NEW Status = iota
	READY
	RUNNING
	BLOCKED
	EXIT
)

type Process struct {
	PID      int
	Status   int
	Priority int
	Tickets  int
	Runtime  int
}

type ProcessTable map[int][]int

func InitializeProcessTable() ProcessTable {
	return make(ProcessTable)
}

func CreateProcess(pid int, status int, priority int, tickets int, runtime int) Process {
	return Process{
		PID:      pid,
		Status:   status,
		Priority: priority,
		Tickets:  tickets,
		Runtime:  runtime,
	}
}

func (pt ProcessTable) AddProcess(process *Process) {
	processSlice := make([]int, 4)
	processSlice[0] = process.Status
	processSlice[1] = process.Priority
	processSlice[2] = process.Tickets
	processSlice[3] = process.Runtime

	pt[process.PID] = processSlice
}

func (pt ProcessTable) UpdateProcess(process *Process) {
	pt[process.PID][process.Priority] = process.Priority
	pt[process.PID][process.Status] = process.Status
	pt[process.PID][process.Runtime] = process.Runtime
	pt[process.PID][process.Tickets] = process.Tickets
}

func (pt ProcessTable) RemoveProcess(process *Process) {
	pt[process.PID] = nil
}
