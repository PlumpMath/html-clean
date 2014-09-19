package main

import (
  "fmt"
  "io/ioutil"
  "os"
  "strings"
)

func check(e error) {
  if e != nil {
    panic(e)
  }
}

func main() {
  args := os.Args[1:]

  if len(args) == 0 {
    fmt.Println("Must provide input html")
    os.Exit(1)
  }

  filename := args[0]

  b, err := ioutil.ReadFile(filename)
  check(err)

  contents := string(b)
  
  contents = strings.Replace(contents, "”", "&rdquo;", -1)
  contents = strings.Replace(contents, "“", "&ldquo;", -1)
  contents = strings.Replace(contents, "‘", "&lsquo;", -1)
  contents = strings.Replace(contents, "’", "&lsquo;", -1)

  werr := ioutil.WriteFile(filename, []byte(contents), 0666)
  check(werr)

  fmt.Println("File cleaned.")
}
