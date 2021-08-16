// Code generated by copypasta/template/atcoder/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// 提交地址：https://atcoder.jp/contests/abc214/submit?taskScreenName=abc214_c
func Test_run(t *testing.T) {
	t.Log("Current test is [c]")
	testCases := [][2]string{
		{
			`3
4 1 5
3 10 100`,
			`3
7
8`,
		},
		{
			`4
100 100 100 100
1 1 1 1`,
			`1
1
1
1`,
		},
		{
			`4
1 2 3 4
1 2 4 7`,
			`1
2
4
7`,
		},
		{
			`8
84 87 78 16 94 36 87 93
50 22 63 28 91 60 64 27`,
			`50
22
63
28
44
60
64
27`,
		},
		
	}
	testutil.AssertEqualStringCase(t, testCases, 0, run)
}
// https://atcoder.jp/contests/abc214/tasks/abc214_c
