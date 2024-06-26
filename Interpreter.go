package main

import ("fmt")


func Interpret(path string) {

    reader := LoadFile(path)
    tape   := NewTape()
    stack  := NewStack()


    for reader.size > reader.pointer{ //while the pointer isnt at the end (or error thrown)
        current := reader.CurrentInstruction()
        //fmt.Println(string(current))
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

        //fmt.Println(tape.pointer, tape.cells[tape.pointer], string(current))
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
        //fmt.Println("end of pointer", endPointer, string(reader.content[endPointer]))
        if(endPointer != -1){
            reader.JumpTo(endPointer)
            return
        }else{
            fmt.Println("no closing bracket")
        }
    }
    //fmt.Println("made a loop")
    s.Push(reader.pointer)
}

func endOfLoop(s *Stack, tape *Tape, reader *Reader){
    value := tape.GetCurrentValue()

    if value != 0 {
        //we need to jump to the beginning again
        startPointer, success := s.Peek()
        if(success){
            //fmt.Println("start pointer", startPointer, string(reader.content[startPointer]))
            reader.JumpTo(startPointer)
            return
        } else {
            fmt.Println("tried to access empty stack")
        }
    }
    s.Pop()
    //if equal to 0 then we dont do anything
}
