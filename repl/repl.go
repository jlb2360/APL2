package repl

import (
    "bufio"
    "fmt"
    "io"
    "APL2/lexer"
    "APL2/parser"
    "APL2/eval"
    "APL2/object"
    "os"
)


const PROMPT = ">> "


func Start(in io.Reader, out io.Writer) {
    scanner := bufio.NewScanner(in)
    env := object.NewEnvironment()
    for {
        fmt.Printf(PROMPT)
        scanned := scanner.Scan()
        if !scanned {
            return
        }

        line := scanner.Text()

        l := lexer.New(line)

        p := parser.New(l)

        program := p.ParseProgram()
        if len(p.Errors()) != 0 {
            printParserErrors(out, p.Errors())
            continue
        }



        evaluated := eval.Eval(program, env)
        if evaluated != nil {
            io.WriteString(out, evaluated.Inspect())
            io.WriteString(out, "\n")
        }
    }
}

func StartFile(filename string, out io.Writer) {
    f, err := os.Open(filename)
    if err != nil {
        fmt.Printf("Error: %s\n", err)
        return
    }
    defer f.Close()

    scanner := bufio.NewScanner(f)

    env := object.NewEnvironment()

    for scanner.Scan() {
        line := scanner.Text()
        l := lexer.New(line)
        p := parser.New(l)

        program := p.ParseProgram()
        if len(p.Errors()) != 0 {
            printParserErrors(out, p.Errors())
            continue
        }

        evaluated := eval.Eval(program, env)
        if evaluated != nil {
            io.WriteString(out, evaluated.Inspect())
            io.WriteString(out, "\n")
        }
    }
}


func printParserErrors(out io.Writer, errors []string) {
    for _, msg := range errors {
        io.WriteString(out, "\t" + msg + "\n")
    }
}
