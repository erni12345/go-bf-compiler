package main


import(
    "fmt"
    "os"
    "time"
    "flag"
)

func main() {



    mode := flag.String("mode", "interpret", "the mode for compilation/interpretation (optional)")
    filePath := flag.String("inputPath", "examples/hanoi.bf", "the path to the BF file (not optional)")
    outputPath := flag.String("outputPath", "output.go", "The output path if compiled (optional)")
    flag.Parse()

    var compileTime, runTime time.Duration



    start := time.Now()

    switch *mode {
    case "interpret":
        interpretStart := time.Now()
        Interpret(*filePath)
        runTime = time.Since(interpretStart)
    case "interpret-opti":
        interpretStart := time.Now()
        InterpreterOpti(*filePath)
        runTime = time.Since(interpretStart)
    case "compile-to-go":
        compileStart := time.Now()
        if err := InterpreterToGo(*filePath, *outputPath); err != nil {
            fmt.Println("Error:", err)
            return
        }
        compileTime = time.Since(compileStart)

        runStart := time.Now()
        if err := RunCompiledGo(*outputPath); err != nil {
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
    fmt.Printf("Mode: %s, Compile Time: %s, Run Time: %s, Total Time: %s, Elapsed Time: %s\n", *mode, compileTime, runTime, total, elapsed)
}

func must(err error){
    if err != nil {
        fmt.Println("Error", err)
        os.Exit(1)
    }
}
