package main



//LOL seems to be even slower to do it this way
func InterpreterOpti(path string) {

    reader := LoadFile(path)
    tape   := NewTape()
    stack  := NewStack()

    for reader.size > reader.pointer{
        current := reader.CurrentInstruction()
        switch current {
            case '+':
                tape.IncrDecrOpti(reader)
            case '-':
                tape.IncrDecrOpti(reader)
            case '>':
                tape.MoveRLOpti(reader)
            case '<':
                tape.MoveRLOpti(reader)
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
        reader.NextInstruction()

    }


}
