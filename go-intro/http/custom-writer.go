package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func main() {
	resp, err := http.Get("http://google.com")

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	// Custom Writer
	lw := logWriter{}
	io.Copy(lw, resp.Body)
}

// Custom Writer (This implements Writer Interface)
func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("==============FINISHED=============")
	return len(bs), nil
}
