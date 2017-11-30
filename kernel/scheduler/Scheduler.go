package scheduler

import (
	"container/heap"
	"fmt"
	"io"
	"log"
	"sort"

	"../../config"
	"../../util/logger"
	"../../vm/ivm"
	"../pageManager"
	"../process"
	"../program"
)

// Scheduler keeps track of system processes.
type Scheduler struct {
	processList        *processList
	processUnloadQueue *processQueue
	pm                 *pageManager.PageManager
	methodName         string
	longTermQueueSize  uint
	logger             *log.Logger
}

// New creates a new scheduler.
func New(c config.Config, p *pageManager.PageManager, a []program.Program) *Scheduler {
	sched := &Scheduler{
		processList: &processList{
			base:       process.MakeArray(a),
			sortMethod: MethodForSwitch(c.Sched),
			logger:     logger.New("processList"),
		},
		processUnloadQueue: &processQueue{
			base:   []uint8{},
			logger: logger.New("processUnloadQueue"),
		},
		pm:                p,
		methodName:        c.Sched,
		longTermQueueSize: c.QSize,
		logger:            logger.New("scheduler"),
	}

	// sort the whole thing
	sort.Sort(sched.processList)

	// make sure each process is set up with the page manager
	sched.Each(func(p *process.Process) {
		err := sched.pm.Setup(p)
		if err != nil {
			sched.logger.Panicf("[Scheduler] New() error: %v", err)
		}
		(*p).SetStatus(process.Ready)
		sched.Update(*p)
	})
	return sched
}

// Tick is used to signal the start of a virtual machine cycle to the kernel.
// This sets up processes and resources before the next cycle begins.
func (sched Scheduler) Tick() {
	sched.logger.Printf("Tick!")

	// run the long-term scheduler
	sched.Long()
}

// Tock is used to signal the end of a virtual machine cycle to the kernel.
// This reacts to the events that occured during the cycle.
func (sched Scheduler) Tock() error {
	sched.logger.Printf("Tock!")

	// make sure terminated processes aren't taking up space anymore
	// (otherwise, there's nothing to fill here and it just stops)
	err := sched.Clean()
	if err != nil {
		sched.logger.Printf("ERROR in Tock: %v", err)
		return err
	}

	// check back through previous requests and try to fulfill them
	sched.pm.HandleWaitlist()

	// make sure any waiting processes have what they need
	sched.Each(func(p *process.Process) {
		if p.Status() == process.Wait {
			err := sched.pm.Reallocate(p)
			if err != nil {
				sched.logger.Printf("ERROR in Tock: %v", err)
			} else {
				p.SetStatus(process.Ready)
			}
		}
	})
	return nil
}

// ProcessForCore returns the appropriate process for the given core.
func (sched Scheduler) ProcessForCore(corenum uint8) process.Process {

	// Look for the first process that is ready to be run
	p := sched.FindBy(func(p *process.Process) bool {
		return p.Status() == process.Ready
	})
	if p == nil {
		sched.logger.Printf("ProcessForCore(%d) => SLEEP", corenum)
		return process.Sleep()
	}
	// make sure the process is ready to be run on the given core
	sched.Short(corenum, p)
	// update this internally (because Short changed it)
	sched.Update(*p)

	sched.logger.Printf(
		"ProcessForCore(%d) => process %d",
		corenum, p.ProcessNumber,
	)
	return *p
}

// Update updates an existing process in the list.
func (sched *Scheduler) Update(p process.Process) error {
	for i := sched.processList.Len() - 1; i >= 0; i-- {
		pX := sched.processList.base[i]
		if p.ProcessNumber == pX.ProcessNumber {
			sched.processList.base[i] = p
			sched.logger.Printf(
				"updated process %d",
				p.ProcessNumber,
			)
			return nil
		}
	}
	err := NotFoundError{}
	sched.logger.Printf(
		"ERROR updating process %d: %v",
		p.ProcessNumber, err,
	)
	return err
}

// Load makes sure the given process is in RAM.
func (sched *Scheduler) Load(p *process.Process) error {
	if p.Status() != process.Ready {
		// this process is not ready to be loaded into RAM
		err := NotReadyError{}
		sched.logger.Printf(
			"ERROR loading process %d: %v",
			p.ProcessNumber, err,
		)
		return err
	}
	// defer to the page manager
	err := sched.pm.Load(p)
	if err != nil {
		sched.logger.Printf(
			"ERROR loading process %d: %v",
			p.ProcessNumber, err,
		)
		return err
	}
	return nil
}

// Save makes sure the given process's RAM is persisted to Disk.
func (sched *Scheduler) Save(p *process.Process) error {
	sched.logger.Printf("save process %d", p.ProcessNumber)

	// defer to the page manager
	err := sched.pm.Save(p)
	if err != nil {
		sched.logger.Printf(
			"ERROR saving process %d: %v",
			p.ProcessNumber, err,
		)
	}
	return err
}

// Unload makes sure the given process is not in RAM.
func (sched *Scheduler) Unload(p *process.Process) error {
	sched.logger.Printf(
		"process %d should be unloaded",
		p.ProcessNumber,
	)

	if p.Status() != process.Terminated {
		sched.logger.Panicf(
			"process %d is not terminated (is %v)",
			p.ProcessNumber, p.Status,
		)
	}

	// defer to the page manager
	err := sched.pm.Unload(p)
	if err != nil {
		sched.logger.Printf(
			"ERROR unloading process %d: %v",
			p.ProcessNumber, err,
		)
		return err
	}
	return err
}

// Complete completes the given processs (marks it Terminated).
func (sched *Scheduler) Complete(p *process.Process) error {
	// mark is Terminated (this will get cleaned up later)
	p.SetStatus(process.Terminated)
	sched.logger.Printf("process %d completed/terminated", p.ProcessNumber)
	return nil
}

// NotFoundError is when the desired process is not in the list.
type NotFoundError struct{}

func (err NotFoundError) Error() string {
	return "process is not in the scheduler"
}

// NotReadyError is when the process is not ready to load RAM.
type NotReadyError struct{}

func (err NotReadyError) Error() string {
	return "process is not ready to load RAM"
}

// Each goes through each process in order and passes to the given function.
func (sched Scheduler) Each(fn func(*process.Process)) {
	for i := sched.processList.Len() - 1; i >= 0; i-- {
		fn(&sched.processList.base[i])
	}
}

// EachWithError goes through each process and checks for an error each time.
func (sched Scheduler) EachWithError(fn func(*process.Process) error) error {
	for i := sched.processList.Len() - 1; i >= 0; i-- {
		if err := fn(&sched.processList.base[i]); err != nil {
			return err
		}
	}
	return nil
}

// EachWhile goes through each process while the function keeps returning true.
func (sched Scheduler) EachWhile(fn func(*process.Process) bool) {
	for i := sched.processList.Len() - 1; i >= 0; i-- {
		if !fn(&sched.processList.base[i]) {
			break
		}
	}
}

// FindBy goes through each until the passed function returns true.
func (sched Scheduler) FindBy(fn func(*process.Process) bool) *process.Process {
	_, p := sched.findPair(fn)
	return p
}

func (sched Scheduler) findPair(fn func(*process.Process) bool) (int, *process.Process) {
	for i := sched.processList.Len() - 1; i >= 0; i-- {
		p := &sched.processList.base[i]
		if fn(p) {
			return i, p
		}
	}
	return -1, nil
}

// FprintProcessTable prints the process table to the given writer.
func (sched Scheduler) FprintProcessTable(w io.Writer) error {
	processListLen := sched.processList.Len()
	header := fmt.Sprintf(
		"Process Table (%d processes, sort method: %s, queue size: %d/%d)\n",
		processListLen, sched.methodName,
		sched.longTermQueueSize, ivm.RAMNumFrames,
	)
	if _, err := w.Write([]byte(header)); err != nil {
		return err
	}
	if err := sched.processList.fprint(w); err != nil {
		return err
	}
	return nil
}

// IsDone returns if the system is done yet.
func (sched Scheduler) IsDone() bool {
	return sched.processList.Len() == 0
}

// NumLeft is the number of processes left in the queue.
func (sched Scheduler) NumLeft() int {
	return sched.processList.Len()
}

// Add adds a process into the process manager.
func (sched Scheduler) Add(p process.Process) {
	heap.Push(sched.processList, p)
}

// Find returns the process with the corresponding process number.
func (sched Scheduler) Find(processNumber uint8) *process.Process {
	_, p := sched.findPair(func(p *process.Process) bool {
		return p.ProcessNumber == processNumber
	})
	return p
}

// ProcessTable returns the system process table
func (sched Scheduler) ProcessTable() []process.Process {
	return sched.processList.base
}
