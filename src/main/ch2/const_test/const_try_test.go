package const_test

import "testing"

const (
	Monday = iota + 1
	Tuesday
	Wednesday
)

const (
	Readable = 1 << iota
	Writeable
	Executable
)

func TestConstantTry(t *testing.T) {
	t.Log(Monday, Tuesday, Wednesday)
}

func TestConstantTryByte(t *testing.T) {
	//a := 7 //0111
	a := 1 //0001
	t.Log(Readable, Writeable, Executable)
	t.Log(a&Readable == Readable, a&Writeable == Writeable, a&Executable == Executable)
}
