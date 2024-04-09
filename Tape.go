package main

import (
    "fmt"
)


const tapeSize = 30000 //TODO find a way to be smarter about size?

//So we are using a struct which holds an array of bytes and a pointer to simulate the Tape
type Tape struct {
    cells []byte
    pointer int
    tapeSize int
}

//return a pointer because its better than copying each time
func NewTape() *Tape{ //we return a pointer thats why we use the *
    return &Tape{ //the & operator instructs to return a pointer
        cells: make([]byte, tapeSize)   ,
        pointer: 0,
        tapeSize: tapeSize,
    }
}

//+
func (t *Tape) Increment() {
    t.cells[t.pointer]++
    if t.cells[t.pointer] > 255 {
        t.cells[t.pointer] = 0
    }
}

//-
func (t *Tape) Decrement() {
    t.cells[t.pointer]--
    if t.cells[t.pointer] < 0 {
        t.cells[t.pointer] = 255
    }
}

//+/- optimized, idea is we agregate the + and -
func (t *Tape) IncrDecrOpti(r *Reader) {
    counter := 0
    for (r.CurrentInstruction() == '+' || r.CurrentInstruction() == '-') {
        value := r.CurrentInstruction()
        switch value {
            case '+':
                counter++
            case '-':
                counter--
        }
        r.NextInstruction()
    }
    t.cells[t.pointer] += byte(counter)
    r.pointer--
}

// >
func (t *Tape) MoveRight(){
    if(t.tapeSize == t.pointer-1){
        t.tapeSize *= 2
        newCells := make([]byte, len(t.cells), t.tapeSize)
        copy(newCells, t.cells)
	t.cells = newCells
    }
    t.pointer++; //TODO : WHAT IF AT THE END, possible solution -> make tape bigger and copy
}

// <
func (t *Tape) MoveLeft(){
    if(t.pointer != 0){
        t.pointer--; //TODO: what if at pos 0?
    } else{
        fmt.Println("you are trying to go to negative pointer")
    }
}

func (t *Tape) MoveRLOpti(r *Reader){
    counter := 0
    for (r.CurrentInstruction() == '>' || r.CurrentInstruction() == '<'){
        value := r.CurrentInstruction()
        switch value{
            case '>':
                counter++
            case '<':
                counter--
        }
        r.NextInstruction()
    }
    t.pointer += counter
    r.pointer--

}


//. TODO : OPTIMIZE by using a buffer, only print at the end or whenever an input is requested
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


