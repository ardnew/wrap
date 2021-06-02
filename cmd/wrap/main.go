package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/ardnew/wrap"
)

func main() {

	var argNumCols int
	var argHyphen string
	var argNewline bool

	flag.IntVar(&argNumCols, "c", 80, "Wrap lines with length greater than or equal to `col` columns.")
	flag.StringVar(&argHyphen, "h", "-", "Append `hyp` to lines containing words longer than wrap length.")
	flag.BoolVar(&argNewline, "n", false, "Suppress printing the final line break.")
	flag.Parse()

	sc := bufio.NewScanner(in(flag.Args()))
	for sc.Scan() {
		w := wrap.String(sc.Text(), argNumCols, argHyphen)
		for i, s := range w {
			fmt.Print(s)
			if i+1 < len(w) || !argNewline {
				fmt.Println()
			}
		}
	}
}

func in(args []string) io.Reader {
	switch len(args) {
	case 0:
		return os.Stdin
	case 1:
		if r, err := os.Open(args[0]); nil == err {
			return r
		}
	}
	return strings.NewReader(strings.Join(args, " "))
}
