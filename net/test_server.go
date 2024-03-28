package net

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
)

func TestServer() {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "The host received by the server is %s\n", r.Host)
	}))
	defer ts.Close()
	fmt.Printf("The server is at %s\n", ts.URL)

	client := ts.Client()

	req, err := http.NewRequest("GET", ts.URL, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Host = "aaa.com"

	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	io.Copy(os.Stdout, res.Body)
}
