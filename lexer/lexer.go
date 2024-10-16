package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Open("input.md")
	if err != nil {
		panic(err)
	}

	// TODO look up defer meaning
	defer func() {
		if err := file.Close(); err != nil {
			panic(err)
		}
	}()

	reader := bufio.NewReader(file)

	output, err := os.Create("output.html")
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := output.Close(); err != nil {
			panic(err)
		}
	}()

	writeBuf := bufio.NewWriter(output)

	buf := make([]byte, 1024)
	for {
		curr, err := reader.Read(buf)
		if err != nil && err != io.EOF {
			panic(err)
		}
		if curr == 0 {
			break
		}

		if _, err := writeBuf.Write(buf[:curr]); err != nil {
			panic(err)
		}
	}

	if err = writeBuf.Flush(); err != nil {
		panic(err)
	}

	fmt.Println("Finished")
}
