package main

import (
  "log"
  "net/http"
)

func main() {
  assets := http.StripPrefix("/", http.FileServer(http.Dir("assets/")))
  http.Handle("/", assets)

  log.Println("Listening at port 3000")
  log.Fatal(http.ListenAndServe(":3000", nil))
}
