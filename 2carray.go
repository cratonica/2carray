// Simple utility to convert a file into a C byte array
// Clint Caywood
// http://github.com/cratonica/2carray
package main

import (
	"code.google.com/p/go.crypto/ssh/terminal"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
	"syscall"
)

func main() {
	flag.Parse()

	if terminal.IsTerminal(syscall.Stdin) {
		flag.Usage()
		fmt.Println("\nPlease pipe the file you wish to encode into stdin\n")
		return
	}

	if len(os.Args) != 2 {
		fmt.Println("Usage: 2carray array_name")
		return
	}
	allCaps := strings.ToUpper(os.Args[1])
	fmt.Printf("#ifndef %s_H_INCLUDED\n", allCaps)
	fmt.Printf("#define %s_H_INCLUDED\n\n", allCaps)
	fmt.Printf("const unsigned char %s[] = {", os.Args[1])
	buf := make([]byte, 1)
	var err error
	var totalBytes uint64
	var n int
	for n, err = os.Stdin.Read(buf); n > 0 && err == nil; {
		if totalBytes > 0 {
			fmt.Print(", ")
		}
		if totalBytes%12 == 0 {
			fmt.Printf("\n\t")
		}
		fmt.Printf("0x%02x", buf[0])
		totalBytes++
		n, err = os.Stdin.Read(buf)
	}
	if err != nil && err != io.EOF {
		fmt.Errorf("Error: %v", err)
	}
	fmt.Print("\n};\n\n")
	fmt.Printf("#endif /* %s_H_INCLUDED */\n\n", allCaps)
}
