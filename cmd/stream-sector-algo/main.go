package main

import (
  "os"
  "bufio"
  "fmt"
)

func main() {
  seen := make(map[string]bool)
  input := bufio.NewScanner(os.Stdin)
  for input.Scan() {
    line := input.Text()
    if !seen[line] {
      seen[line] = true
      fmt.Println(line)
    }
    fmt.Println(seen)

  }
}
