// 関数は０個以上の引数を取れる

package main

import "fmt"

func add(x int, y int) int { // 変数名の 後ろ に型名を書く
  return x + y
}

func main() {
  fmt.Println(add(42, 13))
}
