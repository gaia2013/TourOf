package main

import "fmt"

func add(x, y int) int { // 関数の２つ以上の引数が同じ型なら最後の型を残して省略記述可能
  return x + y
}

func main() {
  fmt.Println(add(42, 13))
}
