
package main

import (
    "fmt"
    "os"
    "os/exec"

)



func InterpreterToGo(inputFile string, outputFileName string) error {
    input, err := os.ReadFile(inputFile)
    if err != nil {
        return fmt.Errorf("error reading input file: %v", err)
    }

    outputFile, err := os.Create(outputFileName)
    if err != nil {
        return fmt.Errorf("error creating output file: %v", err)
    }
    defer outputFile.Close()

    // Write the package and main function declaration
    outputFile.WriteString("package main\n\n")
    outputFile.WriteString("import (\n")
    outputFile.WriteString("\"fmt\"\n")
    outputFile.WriteString(")\n\n")
    outputFile.WriteString("func main() {\n")
    outputFile.WriteString("    memory := make([]byte, 30000)\n")
    outputFile.WriteString("    pointer := 0\n\n")

    // Parse Brainfuck code and generate corresponding Go code
    for _, char := range string(input) {
        switch char {
        case '>':
            outputFile.WriteString("    pointer++\n")
        case '<':
            outputFile.WriteString("    pointer--\n")
        case '+':
            outputFile.WriteString("    memory[pointer]++\n")
        case '-':
            outputFile.WriteString("    memory[pointer]--\n")
        case '.':
            outputFile.WriteString("    fmt.Print(string(memory[pointer]))\n")
        case ',':
            outputFile.WriteString("    fmt.Scanf(\"%c\", &memory[pointer])\n")
        case '[':
            outputFile.WriteString("    for memory[pointer] != 0 {\n")
        case ']':
            outputFile.WriteString("    }\n")
        }
    }

    // Close the main function
    outputFile.WriteString("}\n")

    return nil
}


func RunCompiledGo(goFile string) error {
    cmd := exec.Command("go", "run", goFile)
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("error running compiled Go code: %v", err)
    }
    return nil
}
