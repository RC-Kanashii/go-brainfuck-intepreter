package intepreter

type Interpreter struct {
	tape    *tape
	program string
	pc      int         // program counter
	ls      Stack       // ls records the position of '['
	rs      Stack       // rs records the position of ']'
	jt      map[int]int // jt means jump table, which records the position of matching '[' and ']'
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
			if !i.ls.Has(i.pc) {
				i.ls.Push(i.pc)
			}
			if cell == 0 {
				// pc jumps to the position after the matching ']'
				i.pc = i.rs.Pop() // pc will be incremented by 1 after the switch statement
				i.ls.Pop()
				i.rs.Pop()
			}
		case ']':
			if !i.rs.Has(i.pc) {
				i.rs.Push(i.pc)
			}
			if cell != 0 {
				// pc jumps to the position of the matching '['
				i.pc = i.ls.Top() // pc will be incremented by 1 after the switch statement
			} else {
				i.ls.Pop()
				i.rs.Pop()
			}
		}
		i.pc++
	}
}

func CreateInterpreter(program string, tapeLength int) *Interpreter {
	return &Interpreter{
		tape:    createTape(tapeLength),
		program: program,
		pc:      0,
		ls:      Stack{},
		rs:      Stack{},
		jt:      make(map[int]int),
	}
}
