package prog

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
