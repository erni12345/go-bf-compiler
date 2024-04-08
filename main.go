package main


import(
    "fmt"
    "os"
)

func main() {

    //checking if we have a file name
    if len(os.Args) < 2 {
        fmt.Println("Usage : bf-compiler <filename>")
        return
    }

    filePath := os.Args[1]
    Interpret(filePath)
}

func must(err error){
    if err != nil {
        fmt.Println("Error", err)
        os.Exit(1)
    }
}
