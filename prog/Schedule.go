package prog

import (
	"sort"
)

// Schedule describes a list of programs to execute in a specific order.
type Schedule struct {
	programsList []Program
	cursor       uint8
}

// MakeSchedule makes a Schedule for a given list of programs.
func MakeSchedule(programs []Program) Schedule {
	sc := &Schedule{programsList: []Program{}, cursor: 0}
	sc.AddProgram(programs...)

	return *sc
}

// AddProgram adds a program (or programs) to a given Schedule
func (sc *Schedule) AddProgram(program ...Program) {
	sc.programsList = append(sc.programsList, program...)
	sc.Reschedule(1) // set to sort by Job ID
}

// CurrentProgram returns the current program in the schedule
func (sc Schedule) CurrentProgram() *Program {
	return &sc.programsList[sc.cursor]
}

// NextProgram goes to the next program in the schedule and returns it
// when there is not a next progam, it will just return nil
func (sc *Schedule) NextProgram() *Program {
	if sc.cursor < uint8(len(sc.programsList)-1) {
		// there are still programs in the list
		// increment the cursor to go to the next program
		sc.cursor++
		// return the next program at the new cursor
		return &sc.programsList[sc.cursor]
	}
	return nil
}

// Reschedule applies an order to the list of jobs in the jobs table.
func (sc *Schedule) Reschedule(i int) {
	var sb ScheduleSortBy
	switch i {
  case 1:
		// sort by Job ID
		sb = func(p1, p2 Program) bool {
			return p1.Job.ID < p2.Job.ID
		}
  case 2:
		// sort by PriorityNumber
    sb = func(p1, p2 Program) bool {
			return p1.Job.PriorityNumber < p2.Job.PriorityNumber
		}
  }
	ScheduleSortBy(sb).Sort(sc.programsList)
}

// ScheduleSortBy defines how to sort the schedule
type ScheduleSortBy func(p1, p2 Program) bool

// Sort sorts the programs into the correct order.
func (by ScheduleSortBy) Sort(programs []Program) {
	ps := &programSorter{
		programs: programs,
		by:      by,
	}
	sort.Sort(ps)
}

type programSorter struct {
	programs []Program
	by       ScheduleSortBy
}

// Len is part of sort.Interface.
func (s *programSorter) Len() int {
	return len(s.programs)
}

// Swap is part of sort.Interface.
func (s *programSorter) Swap(i, j int) {
	s.programs[i], s.programs[j] = s.programs[j], s.programs[i]
}

// Less is part of sort.Interface.
// It is implemented by calling the "by" closure in the sorter.
func (s *programSorter) Less(i, j int) bool {
	return s.by(s.programs[i], s.programs[j])
}
