# Brainfuck Compiler

This project is a Brainfuck compiler written in Go. The purpose of this project is to improve my Go programming skills while exploring the intricacies of the Brainfuck programming language.

## About Brainfuck

Brainfuck is an esoteric programming language created in 1993 by Urban Müller. It is known for its minimalism and extreme simplicity. Despite its minimalistic nature, Brainfuck is Turing complete, meaning it can compute anything that a Turing machine can. The language consists of only eight commands:

- `>`: Move the pointer to the right
- `<`: Move the pointer to the left
- `+`: Increment the memory cell at the pointer
- `-`: Decrement the memory cell at the pointer
- `[`: Jump past the matching `]` if the cell at the pointer is 0
- `]`: Jump back to the matching `[` if the cell at the pointer is nonzero
- `.`: Output the character at the pointer
- `,`: Input a character and store it in the cell at the pointer

Despite its simplicity, Brainfuck can be used to write complex programs.

## How to Compile Brainfuck Programs

To compile a Brainfuck program, you first need to compile the Go program included in this repository. Follow these steps:

1. Clone this repository to your local machine:
```
https://github.com/erni12345/go-bf-compiler.git
```

2. Navigate to the project directory:
```
cd brainfuck-compiler
```

3. Compile the Go program:
```
go build
```

This will generate an executable file named `brainfuck-compiler` or `brainfuck-compiler.exe` depending on your operating system.

4. Now, you can use the generated executable to compile Brainfuck programs. Run the following command:
```
./brainfuck-compiler <name-of-file.bf> <mode>
```
Current modes : `interpret`
Replace `<name-of-file.bf>` with the filename of your Brainfuck program, and mode with the mode. **Soon** this will compile the Brainfuck program into executable code. For now just runs.

That's it! You've successfully compiled a Brainfuck program using this compiler.

