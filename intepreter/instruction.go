package intepreter

type Instruction func(*tape)

var InstrSet = map[byte]Instruction{
	'>': (*tape).MoveRight,
	'<': (*tape).MoveLeft,
	'+': (*tape).Incr,
	'-': (*tape).Decr,
	'.': (*tape).Output,
	',': (*tape).Input,
}
