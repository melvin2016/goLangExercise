package main

import (
	"fmt"
	"io"
	"net/http"
)

type consoleLogger struct{}

func main() {
	resp, err := http.Get("http://google.com")
	cl := consoleLogger{}
	if err != nil {
		fmt.Println("Error: ", err)
	}

	io.Copy(cl, resp.Body)

}

func (consoleLogger) Write(bs []byte) (n int, err error) {
	fmt.Println("Response Is : ", string(bs))
	return len(bs), nil
}
