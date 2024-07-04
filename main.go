package main

import "main/intepreter"

func main() {
	// Hello World!
	s := "++++++++[>++++[>++>+++>+++>+<<<<-]>+>+>->>+[<]<-]>>.>---.+++++++..+++.>>.<-.<.+++.------.--------.>>+.>++."

	i := intepreter.CreateInterpreter(s)
	i.Run()
}
