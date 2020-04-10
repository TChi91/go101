package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		fmt.Fprint(res, "Hello Http!\n")
		io.WriteString(res, "Cv chiwaa")
	})
	http.ListenAndServe(":8000", nil)
}
