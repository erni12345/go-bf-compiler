package main


import(
    "fmt"
    "os"
    "time"
)

func main() {
    if len(os.Args) < 4 {
        fmt.Println("Usage: go run compiler.go <mode> <input_file> <output_file>")
        return
    }

    mode := os.Args[1]
    filePath := os.Args[2]

    var compileTime, runTime time.Duration

    start := time.Now()

    switch mode {
    case "interpret":
        interpretStart := time.Now()
        Interpret(filePath)
        runTime = time.Since(interpretStart)
    case "compile-to-go":
        outputPath := os.Args[3]
        compileStart := time.Now()
        if err := InterpreterToGo(filePath, outputPath); err != nil {
            fmt.Println("Error:", err)
            return
        }
        compileTime = time.Since(compileStart)

        runStart := time.Now()
        if err := RunCompiledGo(outputPath); err != nil {
            fmt.Println("Error:", err)
            return
        }
        runTime = time.Since(runStart)
    default:
        fmt.Println("Invalid mode. Use 'interpret' or 'compile-to-go'.")
        return
    }

    elapsed := time.Since(start)
    total := compileTime + runTime
    fmt.Printf("Mode: %s, Compile Time: %s, Run Time: %s, Total Time: %s, Elapsed Time: %s\n", mode, compileTime, runTime, total, elapsed)
}

func must(err error){
    if err != nil {
        fmt.Println("Error", err)
        os.Exit(1)
    }
}
