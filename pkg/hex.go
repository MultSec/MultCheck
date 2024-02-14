package main

import (
	"fmt"
	"io"
	"os"
)

func hexDumpAround(filename string, offset, context int64) error {
	const bytesPerLine = 16

	// Open binary file
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	// Calculate start position, ensure it's not negative
	start := offset - context
	if start < 0 {
		start = 0
	}
	end := offset + context

	// Seek to start position
	_, err = file.Seek(start, io.SeekStart)
	if err != nil {
		return err
	}

	data := make([]byte, end-start)
	_, err = file.Read(data)
	if err != nil {
		return err
	}

	mid := int64(len(data)) / 2

	// Print hex dump from start to end
	for i := int64(0); i < int64(len(data)); i += bytesPerLine {
		chunk := data[i:min(i+bytesPerLine, int64(len(data)))]

		// Print the offset in the data stream
		fmt.Printf("%08x  ", start+i)

		// Print the hex codes
		for j, b := range chunk {
			fmt.Printf("%02x ", b)
			if j%2 == 1 {
				fmt.Print(" ")
			}
		}

		// Pad the line with spaces if it's shorter than bytesPerLine
		if len(chunk) < bytesPerLine {
			for j := len(chunk); j < bytesPerLine; j++ {
				fmt.Print("   ")
				if j%2 == 1 {
					fmt.Print(" ")
				}
			}
		}

		colored := (i == mid-bytesPerLine || i == mid || i == mid+bytesPerLine)
		if colored {
			fmt.Print("\033[1;31m")
		}

		// Print the ASCII representation
		for _, b := range chunk {
			if b >= 32 && b <= 126 {
				fmt.Printf("%c", b)
			} else {
				fmt.Print("·")
			}
		}

		if colored {
			fmt.Print("\033[0m")
		}

		fmt.Println()
	}

	return nil
}

// min returns the smaller of x or y.
func min(x, y int64) int64 {
	if x < y {
		return x
	}
	return y
}
