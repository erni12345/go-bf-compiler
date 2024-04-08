package main

import (
    "fmt"
)


const tapeSize = 30000 //TODO find a way to be smarter about size?

//So we are using a struct which holds an array of bytes and a pointer to simulate the Tape
type Tape struct {
    cells []byte
    pointer int
}

//return a pointer because its better than copying each time
func NewTape() *Tape{ //we return a pointer thats why we use the *
    return &Tape{ //the & operator instructs to return a pointer
        cells: make([]byte, tapeSize)   ,
        pointer: 0,
    }
}

//+
func (t *Tape) Increment () { //hacing the (t *Tape) makes it associated with the Tape struct, meaning i can then call tape.Increment()
    t.cells[t.pointer]++;
}

//-
func (t *Tape) Decrement() {
    t.cells[t.pointer]--;
}

// >
func (t *Tape) MoveRight(){
    t.pointer++; //TODO : WHAT IF AT THE END, possible solution -> make tape bigger and copy
}

// <
func (t *Tape) MoveLeft(){
    t.pointer--; //TODO: what if at pos 0?
}

//.
func (t *Tape) OutputPointer(){
    s := string(t.cells[t.pointer])
    fmt.Print(s);
}

//,
func (t *Tape) InputToPointer(){
    var input byte;
    fmt.Scanf("%c\n", &input)
    t.cells[t.pointer] = input
}

func (t *Tape) GetCurrentValue() byte{
    return t.cells[t.pointer];
}


