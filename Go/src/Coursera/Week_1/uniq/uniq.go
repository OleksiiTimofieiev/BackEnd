package uniq

import (
	"bufio"
	"fmt"
	"io"
)

func uniq(input io.Reader, output io.Writer) error {
	in := bufio.NewScanner(input)
	var prev string

	for in.Scan() {
		txt := in.Text()

		if txt == prev {
			continue
		}

		if txt < prev {
			return fmt.Errorf("file not supported")
		}

		prev = txt

		fmt.Fprintln(output, txt)
	}
	return nil
}

// func main() {
// 	err := uniq(os.Stdin, os.Stdout)
// 	if err != nil {
// 		panic(err.Error())
// 	}
// }
