package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/dpotapov/go-spnego"
)

func main() {
	c := &http.Client{
		Transport: &spnego.Transport{},
	}

	if len(os.Args) < 2 {
		fmt.Printf("Usage:\n\t%s <URL>\n", os.Args[0])
		return
	}

	resp, err := c.Get(os.Args[1])
	if err != nil {
		panic(err)
	}

	fmt.Print(resp.Status)
}
