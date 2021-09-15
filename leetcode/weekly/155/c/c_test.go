// Code generated by generator_test.
package main

import (
	"github.com/EndlessCheng/codeforces-go/leetcode/testutil"
	"testing"
)

func Test(t *testing.T) {
	t.Log("Current test is [c]")
	exampleIns := [][]string{{`"dcab"`, `[[0,3],[1,2]]`}, {`"dcab"`, `[[0,3],[1,2],[0,2]]`}, {`"cba"`, `[[0,1],[1,2]]`}}
	exampleOuts := [][]string{{`"bacd"`}, {`"abcd"`}, {`"abc"`}}
	// custom test cases or WA cases.
	//exampleIns = append(exampleIns, []string{``})
	//exampleOuts = append(exampleOuts, []string{``})
	if err := testutil.RunLeetCodeFunc(t, smallestStringWithSwaps, exampleIns, exampleOuts); err != nil {
		t.Fatal(err)
	}
}