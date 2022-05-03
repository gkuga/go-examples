// https://qiita.com/sxarp/items/cd528a546d1537105b9d
package main

import (
	"fmt"
)

// IntとStringの直和型
type IntStringUnion interface {
	Match(Cases)
}

// Int/String型の場合に実行する処理を格納する構造体
type Cases struct {
	IntFunc    func(Int)
	StringFunc func(String)
}

type Int int

func (i Int) Match(c Cases) { c.IntFunc(i) } // IntStringUnionインターフェイスの実装

type String string

func (i String) Match(c Cases) { c.StringFunc(i) } // IntStringUnionインターフェイスの実装

func main() {
	// IntStringUnion型からなる配列
	unionArray := []IntStringUnion{String("1"), Int(2), String("123"), Int(4)}

	// 総和を計算する、Int => その値そのまま、String => 文字列の長さ、として各要素を評価して加算する
	sum := 0
	for _, item := range unionArray {
		item.Match(Cases{ // ここで型のパターンマッチを行う
			IntFunc:    func(i Int) { sum += int(i) },            // Int型だった場合の処理
			StringFunc: func(s String) { sum += len(string(s)) }, // String型だった場合の処理
		})
	}

	fmt.Printf("%d", sum) // => 10が出力される
}
