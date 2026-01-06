package syscall

import (
	"github.com/rafaelabras/SchedulerLab/process"
)

type ProcessFunc func(number int, number2 int)

func wait(p *process.Process, parentFunc ProcessFunc, childrenFunc ProcessFunc, number int, number2 int) {
	childrenFunc(number, number2)
	parentFunc(number, number2)
}
