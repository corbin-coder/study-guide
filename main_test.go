package main

import "testing"

func TestRewrite2(t *testing.T) {
	nums := []int{1, 2, 3, 5, 7, 8, 9, 10}
	res := rewrite2(nums)
	t.Log(res)
}
