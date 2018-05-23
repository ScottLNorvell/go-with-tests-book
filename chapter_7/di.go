package main

import (
  "fmt"
  _ "os"
  "io"
  "net/http"
)

func Encourage(writer io.Writer, name string) {
  fmt.Fprintf(writer, "%s is Great!", name)
}

func MyEncourageHandler(w http.ResponseWriter, r *http.Request) {
  Encourage(w, "Scott")
}

func main() {
  http.ListenAndServe(":5000", http.HandlerFunc(MyEncourageHandler))
}
