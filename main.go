package main

import (
	"bytes"
	"flag"
	"fmt"
	"image/png"
	"net/http"
	"strconv"

	"github.com/go-mimetex/mimetex"
)

func handler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path[1:] == "" || r.URL.Path[1:] == "\\" {
		http.Error(w, "Bad formula", 500)
		return
	}
	img, err := mimetex.RenderImage(r.URL.Path[1:], 5)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	buffer := new(bytes.Buffer)
	if err := png.Encode(buffer, img); err != nil {
		fmt.Println("Failed to encode image")
	}

	w.Header().Set("Content-Type", "image/png")
	w.Header().Set("Content-Length", strconv.Itoa(len(buffer.Bytes())))
	if _, err := w.Write(buffer.Bytes()); err != nil {
		fmt.Println("Failed to reply with image")
	}
}

func main() {
	portFlag := flag.Int("p", 8888, "Port")
	flag.Parse()

	http.HandleFunc("/", handler)
	http.ListenAndServe(":"+strconv.Itoa(*portFlag), nil)
}
