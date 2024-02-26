package main


import (
    "fmt"
    "os"
    "os/user"
    "APL2/repl"
)


func main() {

    // check if file was inputted
    if len(os.Args) > 1 {
        filename := os.Args[1]
        // check that file ends in .apl
        if filename[len(filename)-4:] != ".apl" {
            fmt.Printf("Error: file must end in .apl\n")
            os.Exit(1)
        }
        repl.StartFile(filename, os.Stdout)
   } else {
        startRepl()
    }


}

func startRepl() {
    user, err := user.Current()
    if err != nil {
        panic(err)
    }
    fmt.Printf("Hello %s! This is the A Programming Language 2.0!\n", user.Username)
    fmt.Printf("Feel free to type in commands\n")
    repl.Start(os.Stdin, os.Stdout)
}
