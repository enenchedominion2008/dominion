package main

import (
	"io"
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	foundError := false

	// If no arguments, read from stdin
	if len(args) == 0 {
		_, err := io.Copy(os.Stdout, os.Stdin)
		if err != nil {
			printError(err.Error())
			foundError = true
		}
	} else {
		// Process each file argument
		for _, filename := range args {
			file, err := os.Open(filename)
			if err != nil {
				printError("ERROR: open " + filename + ": no such file or directory")
				foundError = true
				continue
			}

			_, err = io.Copy(os.Stdout, file)
			if err != nil {
				printError(err.Error())
				foundError = true
			}
			file.Close()
		}
	}

	// If we found any errors, force a non-zero exit using panic
	if foundError {
		panic("exit with error")
	}
}

func printError(msg string) {
	for _, r := range msg {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}
