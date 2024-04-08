package main


import(
    "fmt"
    "os"
    "time"
)

func main() {
    // Checking if we have a file name
    if len(os.Args) < 3 {
        fmt.Println("Usage : bf-compiler <filename> <mode>")
        return
    }

    filePath := os.Args[1]
    mode := os.Args[2]

    if mode == "interpret" {
        start := time.Now()
        Interpret(filePath)
        elapsed := time.Since(start)
        fmt.Printf("Mode: %s, Time taken: %s\n", mode, elapsed)
    }
}

func must(err error){
    if err != nil {
        fmt.Println("Error", err)
        os.Exit(1)
    }
}
