// Code generated by copypasta/template/leetcode/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/leetcode/testutil"
	"testing"
)

func Test(t *testing.T) {
	t.Log("Current test is [b]")
	examples := [][]string{
		{
			`[1,3,5,10]`, `2`, 
			`2`,
		},
		{
			`[3,6,8,10]`, `3`, 
			`0`,
		},
		{
			`[3,4,6,7]`, `2`, 
			`1`,
		},
		{
			`[5]`, `10`, 
			`0`,
		},
		
	}
	targetCaseNum := 0 // -1
	if err := testutil.RunLeetCodeFuncWithExamples(t, addRungs, examples, targetCaseNum); err != nil {
		t.Fatal(err)
	}
}
// https://leetcode-cn.com/contest/weekly-contest-250/problems/add-minimum-number-of-rungs/