package cpu

import (
	"github.com/rafaelabras/SchedulerLab/process"
	"github.com/rafaelabras/SchedulerLab/scheduler"
)

type Cpu struct {
	Tick           int
	CurrentProcess *process.Process
	Scheduler      scheduler.Scheduler
	Quantum        int // time slice
}

func NewCpu() *Cpu {
	return &Cpu{}
}

func (cpu *Cpu) SetupCpu(tick int, quantum int, s scheduler.Scheduler) {
	cpu.Tick = tick
	cpu.Scheduler = s
	cpu.Quantum = quantum
}

func (cpu *Cpu) RunDefault() {
	cpu.Tick++
	flagTimeslice := 0

	if cpu.CurrentProcess.Status == 1 {

		cpu.CurrentProcess.Status = 2
		for flagTimeslice < cpu.Quantum {
			cpu.CurrentProcess.RunInstruction()
			flagTimeslice++
			cpu.Tick++
		}
	}
}
