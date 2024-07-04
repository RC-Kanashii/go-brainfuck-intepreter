package intepreter

type Interpreter struct {
	tape    *tape
	program string
	pc      int         // program counter
	s       Stack       // s records the position of '['
	jt      map[int]int // jt means jump table, which records the position of matching '[' and ']'
}

func (i *Interpreter) createJumpTable() {
	for idx, instr := range i.program {
		if instr == '[' {
			i.s.Push(idx)
		} else if instr == ']' {
			lp := i.s.Pop()
			i.jt[lp] = idx
			i.jt[idx] = lp
		}
	}
}

func (i *Interpreter) Run() {
	for i.pc < len(i.program) {
		instr := i.program[i.pc]
		cell := i.tape.Get()
		switch instr {
		case '.':
			i.tape.Output()
		case ',':
			i.tape.Input()
		case '+':
			i.tape.Incr()
		case '-':
			i.tape.Decr()
		case '>':
			i.tape.MoveRight()
		case '<':
			i.tape.MoveLeft()
		case '[':
			if cell == 0 {
				// pc jumps to the position after the matching ']'
				i.pc = i.jt[i.pc] // pc will be incremented by 1 after the switch statement
			}
		case ']':
			if cell != 0 {
				// pc jumps to the position of the matching '['
				i.pc = i.jt[i.pc] // pc will be incremented by 1 after the switch statement
			}
		}
		i.pc++
	}
}

func CreateInterpreter(program string) *Interpreter {
	i := Interpreter{
		tape:    createTape(32),
		program: program,
		pc:      0,
		s:       Stack{},
		jt:      make(map[int]int),
	}
	i.createJumpTable()
	return &i
}
