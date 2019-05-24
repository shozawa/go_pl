package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/shozawa/go_pl/ch07/evaluator"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	env := make(evaluator.Env)
	for {
		fmt.Printf(">> ")
		scanned := scanner.Scan()
		if !scanned {
			return
		}
		fmt.Println(evaluator.Eval(scanner.Text(), env))
	}
}
