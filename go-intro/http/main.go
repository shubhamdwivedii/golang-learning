package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("http://google.com")

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	// fmt.Println(resp) // This will not print out the response body. Just bunch of response info data such as 200 OK Content-Type etc.

	/* Response Struct {
		Status string
		StatusCode int
		Body io.ReadCloser // Structs can have interfaces as field types
	}

	io.ReadCloser Interface {
		Reader
		Closer // Interface can have other interfaces as fields
		// This means that ReadCloser interface will have all methods of both Reader and Closer interfaces combined.
	}

	Reader Interface {
		Read([]byte) (int, error) // int is the number of bytes just read.
	}
	The Read([]byte) takes a byte slice to load data into.

	Reader Interface takes radically different sources of input (text, images, html etc) and
	translates them into some common medium that we can easily work with.
	*/

	bs := make([]byte, 99999) // this byte slice will hold the data we read using Read().
	resp.Body.Read(bs)        // response Body implements the Reader Interface (thus has a Read function)
	fmt.Println(string(bs))   // This will print the HTML code for google homepage

	// This will do exact same thing as above code.
	io.Copy(os.Stdout, resp.Body) // os.Stdout implements Writer Interface

	/* The Writer Interface takes some data (byte slice) and sends it to some output channel (Stdout in this case)

	Writer interface {
		Write([]byte) (int, error)
	}

	The []byte here is used as a source of input (unlike in Read)

	func Copy(dst Writer, src Reader) (written int64, err error)

	io.Copy copies from src to dst until either EOF is reached or error occur. */

	// ###### See Custom Writer (custom-writer.go) #######
}
