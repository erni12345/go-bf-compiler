package main

import (
    "os"
    //"fmt"
)

type Reader struct {
    content []byte
    pointer int
    size int
}

//Make a loader, takes in the file path of the brainfuck file, loads it into the Reader
func LoadFile(path string) *Reader {
    file, err := os.Open(path)
    must(err)
    defer file.Close() //This makes sure that the function is called at the end of the current function

    fileInfo, err := file.Stat()
    must(err)
    fileSize := fileInfo.Size()

    buffer := make([]byte, fileSize) //we allocate memory to hold all the content

    _, err = file.Read(buffer)
    must(err)

    return &Reader{
        content: buffer,
        pointer: int(0),
        size : int(fileSize),
    }

}


//Returns the current instruction
func (r* Reader) CurrentInstruction() byte{
    return r.content[r.pointer]
}

//This function returns the next instruction
func (r *Reader) NextInstruction() {
    r.pointer++

}

//Jump to pointer index
func (r *Reader) JumpTo(n int) {
    r.pointer = n
}

//This function returns the location of the next ']' to be used with the looping logic
//TODO : maybe add a check if the current pos is a '['
func (r *Reader) EndOfLoopIndex() int {
    stack := NewStack()

    for i := r.pointer; i < r.size; i++ {
        if r.content[i] == '[' {
            stack.Push(i)
        } else if r.content[i] == ']' {
            // If stack is empty, then this ']' has no matching '['
            if stack.IsEmpty() {
                return -1
            }
            // Pop the top index from stack
            stack.Pop()

            // If this is the corresponding ']' for the outermost '[', return its index
            if stack.IsEmpty() {
                return i
            }
        }
    }
    // If the stack is not empty, then there are unmatched '['
    if !stack.IsEmpty() {
        return -1
    }

    // No closing ']' found
    return -1
}



