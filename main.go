package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/google/uuid"
)

const genArg = "gen"

func main() {
	args := os.Args[1:]
	if len(args) < 1 {
		fmt.Println("invalid arguments")
		return
	}
	if strings.ToLower(args[0]) == genArg {
		fmt.Println(generate())
		return
	}
	fmt.Println("invalid arguments")
}

func generate() uuid.UUID {
	return uuid.New()
}
