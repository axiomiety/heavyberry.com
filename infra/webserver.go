package main

import (
  "net/http"
  "log"
)

func main() {
  err := http.ListenAndServe(":80", http.FileServer(http.Dir("../site")))
  if err != nil {
    log.Printf("error running docs webserver: %v", err)
  }
}
