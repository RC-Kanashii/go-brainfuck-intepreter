package main

import (
	"flag"
	"main/intepreter"
	"os"
)

func main() {
	file := flag.String("f", "", "file path of the source code")
	flag.Parse()

	if *file == "" {
		panic("file path is required")
	}
	program, err := os.ReadFile(*file)
	if err != nil {
		panic(err)
	}

	i := intepreter.CreateInterpreter(string(program))
	i.Run()
}
