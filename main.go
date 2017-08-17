package main

import (
	"io"
	"strconv"

	"github.com/mattn/go-colorable"
)

func printSGR(w io.Writer, params ...int) {
	io.WriteString(w, "\x1b[")
	for i, p := range params {
		if i > 0 {
			io.WriteString(w, ";")
		}
		io.WriteString(w, strconv.Itoa(p))
	}
	io.WriteString(w, "m")
}

func printReset(w io.Writer) {
	printSGR(w, 0)
}

func print256Color(w io.Writer, color int) {
	printSGR(w, 48, 5, color)
	io.WriteString(w, "  ")
}

func main() {
	w := colorable.NewColorableStdout()

	// System colors
	io.WriteString(w, "System colors:\n")
	for n := 0; n < 16; n++ {
		print256Color(w, n)

		if (n+1)%8 == 0 {
			printReset(w)
			io.WriteString(w, "\n")
		}
	}

	// 6x6x6 color cube
	io.WriteString(w, "\nColor cube 6x6x6:\n")
	for g := 0; g < 6; g++ {
		for r := 0; r < 6; r++ {
			for b := 0; b < 6; b++ {
				color := 16 + b + g*6 + r*36
				print256Color(w, color)
			}
			printReset(w)
			io.WriteString(w, "  ")
		}
		printReset(w)
		io.WriteString(w, "\n")
	}

	// Grayscale colors
	io.WriteString(w, "\nGrayscale colors:\n")
	for i := 232; i < 256; i++ {
		print256Color(w, i)
	}
	printReset(w)
	io.WriteString(w, "\n")
}
