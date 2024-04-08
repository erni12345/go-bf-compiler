package main

import ("fmt")


func Interpret(path string) {

    reader := LoadFile(path)
    tape   := NewTape()
    stack  := NewStack()


    for reader.size > reader.pointer{ //while the pointer isnt at the end (or error thrown)
        current := reader.CurrentInstruction()
        switch current {
            case '+':
                tape.Increment()
            case '-':
                tape.Decrement()
            case '>':
                tape.MoveRight()
            case '<':
                tape.MoveLeft()
            case '.':
                tape.OutputPointer()
            case ',':
                tape.InputToPointer()
            case '[': //determine if jump to end or not. if not then we need to store a comeback
                startOfLoop(stack, tape, reader)
            case ']':
                endOfLoop(stack, tape, reader)
            default:
        }

        fmt.Println(tape.pointer, tape.cells[tape.pointer], string(current))
        //fmt.Println("before", string(reader.CurrentInstruction()))
        reader.NextInstruction()
        //if(reader.size > reader.pointer){
        //    fmt.Println("after", string(reader.CurrentInstruction()))

        //}
    }

}

func startOfLoop(s *Stack, tape *Tape, reader *Reader){
    value := tape.GetCurrentValue()

    if value == 0{
        endPointer := reader.EndOfLoopIndex()
        fmt.Println("end of pointer", endPointer, string(reader.content[endPointer]))
        reader.JumpTo(endPointer)
        return
    }

    s.Push(reader.pointer)
}

func endOfLoop(s *Stack, tape *Tape, reader *Reader){
    value := tape.GetCurrentValue()

    if value != 0 {
        //we need to jump to the beginning again
        startPointer, _ := s.Pop()
        fmt.Println("start pointer", startPointer, string(reader.content[startPointer]))
        reader.JumpTo(startPointer)
        return
    }

    //if equal to 0 then we dont do anything
}
