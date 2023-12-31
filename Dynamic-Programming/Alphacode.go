// https://www.hackerearth.com/problem/algorithm/acode-alphacode-3/?purpose=login&source=problem-page&update=google

package main

import (
	"fmt"
)

var s string
var memo []int

func dp(i int) int {
	if i == len(s) {
		return 1
	}

	if memo[i] != -1 {
		return memo[i]
	}

	var ans int = 0
	if s[i] >= '1' && s[i] <= '9' {
		ans += dp(i + 1)
	}

	if i+1 < len(s) && (s[i] == '1') {
		ans += dp(i + 2)
	}

	if i+1 < len(s) && (s[i] == '2' && s[i+1] <= '6') {
		ans += dp(i + 2)
	}

	memo[i] = ans
	return ans
}

func main() {
	for {
		fmt.Scan(&s)
		if s[0] == '0' {
			break
		}
		memo = make([]int, len(s))
		for i := range memo {
			memo[i] = -1
		}
		fmt.Println(dp(0))
	}
}
