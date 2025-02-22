package main

import "github.com/CrazyCatViking/go-http/http"

func main() {
  server, _ := http.Create(":8080")
  defer server.Close()

  server.Get("/", func (r *http.Request) {
    r.Html("Hello World")
  })

  server.Listen()
}

