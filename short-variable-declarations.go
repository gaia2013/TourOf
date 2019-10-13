package main

import "fmt"

func main() {
    var i, j int = 1, 2
    k := 3 // 暗黙的な型宣言
    c, python, java := true, false, "no!"

    fmt.Println(i, j, k, c, python, java)
}

// ※関数外では、キーワードで始まる宣言（var,func,など）が必要
