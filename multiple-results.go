package main

import "fmt"

func swap(x, y string) (string, string) { // swap関数は２つのstringを返す
  return y, x
}

func main() {
  a, b := swap("hello", "world")
  fmt.Println(a, b)
}
