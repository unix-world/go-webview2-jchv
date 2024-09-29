package main

// GOOS=windows GOARCH=amd64 go build ./cmd/demo/main.go

import (
	"log"

	webview2 "github.com/unix-world/go-webview2-jchv"
)

func main() {
	w := webview2.NewWithOptions(webview2.WebViewOptions{
		Debug:     true,
		AutoFocus: true,
		WindowOptions: webview2.WindowOptions{
			Title:  "Minimal webview2 example",
			Width:  800,
			Height: 600,
			IconId: 2, // icon resource id
			Center: true,
		},
	})
	if w == nil {
		log.Fatalln("Failed to load webview2.")
	}
	defer w.Destroy()
	w.SetSize(800, 600, webview2.HintFixed)
	w.Navigate("https://en.m.wikipedia.org/wiki/Main_Page")
	w.Run()
}
