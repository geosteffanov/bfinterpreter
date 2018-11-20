package bf

import (
	"bufio"
	"io"
	"os"
	"testing"
)

type flushingWriter struct {
	w *bufio.Writer
}

func (w flushingWriter) Write(p []byte) (int, error) {
	count, err := w.w.Write(p)
	if err != nil {
		return 0, err
	}

	w.w.Flush()

	return count, nil
}

func TestInterpretInstruction(t *testing.T) {
	var writer io.Writer

	writer = flushingWriter{
		w: bufio.NewWriter(os.Stdout),
	}

	src := []instruction{
		{
			idx:             0,
			instructionType: incD,
		},
	}

	for i := 0; i < 65; i++ {
		src = append(src, instruction{
			idx:             uint(i),
			instructionType: incD,
		})
	}

	src = append(src, instruction{
		idx:             65,
		instructionType: outD,
	})

	intepreterState := state{
		src:    src,
		buffer: make([]byte, 1),
		writer: writer,
	}

	writer.Write([]byte("hello\n"))

	for i := 0; i < 66; i++ {
		intepreterState.interpretInstruction()
	}

	intepreterState.interpretInstruction()

}

func TestInterpreter(t *testing.T) {
	input :=  "++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++[.]"
	instructions := parseInput(tokenizeInput(input))
	var writer io.Writer

	writer = flushingWriter{
		w: bufio.NewWriter(os.Stdout),
	}

	state := state{
		instructionPtr: 0,
		src: instructions,
		buffer: make([]byte, 1),
		writer: writer,
	}

	for i := 0; i < 82; i++ {
		state.interpretInstruction()
	}


}
