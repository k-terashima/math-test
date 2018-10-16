/*

	- 場合の数
	- 用意した数から無作為に選択し並べた場合のパターン数
*/

package main

import (
	"fmt"
	"math/rand"
	"reflect"
	"time"
)

var (
	// 被り判定を行うか
	checkDouble = false
)

func main() {
	start := time.Now()
	defer func() {
		end := time.Now()
		fmt.Println("exec time: ", end.Sub(start))
	}()

	var (
		// 用意する数
		size = 1000001

		// いくつ部屋あるの？
		num = 5

		// いくつパターンか保存
		pattern [][]int
	)

	rand.Seed(time.Now().UnixNano())
	p := make([]int, size)
	for i := 0; i < size; i++ {
		p[i] = rand.Intn(10)
	}

	for range p {
		room := make([]int, num)
		for i := 0; i < num; i++ {
			l := len(p)
			n := rand.Intn(l)
			// room配列に必要な素材がなくなればbreak
			if num-i > l {
				break
			}

			// パターンを作る
			room[i] = p[n]

			// 使った素材は都度除外
			p = append(p[:n], p[n+1:]...)
		}

		// fmt.Printf("%v\n", r)
		pattern = append(pattern, room)

		// 使った素材を除外し続け，パターン長より少なくなったら作業終了
		l := len(p)
		if l < num {
			break
		}
	}

	// 最後に一度被りチェック
	var (
		copy      = pattern
		isDoubles int
	)
	if checkDouble {
		for k, v := range copy {
			var (
				isDouble int
			)
			for i, value := range pattern {
				if reflect.DeepEqual(v, value) {
					isDouble++
					if isDouble > 1 {
						isDoubles++
						// fmt.Println("is double: ", k, i, len(pattern), v, value)
						pattern = append(pattern[:i], pattern[i+1:]...)

						if isDouble > 2 {
							fmt.Println("orver double: ", k, i, len(pattern), v, value)
						}
					}
				}
			}
		}
	}

	fmt.Println("パターン数: ", len(pattern))
	// fmt.Printf("%+v\n", pattern)
	fmt.Printf("あまり配列: %+v\n", p)

	if checkDouble {
		fmt.Println("被り配列: ", isDoubles)
	}

}
