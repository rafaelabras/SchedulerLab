package syscall

import (
	"github.com/rafaelabras/SchedulerLab/process"
)

func fork(p *process.Process, p2 *process.Process) {
	p.ChildrenPids[0] = p2.PID
	p2.ParentPID = p.PID
}
