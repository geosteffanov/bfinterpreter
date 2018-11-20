package bf

import (
	"errors"
	"io"
)

type instruction struct {
	idx             uint
	instructionType token

	matchingBracketIdx uint
}

type state struct {
	mem            memoryState
	src            []instruction
	instructionPtr uint

	buffer []byte
	reader io.Reader
	writer io.Writer
}

func (s *state) moveRight() {
	s.mem.incrementPtr()
}

func (s *state) moveLeft() {
	s.mem.decrementPtr()
}

func (s *state) incrementValue() {
	s.mem.incrementCell()
}

func (s *state) decrementValue() {
	s.mem.decerementCell()
}

func (s *state) readCellValue() error {
	count, err := s.reader.Read(s.buffer)
	if err == io.EOF {
		return nil
	}

	if err != nil {
		return errors.New("couldn't read from input")
	}

	if count < 1 {
		return errors.New("couldn't read from input")
	}

	s.mem.setCell(s.buffer[0])

	return nil
}

func (s *state) writeCellValue() error {
	count, err := s.writer.Write([]byte{s.mem.getCell()})

	if err != nil {
		return errors.New("couldn't write to output")
	}

	if count != 1 {
		return errors.New("couldn't write to output")
	}

	return nil
}

func (s *state) interpretInstruction() {
	instr := s.src[s.instructionPtr]

	switch instr.instructionType {
	case movR:
		s.moveRight()

	case movL:
		s.moveLeft()
	case incD:
		s.incrementValue()
	case decD:
		s.decrementValue()
	case outD:
		s.writeCellValue()
	case inD:
		s.readCellValue()
	case loopS:
		if s.mem.getCell() == 0 {
			s.instructionPtr = instr.matchingBracketIdx
		}
	case loopE:
		if s.mem.getCell() != 0 {
			s.instructionPtr = instr.matchingBracketIdx
		}
	}
	s.instructionPtr += 1
}
