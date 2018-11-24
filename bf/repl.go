package bf

import (
	"fmt"
	"bufio"
	"os"
)

type repl struct {
	interpreter *state
	reader *bufio.Reader
	writer *bufio.Writer
}


func NewRepl() repl {
	intepreter := NewInterpreter("")
	return repl{
		interpreter: &intepreter,
		reader: bufio.NewReader(os.Stdin),
		writer: bufio.NewWriter(os.Stdout),
	}
}



func (r *repl) Read() {
	fmt.Print(">>")
	instructions, err := r.reader.ReadString('\n')
	instructions = instructions[:len(instructions) - 1]
	if err != nil {
		panic("not implemented")
	}

	r.interpreter.addInstructions(instructions)
}

func (r *repl) Eval() {
		Run(r.interpreter)
}

func (r *repl) Start() {
	for {
		r.Read()
		r.Eval()
	}
}
