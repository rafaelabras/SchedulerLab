package syscall

import (
	"github.com/rafaelabras/SchedulerLab/process"
)

func exec(p *process.Process, p2 *process.Process) {
	p.Priority = p2.Priority
	p.Tickets = p2.Tickets
	p.CPUTime = p.CPUTime
	p.RemainingInstructions = p2.RemainingInstructions
	p.Runtime = p2.Runtime
}
